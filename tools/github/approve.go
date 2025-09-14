package github

import (
	"fmt"
	"sync"

	"github.com/bjornnorgaard/toolbox/tools/github/prs"
	"github.com/bjornnorgaard/toolbox/tools/github/review"
	"github.com/bjornnorgaard/toolbox/tools/github/types"
)

func Approve(force bool) error {
	fmt.Println("🕓 Fetching pull requests to approve...")

	// Get PRs with successful checks
	var successOpts []prs.OptsApply
	if force {
		// When force is enabled, include already approved PRs
		successOpts = append(successOpts, prs.WithReviewApproved())
	}

	successPulls, err := prs.Get(successOpts...)
	if err != nil {
		return fmt.Errorf("🔥 Failed to fetch successful pull requests: %w", err)
	}

	// Get PRs with failed checks
	var failedOpts []prs.OptsApply
	failedOpts = append(failedOpts, prs.WithChecksFailed())
	if force {
		failedOpts = append(failedOpts, prs.WithReviewApproved())
	}

	failedPulls, err := prs.Get(failedOpts...)
	if err != nil {
		return fmt.Errorf("🔥 Failed to fetch failed pull requests: %w", err)
	}

	totalPulls := len(successPulls) + len(failedPulls)
	if totalPulls == 0 {
		fmt.Println("👍 No pull requests to approve")
		return nil
	}

	fmt.Printf("👀 Loaded %d pull requests (%d successful, %d failed)\n", totalPulls, len(successPulls), len(failedPulls))

	wg := sync.WaitGroup{}

	// Process successful PRs with squash and merge
	for _, pull := range successPulls {
		wg.Add(1)
		go func(pr types.PR) {
			defer wg.Done()
			if err = review.ApproveSquash(pr, review.WithSquash()); err != nil {
				fmt.Printf("❗️Failed to approve %s PR#%d '%s': %v\n", pr.Repository, pr.Number, pr.Title, err)
				return
			}
			fmt.Printf("✅ Approved %s PR#%d '%s' (squash and merge)\n", pr.Repository, pr.Number, pr.Title)
		}(pull)
	}

	// Process failed PRs with recreate
	for _, pull := range failedPulls {
		wg.Add(1)
		go func(pr types.PR) {
			defer wg.Done()
			if err = review.ApproveSquash(pr, review.WithRecreate()); err != nil {
				fmt.Printf("❗️Failed to approve %s PR#%d '%s': %v\n", pr.Repository, pr.Number, pr.Title, err)
				return
			}
			fmt.Printf("✅ Approved %s PR#%d '%s' (recreate)\n", pr.Repository, pr.Number, pr.Title)
		}(pull)
	}

	wg.Wait()
	fmt.Printf("🚀 Finished approving %d pull requests\n", totalPulls)
	return nil
}
