package main

import (
	"errors"
	"fmt"
	"os"
)

const (
	CLIENT_CONFIG_FILE = "config-client.json"

	ACCESS_TOKEN_URI    = "https://bitbucket.org/site/oauth2/access_token"
	AUTHORIZE_URI       = "https://bitbucket.org/site/oauth2/authorize?client_id=%s&response_type=code"
	LIST_REPOSITORY_URI = "https://api.bitbucket.org/2.0/repositories/%s"
)

type ClientConfig struct {
	ClientID     string   `json:"client_id"`
	ClientSecret string   `json:"client_secret"`
	RefreshToken string   `json:"refresh_token"`
	Usernames    []string `json:"usernames"`
}

type Oauth2Config struct {
	AccessToken  string `json:"access_token"`
	Scopes       string `json:"scopes"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
}

func main() {

	args := os.Args

	if len(args) != 2 || (args[1] != "config" && args[1] != "archive") {
		fmt.Printf("Usage: ./bitbucket-archiver config | archive\n")
		os.Exit(1)
	}

	var err error

	switch args[1] {
	case "config":
		err = runConfig()
	case "archive":
		err = runArchive()
	default:
		err = errors.New("Usage: ./bitbucket-archiver config | archive\n")
	}

	if err != nil {
		fmt.Printf("ERR: %s\n", err)
		os.Exit(1)
	}

	os.Exit(0)
}
