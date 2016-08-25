package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func getAccessToken(c ClientConfig) (string, error) {

	values := url.Values{"grant_type": {"refresh_token"}, "refresh_token": {c.RefreshToken}}.Encode()
	req, err := http.NewRequest("POST", ACCESS_TOKEN_URI, strings.NewReader(values))

	if err != nil {
		return "", err
	}

	req.SetBasicAuth(c.ClientID, c.ClientSecret)
	req.Header.Add("content-type", "application/x-www-form-urlencoded")

	client := http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.Do(req)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	var oc Oauth2Config
	err = json.Unmarshal(respBody, &oc)

	if err != nil {
		return "", err
	}

	c.RefreshToken = oc.RefreshToken

	err = saveClientConfig(c)

	if err != nil {
		return "", err
	}

	return oc.AccessToken, nil
}
