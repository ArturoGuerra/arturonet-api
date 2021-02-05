package projects

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo/v4"
)

const githubAPIBase = "https://api.github.com"
const githubUser = "ArturoGuerra"

var githubOrgs = []string{
	"destinyarena",
}

type (
	// RepoOwner is a struct representation of a github user that owns a repo
	RepoOwner struct {
		Login     string `json:"login"`
		ID        int    `json:"id"`
		NodeID    string `json:"node_id"`
		AvatarURL string `json:"avatar_url"`
		URL       string `json:"url"`
		Type      string `json:"type"`
	}

	// Repo is a struct representation of a github repo
	Repo struct {
		ID          int        `json:"id"`
		NodeID      string     `json:"node_id"`
		Name        string     `json:"name"`
		FullName    string     `json:"full_name"`
		Private     bool       `json:"private"`
		Owner       *RepoOwner `json:"owner"`
		Description string     `json:"description"`
		Fork        bool       `json:"fork"`
		URL         string     `json:"url"`
		HTMLURL     string     `json:"html_url"`
		Language    string     `json:"language"`
		ForksCount  int        `json:"forks_count"`
		Archived    bool       `json:"archived"`
		Watchers    int        `json:"watchers_count"`
		Stars       int        `json:"stargazers_count"`
		CreatedAt   string     `json:"created_at"`
		UpdatedAt   string     `json:"updated_at"`
	}
)

func (p *projects) getRepos(url string) ([]*Repo, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	data := make([]*Repo, 0)
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	return data, nil
}

func (p *projects) getProjects() echo.HandlerFunc {
	return func(c echo.Context) error {
		// Repos that will be turned in json for return payload
		repos := make([]*Repo, 0)

		userurl := fmt.Sprintf("%s/user/%s/repos", githubAPIBase, githubUser)

		urls := []string{
			userurl,
		}

		for _, v := range githubOrgs {
			url := fmt.Sprintf("%s/orgs/%s/repos", githubAPIBase, v)
			urls = append(urls, url)
		}

		for _, url := range urls {
			newrepos, err := p.getRepos(url)
			if err == nil {
				repos = append(repos, newrepos...)
			} else {
				p.Logger.Error(err)
			}
		}

		return c.JSON(200, repos)
	}
}
