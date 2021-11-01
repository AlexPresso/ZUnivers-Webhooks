package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/alexpresso/zunivers-webhooks/structures/discord"
	"github.com/alexpresso/zunivers-webhooks/utils"
	"github.com/spf13/viper"
	"net/http"
	"reflect"
	"strings"
)

func DispatchEvent(event string, oldObject, newObject interface{}) {
	body, _ := json.Marshal(makeFormData(event, oldObject, newObject))

	for _, url := range viper.GetStringSlice(fmt.Sprintf("webhooks.%s.urls", event)) {
		url := url

		res, err := http.Post(fmt.Sprintf("%s?wait=true", url), "application/json", bytes.NewBuffer(body))
		if err != nil {
			utils.Log(fmt.Sprintf("Failed to dispatch %s event", event))
		}

		fmt.Println(res)
	}

	utils.Log("Dispatched event: " + event)
}

func makeFormData(event string, oldObject, newObject interface{}) *discord.WebhookFormData {
	embed := &discord.Embed{
		Title:       "",
		Type:        "rich",
		Description: viper.GetString(fmt.Sprintf("webhooks.%s.message", event)),
		Color:       374272,
		Author: &discord.Author{
			Name:    "ZUnivers",
			IconURL: viper.GetString("frontBaseUrl") + "/img/logo-mini.aea51074.png",
			URL:     viper.GetString("frontBaseUrl"),
		},
	}

	fillEmbed(embed, oldObject, newObject)

	return &discord.WebhookFormData{
		Username:  "ZUnivers-Webhooks",
		AvatarURL: viper.GetString("frontBaseUrl") + "/img/logo-mini.aea51074.png",
		Embeds:    []*discord.Embed{embed},
	}
}

func fillEmbed(embed *discord.Embed, oldObject, newObject interface{}) {
	if newObject == nil {
		return
	}

	if oldObject != nil {
		oldObject = reflect.ValueOf(oldObject)
	}

	newType := reflect.TypeOf(newObject)
	newObject = reflect.ValueOf(newObject)
	embedField := &discord.EmbedField{
		Name:   "Détails",
		Value:  "",
		Inline: false,
	}

	for i := 0; i < newType.NumField(); i++ {
		if tagValue, hasTag := newType.Field(i).Tag.Lookup("zu"); hasTag {
			parts := strings.Split(tagValue, ";")

			var oldValue interface{}
			newValue := newObject.(reflect.Value).Field(i).Interface()
			if oldObject != nil {
				oldValue = oldObject.(reflect.Value).Field(i).Interface()
			}

			for _, part := range parts {
				part := strings.Split(part, "=")

				switch part[0] {
				case "display":
					processDisplay(embedField, oldValue, newValue, part)
					break
				case "imageUrl":
					processImage(embed, newValue, part)
					break
				case "url":
					embedField.Value += fmt.Sprintf("\n[Page de l'entité](%s)", viper.GetString("frontBaseUrl")+fmt.Sprintf(part[1], newValue))
					break
				}
			}
		}
	}

	embed.Fields = []*discord.EmbedField{embedField}
}

func processDisplay(field *discord.EmbedField, oldValue, newValue interface{}, parts []string) {
	oldValueText := ""
	//TODO: check for time castable and time.Time.Equal
	if oldValue != nil && oldValue != newValue {
		oldValueText = fmt.Sprintf("`%s` → ", fmt.Sprint(oldValue))
	}

	field.Value += fmt.Sprintf("__%s:__ %s`%s`\n", parts[1], oldValueText, fmt.Sprint(newValue))
}

func processImage(embed *discord.Embed, newValue interface{}, parts []string) {
	cdnBase := viper.GetString("cdnBaseUrl")
	if strings.Compare(parts[1], "%s") == 0 {
		cdnBase = ""
	}

	embed.Thumbnail = &discord.EmbedMedia{
		Url: cdnBase + fmt.Sprintf(parts[1], newValue),
	}
}
