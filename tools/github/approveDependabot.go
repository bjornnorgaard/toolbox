package github

import (
	"fmt"
	"github.com/cli/go-gh"
	"sync"
	"sync/atomic"
)

var (
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

	wg := sync.WaitGroup{}

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
	pullRequests, err := getPullRequests(r, "-status:failure -review:approved author:app/dependabot")
	if err != nil {
		return err
	}

	if len(pullRequests) == 0 {
		return nil
	}

	atomic.AddInt32(&requests, int32(len(pullRequests)))

	for _, pr := range pullRequests {
		atomic.AddInt32(&approved, 1)

		if dryRun {
			fmt.Printf("✅ Would have approved %s PR#%d %s\n", r, pr.Number, pr.Title)
			continue
		}

		_, _, err = gh.Exec("pr", "review", fmt.Sprintf("%d", pr.Number), "--repo", r, "--approve", "--body", "@dependabot squash and merge")
		if err != nil {
			return fmt.Errorf("failed to approve %s %d %s - %v", r, pr.Number, pr.Title, err)
		}

		fmt.Println("✅ Approved PR", pr.Number, pr.Title)
	}

	return nil
}