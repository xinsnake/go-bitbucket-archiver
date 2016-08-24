package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	AccessTokenURI    = "https://bitbucket.org/site/oauth2/access_token"
	ListRepositoryURI = "https://api.bitbucket.org/2.0/repositories/%s"
)

var (
	clientConfig ClientConfig
)

type Repository struct {
}

func runArchive() (err error) {
	if err = o2c.getRefreshToken(); err != nil {
		return
	}

	for _, username := range clientConfig.Usernames {
		err = runArchiveUser(username)
	}

	return
}

func runArchiveUser(username string) (err error) {
	var (
		req      *http.Request
		resp     *http.Response
		respBody []byte
	)

	req, err = http.NewRequest("GET", fmt.Sprintf(ListRepositoryURI, username), nil)
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
	fmt.Printf("%s\n", respBody)

	return
}
