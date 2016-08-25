package main

import (
	"encoding/json"
	"io/ioutil"
)

func saveClientConfig(c ClientConfig) (err error) {

	var b []byte

	if b, err = json.Marshal(c); err != nil {
		return
	}

	return ioutil.WriteFile(CLIENT_CONFIG_FILE, b, 0644)
}

func loadClientConfig() (c ClientConfig, err error) {

	var b []byte

	if b, err = ioutil.ReadFile(CLIENT_CONFIG_FILE); err != nil {
		return
	}

	err = json.Unmarshal(b, &c)

	return
}
