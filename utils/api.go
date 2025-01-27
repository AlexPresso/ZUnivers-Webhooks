package utils

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"github.com/alexpresso/zunivers-webhooks/structures"
	"github.com/spf13/viper"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"net/http"
	"time"
)

const ResponseChangedEvent = "response_changed"

func Request(uri, method string, body []byte, structure interface{}, resSpec *structures.JsonResponseSpec) (err error) {
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
	req.Header.Set("user-agent", "github.com/alexpresso/zunivers-webhooks")
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

	bodyBytes, err := io.ReadAll(reader)
	if err != nil {
		Log(fmt.Sprintf("Error reading request body: %v", err))
		return
	}

	switch structure.(type) {
	case *image.Image:
		*(structure.(*image.Image)), _, err = image.Decode(bytes.NewReader(bodyBytes))
		if err != nil {
			return err
		}
	default:
		err = json.NewDecoder(bytes.NewReader(bodyBytes)).Decode(structure)

		if !EventsAllDisabled([]string{ResponseChangedEvent}) {
			var specMap interface{}

			err = json.NewDecoder(bytes.NewReader(bodyBytes)).Decode(&specMap)
			if err != nil {
				Log("Error while decoding data")
				return
			}

			var prettySpec []byte

			prettySpec, err = json.MarshalIndent(replaceValuesWithTypes(specMap), "", "  ")
			if err != nil {
				Log("Error while decoding data")
				return
			}

			*resSpec = structures.JsonResponseSpec{
				EndpointURI: uri,
				Value:       string(prettySpec),
			}
		}
	}

	return
}

func replaceValuesWithTypes(data interface{}) interface{} {
	switch v := data.(type) {
	case string:
		return "string"
	case float64:
		return "number"
	case bool:
		return "bool"
	case nil:
		return "null"
	case map[string]interface{}:
		for key, value := range v {
			v[key] = replaceValuesWithTypes(value)
		}
		return v
	case []interface{}:
		if len(v) == 0 {
			return v
		}

		v[0] = replaceValuesWithTypes(v[0])
		return []interface{}{v[0]}
	default:
		return fmt.Sprintf("unknown (%T)", v)
	}
}
