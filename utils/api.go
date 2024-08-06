package utils

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"net/http"
	"time"
)

const ResponseChangedEvent = "response_changed"

func Request(uri, method string, body []byte, structure interface{}, resSpec map[string]interface{}) (err error) {
	client := &http.Client{
		Timeout: viper.GetDuration("api.timeout") * time.Second,
	}

	var baseUrl string
	switch structure.(type) {
	case *image.Image:
		baseUrl = viper.GetString("frontBaseUrl")
	default:
		baseUrl = viper.GetString("api.baseUrl")
	}

	req, err := http.NewRequest(method, baseUrl+uri, bytes.NewBuffer(body))
	if err != nil {
		return
	}

	req.Header.Set("accept", "application/json, text/plain, */*")
	req.Header.Set("accept-language", "fr-FR,fr;q=0.9,en-US;q=0.8,en;q=0.7")
	req.Header.Set("origin", viper.GetString("frontBaseUrl"))
	req.Header.Set("referer", viper.GetString("frontBaseUrl"))
	req.Header.Set("dnt", "1")
	req.Header.Set("sec-ch-ua-platform", "\"Chromium\";v=\"94\", \"Google Chrome\";v=\"94\", \";Not A Brand\";v=\"99\"")
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", "Windows")
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-site", "same-site")
	req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/94.0.4606.81 Safari/537.36")
	req.Header.Set("x-zunivers-rulesettype", viper.GetString("api.rulesettype"))

	r, err := client.Do(req)
	if err != nil {
		return
	}

	defer r.Body.Close()
	var reader io.Reader = r.Body
	if r.Header.Get("Content-Encoding") == "gzip" {
		reader, err = gzip.NewReader(r.Body)
		if err != nil {
			return
		}
		defer reader.(*gzip.Reader).Close()
	}

	switch structure.(type) {
	case *image.Image:
		*(structure.(*image.Image)), _, err = image.Decode(r.Body)
		if err != nil {
			return err
		}
	default:
		err = json.NewDecoder(r.Body).Decode(structure)

		if EventsEnabled([]string{ResponseChangedEvent}) {
			err = json.NewDecoder(reader).Decode(&resSpec)
			if err != nil {
				replaceValuesWithTypes(resSpec)
			}
		}
	}

	return
}

func replaceValuesWithTypes(data map[string]interface{}) {
	for key, value := range data {
		switch v := value.(type) {
		case map[string]interface{}:
			replaceValuesWithTypes(v)
		case []interface{}:
			data[key] = "array"
		case string:
			data[key] = "string"
		case float64:
			data[key] = "number"
		case bool:
			data[key] = "boolean"
		case nil:
			data[key] = "null"
		default:
			data[key] = fmt.Sprintf("unknown (%T)", v)
		}
	}
}
