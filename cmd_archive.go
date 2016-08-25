package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
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

	sem := make(chan bool, 5)

	client := http.Client{
		Timeout: 10 * time.Second,
	}

	req, err = http.NewRequest("GET", fmt.Sprintf(LIST_REPOSITORY_URI, username), nil)

	for {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", accessToken))

		if resp, err = client.Do(req); err != nil {
			return
		}

		defer resp.Body.Close()

		if respBody, err = ioutil.ReadAll(resp.Body); err != nil {
			return
		}

		var repoResp RepositoryResponse
		if err = json.Unmarshal(respBody, &repoResp); err != nil {
			return
		}

		for _, repo := range repoResp.Values {
			sem <- true
			go processRepository(sem, repo)
		}

		if repoResp.Next == "" {
			break
		}

		req, err = http.NewRequest("GET", repoResp.Next, nil)
	}

	return
}

func processRepository(sem chan bool, repo Repository) {
	r := fmt.Sprintf("%s/%s", repo.Owner.Username, repo.Name)

	fmt.Printf("Processing repository %s...\n", r)
	startTime := time.Now()

	if repo.Scm != "git" {
		fmt.Printf("Repository %s is not a git repository, skipping!\n", r)
		<-sem
		return
	}

	var cloneUri string
	for _, link := range repo.Links.Clone {
		if link.Name != "ssh" {
			continue
		}
		cloneUri = link.Href
	}

	if cloneUri == "" {
		fmt.Printf("Cannot find SSH link for repository %s, skipping!\n", r)
		<-sem
		return
	}

	repoOwner := repo.Owner.Username
	repoId := strings.Replace(repo.FullName, repoOwner+"/", "", -1)

	if err := cloneRepository(repoOwner, repoId, cloneUri); err != nil {
		fmt.Printf("Error processing repository %s: %s\n", r, err.Error())
		<-sem
		return
	}

	duration := time.Since(startTime)
	fmt.Printf("Finished processing repository %s (time used: %f seconds).\n", r, duration.Seconds())
	<-sem
}
