package github

import (
	"encoding/json"
	"fmt"
	"github.com/cli/go-gh"
)

// getRepositories fetches all repos owned by current gh cli user.
func getRepositories() ([]repository, error) {
	fmt.Println("🕓 Fetching list of repositories")
	buffer, _, err := gh.Exec("repo", "list", "--no-archived", "--limit", "1000", "--source", "--json", "name,owner")
	if err != nil {
		return nil, fmt.Errorf("failed to fetch list of repos - %v", err)
	}

	var repos []repository
	err = json.Unmarshal(buffer.Bytes(), &repos)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal repos - %v", err)
	}

	for i, r := range repos {
		repos[i].FullName = fmt.Sprintf("%s/%s", r.Owner.Login, r.Name)
	}

	fmt.Println("👀 Found", len(repos), "repositories")
	return repos, nil
}

// getPullRequests fetches all pull requests using passed query filter.
func getPullRequests(repo string, filter string) ([]pullRequest, error) {
	query := fmt.Sprintf("is:open is:pr %s", filter)

	buffer, _, err := gh.Exec("pr", "list", "--repo", repo, "--search", query, "--json", "number,title")
	if err != nil {
		return nil, fmt.Errorf("failed to fetch pull-requests for repo %s - %v", repo, err)
	}

	var pullRequests []pullRequest
	err = json.Unmarshal(buffer.Bytes(), &pullRequests)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal pull-requests for %s - %v", repo, err)
	}
	return pullRequests, nil
}

type repository struct {
	FullName string
	Name     string `json:"name"`
	Owner    struct {
		Id    string `json:"id"`
		Login string `json:"login"`
	} `json:"owner"`
}

type pullRequest struct {
	Number int    `json:"number"`
	Title  string `json:"title"`
}
