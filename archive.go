package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

const (
	AccessTokenURI    = "https://bitbucket.org/site/oauth2/access_token"
	AccessTokenBody   = "grant_type=authorization_code&code=%s"
	ListRepositoryURI = "https://api.bitbucket.org/2.0/repositories/%s"
)

var (
	clientConfig ClientConfig
	o2c          = Oauth2Config{}
)

type Repository struct {
}

func runArchive() (err error) {
	var (
		req      *http.Request
		resp     *http.Response
		respBody []byte
	)

	if err = o2c.getRefreshToken(); err != nil {
		return
	}

	req, err = http.NewRequest("GET", fmt.Sprintf(ListRepositoryURI, clientConfig.Usernames[0]), nil)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", o2c.AccessToken))

	client := http.Client{
		Timeout: 10 * time.Second,
	}
	if resp, err = client.Do(req); err != nil {
		return
	}

	defer resp.Body.Close()

	if respBody, err = ioutil.ReadAll(resp.Body); err != nil {
		return
	}
	fmt.Printf("%s, %s\n", resp.Status, respBody)

	return
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
	// TODO here is the problem
	req.SetBasicAuth(clientConfig.ClientID, clientConfig.ClientSecret)

	client := http.Client{
		Timeout: 10 * time.Second,
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
