package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func runConfig() {

	var clientID, clientSecret, code string

	fmt.Printf("Please enter your client id: ")
	fmt.Scanln(&clientID)

	fmt.Printf("Please enter your client secret: ")
	fmt.Scanln(&clientSecret)

	if len(clientID) < 10 || len(clientSecret) < 20 {
		fmt.Printf("Unable to verify client ID / secret\n")
		os.Exit(2)
	}

	authUri := fmt.Sprintf(AuthorizeURI, clientID)
	fmt.Printf("Please visit this URI and finish authorization: \n%s\n", authUri)
	fmt.Printf("After authorization, please copy the code URI parameter and paste it here: ")
	fmt.Scanln(&code)

	fmt.Printf("Now please enter the usernames/team names one by one, \n" +
		"separated by Enter, and finish with an empty line:\n")

	var usernames []string

	for {
		var username string
		fmt.Scanln(&username)

		if len(username) < 1 {
			break
		}

		usernames = append(usernames, username)
	}

	if len(usernames) == 0 {
		fmt.Printf("No usernames were specified, exising!")
		os.Exit(3)
	}

	fmt.Printf("Generating configuration file...")

	c := ClientConfig{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Code:         code,
		Usernames:    usernames,
	}

	var (
		bytes []byte
		err   error
	)

	if bytes, err = json.Marshal(c); err != nil {
		fmt.Printf(err.Error())
		os.Exit(4)
	}

	if err := ioutil.WriteFile(ClientConfigFile, bytes, 0644); err != nil {
		fmt.Printf("Unable to write file, please check the permission settings.")
		os.Exit(5)
	}
}
