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

		_, err := http.Post(fmt.Sprintf("%s?wait=true", url), "application/json", bytes.NewBuffer(body))
		if err != nil {
			utils.Log(fmt.Sprintf("Failed to dispatch %s event", event))
		}
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
		if name, hasTag := newType.Field(i).Tag.Lookup("display"); hasTag {
			oldValueText := ""
			newValue := newObject.(reflect.Value).Field(i).Interface()

			if oldObject != nil {
				oldValue := oldObject.(reflect.Value).Field(i).Interface()
				if oldValue != newValue {
					oldValueText = fmt.Sprintf("`%s` → ", fmt.Sprint(oldValue))
				}
			}

			embedField.Value += fmt.Sprintf("__%s:__ %s`%s`\n", name, oldValueText, fmt.Sprint(newValue))
		} else if uri, hasTag := newType.Field(i).Tag.Lookup("imageUrl"); hasTag {
			uriParts := strings.Split(uri, ";")
			cdnBase := viper.GetString("cdnBaseUrl")
			if strings.Compare(uri, "%s") == 1 {
				cdnBase = ""
			}

			media := &discord.EmbedMedia{
				Url: cdnBase + fmt.Sprintf(uriParts[0], newObject.(reflect.Value).Field(i).Interface()),
			}

			if len(uriParts) > 0 && uriParts[1] == "image" {
				embed.Image = media
			} else {
				embed.Thumbnail = media
			}

		}
	}

	embed.Fields = []*discord.EmbedField{embedField}
}
