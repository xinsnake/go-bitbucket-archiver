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

func runConfig() (err error) {

	clientID := getClientID()
	clientSecret := getClientSecret()
	code := getCode(clientID)
	refreshToken, err := getRefreshToken(clientID, clientSecret, code)
	if err != nil {
		return err
	}
	usernames := getTeamAndUserNames()

	fmt.Printf("Saving your configuration file...\n")

	c := ClientConfig{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RefreshToken: refreshToken,
		Usernames:    usernames,
	}

	return saveClientConfig(c)
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

	authUri := fmt.Sprintf(AUTHORIZE_URI, clientID)
	fmt.Printf("Please visit this URI and finish authorization: \n%s\n", authUri)
	fmt.Printf("After authorization, please copy the code URI parameter and paste it here: ")
	fmt.Scanln(&code)

	return
}

func getRefreshToken(clientID, clientSecret, code string) (string, error) {

	values := url.Values{"grant_type": {"authorization_code"}, "code": {code}}.Encode()
	req, err := http.NewRequest("POST", ACCESS_TOKEN_URI, strings.NewReader(values))

	if err != nil {
		return "", err
	}

	req.SetBasicAuth(clientID, clientSecret)
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

	return oc.RefreshToken, nil
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
