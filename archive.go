package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	AccessTokenURI  = "https://bitbucket.org/site/oauth2/access_token"
	AccessTokenBody = "grant_type=authorization_code&code=%s"
)

var clientConfig ClientConfig

type Repository struct {
}

func runArchive() (err error) {
	var o2c *Oauth2Config

	if err = o2c.getRefreshToken(); err != nil {
		return
	}

	return
}

func appendAuthHeader(req *http.Request) {

}

func (o2c *Oauth2Config) getRefreshToken() (err error) {
	var (
		cbytes []byte
		req    *http.Request
		resp   *http.Response
	)

	if cbytes, err = ioutil.ReadFile(ClientConfigFile); err != nil {
		return
	}

	if err = json.Unmarshal(cbytes, &clientConfig); err != nil {
		return
	}
	accessTokenBody := fmt.Sprintf(AccessTokenBody, clientConfig.Code)
	if req, err = http.NewRequest("POST", AccessTokenURI, strings.NewReader(accessTokenBody)); err != nil {
		return
	}
	req.SetBasicAuth(clientConfig.ClientID, clientConfig.ClientSecret)

	client := http.Client{
		Timeout: 10,
	}
	if resp, err = client.Do(req); err != nil {
		return
	}

	defer resp.Body.Close()

	var respBody []byte
	if respBody, err = ioutil.ReadAll(resp.Body); err != nil {
		return
	}

	if err = json.Unmarshal(respBody, &o2c); err != nil {
		return
	}

	return nil
}
