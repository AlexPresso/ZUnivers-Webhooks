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

func DispatchEmbeds(embeds *[]discord.Embed) {
	if len(*embeds) == 0 {
		return
	}

	var formDatas []*discord.WebhookFormData
	for i := 0; i < len(*embeds); i += 10 {
		end := i + 10
		if end > len(*embeds) {
			end = len(*embeds)
		}

		formData := &discord.WebhookFormData{
			Username:  "ZUnivers-Webhooks",
			AvatarURL: viper.GetString("frontBaseUrl") + "/img/logo-mini.aea51074.png",
			Embeds:    &[]discord.Embed{},
		}

		ping := ""
		for _, embed := range (*embeds)[i:end] {
			embed := embed
			*formData.Embeds = append(*formData.Embeds, embed)

			if len(embed.Role) > 0 {
				ping = ping + fmt.Sprintf("<@&%s> ", embed.Role)
			}
		}

		if len(ping) > 0 {
			formData.Content = ping
		}

		formDatas = append(formDatas, formData)
	}

	urls := viper.GetStringSlice("webhooks")
	perWh := (len(formDatas) / len(urls)) + 1

	if perWh > 5 {
		utils.Log("Too much messages to send, please add one ore more new webhooks URLS.")
	}

	for i := 0; i < len(urls); i++ {
		url := urls[i]
		offset := i * perWh
		end := offset + perWh

		if offset > len(formDatas) {
			break
		}

		if end > len(formDatas) {
			end = len(formDatas)
		}

		for _, formData := range formDatas[offset:end] {
			formData := &formData
			body, _ := json.Marshal(formData)

			res, err := http.Post(fmt.Sprintf("%s?wait=true", url), "application/json", bytes.NewBuffer(body))
			if err != nil {
				utils.Log(fmt.Sprintf("Failed to dispatch to: %s", url))
				fmt.Println(res)
			}
		}
	}

	utils.Log("Dispatched embeds.")
}

func DefaultEmbed(event string, placeholder string) *discord.Embed {
	description := viper.GetString(fmt.Sprintf("events.%s.message", event))
	role := viper.GetString(fmt.Sprintf("events.%s.role", event))

	if strings.Contains(description, "%s") {
		description = fmt.Sprintf(description, placeholder)
	}

	embed := &discord.Embed{
		Title:       "",
		Type:        "rich",
		Description: description,
		Color:       374272,
		Author: &discord.Author{
			Name:    "ZUnivers",
			IconURL: viper.GetString("frontBaseUrl") + "/img/logo-mini.aea51074.png",
			URL:     viper.GetString("frontBaseUrl"),
		},
	}

	if len(role) > 0 {
		embed.Role = role
	}

	return embed
}

func MakeEmbed(event string, oldObject, newObject interface{}) *discord.Embed {
	embed := DefaultEmbed(event, "")
	fillEmbed(embed, oldObject, newObject)

	return embed
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
	url := ""
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
					url = viper.GetString("frontBaseUrl") + fmt.Sprintf(part[1], newValue)
					break
				}
			}
		}
	}

	if url != "" {
		embedField.Value += fmt.Sprintf("\n[Page de l'entité](%s)", strings.ReplaceAll(url, " ", "-"))
	}

	embed.Fields = []*discord.EmbedField{embedField}
}

func processDisplay(field *discord.EmbedField, oldValue, newValue interface{}, parts []string) {
	split := strings.Split(parts[1], "|")
	format := "`%v`"
	if len(split) > 1 {
		format = split[1]
	}

	oldValueText := ""
	if oldValue != nil {
		if utils.IsTime(oldValue) {
			if utils.TimeDifference(oldValue, newValue) {
				oldValueText = fmt.Sprintf("`%s` → ", fmt.Sprintf(format, oldValue))
			}
		} else if oldValue != newValue {
			oldValueText = fmt.Sprintf("`%s` → ", fmt.Sprintf(format, oldValue))
		}
	}

	value := fmt.Sprintf("__%s:__ %s%s\n", split[0], oldValueText, fmt.Sprintf(format, newValue))
	field.Value += value
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
