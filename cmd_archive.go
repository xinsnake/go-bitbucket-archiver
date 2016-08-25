package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func runArchive() (err error) {

	clientConfig, err := loadClientConfig()

	if err != nil {
		return
	}

	accessToken, err := getAccessToken(clientConfig)

	if err != nil {
		return
	}

	for _, username := range clientConfig.Usernames {
		err = runArchiveUser(accessToken, username)
	}

	return
}

func runArchiveUser(accessToken, username string) (err error) {

	var (
		req      *http.Request
		resp     *http.Response
		respBody []byte
	)

	req, err = http.NewRequest("GET", fmt.Sprintf(LIST_REPOSITORY_URI, username), nil)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", accessToken))

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

	fmt.Printf("%s\n", respBody)

	return
}
