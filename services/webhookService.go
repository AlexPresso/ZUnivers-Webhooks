package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"net/http"
	"reflect"
	"zunivers-webhooks/structures/discord"
	"zunivers-webhooks/utils"
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
	return &discord.WebhookFormData{
		Username:  "ZUnivers-Webhooks",
		AvatarURL: viper.GetString("frontBaseUrl") + "/img/logo-mini.aea51074.png",
		Embeds: []*discord.DiscordEmbed{
			{
				Title:       "",
				Type:        "rich",
				Description: viper.GetString(fmt.Sprintf("webhooks.%s.message", event)),
				Color:       374272,
				Fields:      makeFields(oldObject, newObject),
				Author: &discord.DiscordAuthor{
					Name:    "ZUnivers",
					IconURL: viper.GetString("frontBaseUrl") + "/img/logo-mini.aea51074.png",
					URL:     viper.GetString("frontBaseUrl"),
				},
			},
		},
	}
}

func makeFields(oldObject, newObject interface{}) (fields []*discord.DiscordEmbedField) {
	if newObject == nil {
		return
	}

	if oldObject != nil {
		oldObject = reflect.ValueOf(oldObject)
	}

	newType := reflect.TypeOf(newObject)
	newObject = reflect.ValueOf(newObject)
	embedField := &discord.DiscordEmbedField{
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
		}
	}

	fields = append(fields, embedField)
	return
}
