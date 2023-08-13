package github

import "time"

type repository struct {
	FullName string
	Name     string `json:"name"`
	Owner    struct {
		Id    string `json:"id"`
		Login string `json:"login"`
	} `json:"owner"`
}

type PullRequestsV1 struct {
	Number int    `json:"number"`
	Title  string `json:"title"`
}

type PullRequest struct {
	Author              string    `json:"author"`
	CreatedAt           time.Time `json:"createdAt"`
	Id                  string    `json:"id"`
	Number              int       `json:"number"`
	Repository          string    `json:"repository"`
	RepositoryWithOwner string    `json:"repositoryWithOwner"`
	State               string    `json:"state"`
	Title               string    `json:"title"`
	UpdatedAt           time.Time `json:"updatedAt"`
}
