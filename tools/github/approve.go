package github

import (
	"encoding/json"
	"fmt"
	"github.com/cli/go-gh"
)

var (
	dryRun = true
)

func ApproveDependabotPullRequests() error {
	repos, err := getRepositories()
	if err != nil {
		return err
	}

	for _, r := range repos {
		repo := fmt.Sprintf("%s/%s", r.Owner.Login, r.Name)
		err = handlePullRequests(repo)
		if err != nil {
			fmt.Printf("failed to handle pull-request for %s", r.Name)
		}
	}

	fmt.Printf("\n🚀 Done approving PRs from dependabot\n")
	return nil
}

func handlePullRequests(r string) error {
	fmt.Printf("👀 Checking %s\n", r)
	pullRequests, err := getPullRequests(r)
	if err != nil {
		return err
	}

	for _, pr := range pullRequests {
		if dryRun {
			fmt.Printf("\t✅ Would have approved PR#%d: %s\n", pr.Number, pr.Title)
			return nil
		}

		_, _, err = gh.Exec("pr", "review", fmt.Sprintf("%d", pr.Number), "--repo", r, "--approve", "--body", "@dependabot squash and merge")
		if err != nil {
			return fmt.Errorf("failed to approve %s %d %s - %v", r, pr.Number, pr.Title, err)
		}

		fmt.Printf("\t✅ Approved PR#%d: %s\n", pr.Number, pr.Title)
	}

	return nil
}

func getRepositories() ([]repository, error) {
	fmt.Printf("Fetching list of repositories\n\n")
	buffer, _, err := gh.Exec("repo", "list", "--no-archived", "--limit", "10", "--source", "--json", "name,owner")
	if err != nil {
		return nil, fmt.Errorf("failed to fetch list of repos - %v", err)
	}

	var repos []repository
	err = json.Unmarshal(buffer.Bytes(), &repos)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal repos - %v", err)
	}
	return repos, nil
}

func getPullRequests(r string) ([]pullRequest, error) {
	buffer, _, err := gh.Exec("pr", "list", "--repo", r, "--search", "is:open is:pr -status:failure -review:approved author:app/dependabot", "--json", "number,title")
	if err != nil {
		return nil, fmt.Errorf("failed to fetch pull-requests for repo %s - %v", r, err)
	}

	var pullRequests []pullRequest
	err = json.Unmarshal(buffer.Bytes(), &pullRequests)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal pull-requests for %s - %v", r, err)
	}
	return pullRequests, nil
}

type repository struct {
	Name  string `json:"name"`
	Owner struct {
		Id    string `json:"id"`
		Login string `json:"login"`
	} `json:"owner"`
}

type pullRequest struct {
	Number int    `json:"number"`
	Title  string `json:"title"`
}
