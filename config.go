package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func runConfig() {

	clientID := getClientID()
	clientSecret := getClientSecret()
	code := getCode(clientID)
	refreshToken := getRefreshToken(clientID, clientSecret, code)
	usernames := getTeamAndUserNames()

	c := ClientConfig{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RefreshToken: refreshToken,
		Usernames:    usernames,
	}

	saveStructToFileAsJson(c, CLIENT_CONFIG_FILE)
}

func getClientID() (clientID string) {
	fmt.Printf("Please enter your client id: ")
	fmt.Scanln(&clientID)
	return
}

func getClientSecret() (clientSecret string) {
	fmt.Printf("Please enter your client secret: ")
	fmt.Scanln(&clientSecret)
	return
}

func getCode(clientID string) (code string) {
	authUri := fmt.Sprintf(AuthorizeURI, clientID)
	fmt.Printf("Please visit this URI and finish authorization: \n%s\n", authUri)
	fmt.Printf("After authorization, please copy the code URI parameter and paste it here: ")
	fmt.Scanln(&code)
	return
}

func getRefreshToken(clientID, clientSecret, code string) (token string) {
	var (
		cbytes []byte
		req    *http.Request
		resp   *http.Response
		err    error
	)

	if cbytes, err = ioutil.ReadFile(CLIENT_CONFIG_FILE); err != nil {
		return
	}

	if err = json.Unmarshal(cbytes, &clientConfig); err != nil {
		return
	}

	values := url.Values{"grant_type": {"authorization_code"}, "code": {code}}
	valuesStr := values.Encode()
	if req, err = http.NewRequest("POST", AccessTokenURI, strings.NewReader(valuesStr)); err != nil {
		return
	}
	req.SetBasicAuth(clientID, clientSecret)
	req.Header.Add("content-type", "application/x-www-form-urlencoded")

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
	fmt.Printf("%s\n", respBody)

	var o2c Oauth2Config
	if err = json.Unmarshal(respBody, &o2c); err != nil {
		return
	}

	return o2c.RefreshToken
}

func getTeamAndUserNames() (usernames []string) {
	fmt.Printf("Now please enter the usernames/team names one by one, \n" +
		"separated by Enter, and finish with an empty line:\n")

	for {
		var username string
		fmt.Scanln(&username)

		if len(username) < 1 {
			break
		}

		usernames = append(usernames, username)
	}

	return
}
