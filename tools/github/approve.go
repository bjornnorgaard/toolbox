package github

import (
	"encoding/json"
	"fmt"
	"github.com/cli/go-gh"
	"sync"
	"sync/atomic"
)

var (
	wg       sync.WaitGroup
	requests int32
	approved int32
)

func ApproveDependabotPullRequests(dryRun bool) error {
	if dryRun {
		fmt.Println("🐵 Dry run enabled")
	}
	repos, err := getRepositories()
	if err != nil {
		return err
	}

	for _, r := range repos {
		wg.Add(1)
		repoName := r.Name
		repo := fmt.Sprintf("%s/%s", r.Owner.Login, r.Name)
		go func() {
			defer wg.Done()
			err = handlePullRequests(dryRun, repo)
			if err != nil {
				fmt.Println("failed to handle pull-request for", repoName)
			}
		}()
	}

	wg.Wait()

	if requests == 0 {
		fmt.Println("🚀 No PRs to approve")
		return nil
	}
	if approved == 0 {
		fmt.Println("🚀 Found", requests, "PRs but none could be approved")
		return nil
	}

	fmt.Println("🚀 Done approving", approved, "of", requests, "PRs from dependabot")
	return nil
}

func handlePullRequests(dryRun bool, r string) error {
	pullRequests, err := getPullRequests(r)
	if err != nil {
		return err
	}

	if len(pullRequests) == 0 {
		return nil
	}

	atomic.AddInt32(&requests, int32(len(pullRequests)))
	fmt.Println("👀 Checking", r)

	for _, pr := range pullRequests {
		atomic.AddInt32(&approved, 1)

		if dryRun {
			fmt.Println("\t✅ Would have approved PR", pr.Number, pr.Title)
			continue
		}

		_, _, err = gh.Exec("pr", "review", fmt.Sprintf("%d", pr.Number), "--repo", r, "--approve", "--body", "@dependabot squash and merge")
		if err != nil {
			return fmt.Errorf("failed to approve %s %d %s - %v", r, pr.Number, pr.Title, err)
		}

		fmt.Println("\t✅ Approved PR", pr.Number, pr.Title)
	}

	return nil
}

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

	fmt.Println("👀 Found", len(repos), "repositories")
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
