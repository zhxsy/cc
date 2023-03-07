package net

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func Get(url string) (string, error) {
	res, err := http.Get(url)
	defer res.Body.Close()
	if err != nil {
		return "", err
	}
	data, err := ioutil.ReadAll(res.Body)
	return string(data), nil
}

func Post(url string, param map[string]interface{}) (string, error) {
	contentType := "application/json;charset=utf-8"
	byteData, _ := json.Marshal(param)
	res, err := http.Post(url, contentType, bytes.NewBuffer(byteData))
	defer res.Body.Close()
	if err != nil {
		return "", err
	}

	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	return string(content), nil
}
