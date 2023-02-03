package github

import (
	"fmt"
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
				fmt.Printf("failed to handle pull-request for %s with err: %v", repoName, err)
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

		err = setToAutoMerge(r, pr)
		if err != nil {
			return err
		}

		err = commentOnPR(r, pr.Number, "Set to merge automatically when requirements are meet.")
		if err != nil {
			return err
		}

		err = approveDependabotPR(r, pr.Number)
		if err != nil {
			return err
		}

		fmt.Printf("✅ Approved PR#%d %s\n", pr.Number, pr.Title)
	}

	return nil
}
