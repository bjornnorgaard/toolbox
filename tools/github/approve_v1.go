package github

import (
	"fmt"
	"log"
	"sync"
	"sync/atomic"
)

var (
	requests int32
	approved int32
)

func ApproveV1(dryRun bool) error {
	if dryRun {
		log.Println("🐵 Dry run enabled")
	}

	repos, err := getRepos()
	if err != nil {
		return err
	}

	wg := sync.WaitGroup{}

	for _, r := range repos {
		wg.Add(1)
		capturedRepo := r
		go func() {
			defer wg.Done()
			err = handlePullRequests(dryRun, capturedRepo.FullName)
			if err != nil {
				fmt.Printf("failed to handle pull-request for %s with err: %v", capturedRepo.Name, err)
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
	pullRequests, err := getPullRequestsV1(r, "author:app/dependabot")
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
			log.Printf("continuing after err: %v", err)
		}

		err = approveDependabotPRV1(r, pr.Number)
		if err != nil {
			return err
		}

		fmt.Printf("✅ Approved PR#%d %s\n", pr.Number, pr.Title)
	}

	return nil
}
