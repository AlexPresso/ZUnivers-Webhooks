package utils

import (
	"bytes"
	"encoding/json"
	"github.com/spf13/viper"
	"net/http"
	"time"
)

func Request(uri, method string, body []byte, structure interface{}) (err error) {
	client := &http.Client{
		Timeout: viper.GetDuration("api.timeout") * time.Second,
	}

	req, err := http.NewRequest(method, viper.GetString("api.baseUrl")+uri, bytes.NewBuffer(body))
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

	err = json.NewDecoder(r.Body).Decode(structure)
	return
}
