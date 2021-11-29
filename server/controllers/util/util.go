package util

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/viper"
)

func Request(method, uri string, body []byte) (map[string]interface{}, error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, uri, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	auth := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s/token:%s", viper.GetString("ZENDESK_CRED_EMAIL"), viper.GetString("ZENDESK_CRED_API_TOKEN"))))
	req.Header.Add("Authorization", fmt.Sprintf("Basic %s", auth))
	req.Header.Add("Content-Type", "application/json;charset=UTF-8")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	rawData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var data map[string]interface{}
	err = json.Unmarshal([]byte(rawData), &data)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 && res.StatusCode != 300 {
		return data, fmt.Errorf("%d", res.StatusCode)
	}

	return data, nil
}
