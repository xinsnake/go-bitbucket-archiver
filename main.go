package main

import (
	"fmt"
	"os"
)

const (
	Oauth2ConfigFile = "config-oauth2.json"
	ClientConfigFile = "config-client.json"

	AuthorizeURI = "https://bitbucket.org/site/oauth2/authorize?client_id=%s&response_type=code"
)

type ClientConfig struct {
	ClientID     string   `json:"client_id"`
	ClientSecret string   `json:"client_secret"`
	Code         string   `json:"code"`
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

	switch args[1] {
	case "config":
		runConfig()
	case "archive":
		runArchive()
	default:
		fmt.Printf("Usage: ./bitbucket-archiver config | archive\n")
		os.Exit(1)
	}

	os.Exit(0)
}
