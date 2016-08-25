package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

type RepositoryResponse struct {
	Pagelen int          `json:"pagelen"`
	Size    int          `json:"size"`
	Values  []Repository `json:"values"`
	Page    int          `json:"page"`
	Next    string       `json:"next"`
}

type Repository struct {
	Scm     string `json:"scm"`
	Website string `json:"website"`
	HasWiki bool   `json:"has_wiki"`
	Name    string `json:"name"`
	Links   struct {
		Watchers struct {
			Href string `json:"href"`
		} `json:"watchers"`
		Branches struct {
			Href string `json:"href"`
		} `json:"branches"`
		Tags struct {
			Href string `json:"href"`
		} `json:"tags"`
		Commits struct {
			Href string `json:"href"`
		} `json:"commits"`
		Clone []struct {
			Href string `json:"href"`
			Name string `json:"name"`
		} `json:"clone"`
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
		HTML struct {
			Href string `json:"href"`
		} `json:"html"`
		Avatar struct {
			Href string `json:"href"`
		} `json:"avatar"`
		Hooks struct {
			Href string `json:"href"`
		} `json:"hooks"`
		Forks struct {
			Href string `json:"href"`
		} `json:"forks"`
		Downloads struct {
			Href string `json:"href"`
		} `json:"downloads"`
		Pullrequests struct {
			Href string `json:"href"`
		} `json:"pullrequests"`
	} `json:"links"`
	ForkPolicy string `json:"fork_policy"`
	UUID       string `json:"uuid"`
	Project    struct {
		Key   string `json:"key"`
		Type  string `json:"type"`
		UUID  string `json:"uuid"`
		Links struct {
			Self struct {
				Href string `json:"href"`
			} `json:"self"`
			HTML struct {
				Href string `json:"href"`
			} `json:"html"`
			Avatar struct {
				Href string `json:"href"`
			} `json:"avatar"`
		} `json:"links"`
		Name string `json:"name"`
	} `json:"project"`
	Language  string    `json:"language"`
	CreatedOn time.Time `json:"created_on"`
	FullName  string    `json:"full_name"`
	HasIssues bool      `json:"has_issues"`
	Owner     struct {
		Username    string `json:"username"`
		DisplayName string `json:"display_name"`
		Type        string `json:"type"`
		UUID        string `json:"uuid"`
		Links       struct {
			Self struct {
				Href string `json:"href"`
			} `json:"self"`
			HTML struct {
				Href string `json:"href"`
			} `json:"html"`
			Avatar struct {
				Href string `json:"href"`
			} `json:"avatar"`
		} `json:"links"`
	} `json:"owner"`
	UpdatedOn   time.Time `json:"updated_on"`
	Size        int       `json:"size"`
	Type        string    `json:"type"`
	IsPrivate   bool      `json:"is_private"`
	Description string    `json:"description"`
}

func cloneRepository(owner, id, uri string) (err error) {
	ownerDir := fmt.Sprintf("%s/%s", REPO_ARCHIVE_DIR, owner)

	if _, err = os.Stat(ownerDir); os.IsNotExist(err) {
		os.MkdirAll(ownerDir, os.FileMode(int(0755)))
	}

	repoDir := fmt.Sprintf("%s/%s", ownerDir, id)

	var cmd *exec.Cmd

	_, err = os.Stat(repoDir)
	if err == nil {
		fmt.Printf("Updating repository %s/%s\n", owner, id)
		cmd = exec.Command("git", "pull")
		cmd.Dir = repoDir
	} else if os.IsNotExist(err) {
		fmt.Printf("Cloning repository %s/%s\n", owner, id)
		cmd = exec.Command("git", "clone", uri)
		cmd.Dir = ownerDir
	} else {
		return err
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}
