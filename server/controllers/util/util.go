package util

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Request(method, uri string, body []byte) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, uri, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Basic ")
	req.Header.Add("Content-Type", "application/json;charset=UTF-8")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 && res.StatusCode != 300 {
		return nil, fmt.Errorf("Failed request: %d", res.StatusCode)
	}

	rawData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return rawData, nil
}
