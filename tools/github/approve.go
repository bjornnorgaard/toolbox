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

	var prOpts []prs.OptsApply
	if force {
		// When force is enabled, include already approved PRs
		prOpts = append(prOpts, prs.WithReviewApproved())
	}

	pulls, err := prs.Get(prOpts...)
	if err != nil {
		return fmt.Errorf("🔥 Failed to fetch pull requests: %w", err)
	}

	if len(pulls) == 0 {
		fmt.Println("👍 No pull requests to approve")
		return nil
	}

	fmt.Printf("👀 Loaded %d pull requests\n", len(pulls))

	wg := sync.WaitGroup{}
	for _, pull := range pulls {
		wg.Add(1)
		go func(pr types.PR) {
			defer wg.Done()
			var reviewOpt review.ApplyOpts
			reviewOpt = review.WithSquash()
			if err = review.ApproveSquash(pr, reviewOpt); err != nil {
				fmt.Printf("❗️Failed to approve %s PR#%d '%s': %v\n", pr.Repository, pr.Number, pr.Title, err)
				return
			}
			fmt.Printf("✅ Approved %s PR#%d '%s'\n", pr.Repository, pr.Number, pr.Title)
		}(pull)
	}

	wg.Wait()
	fmt.Printf("🚀 Finished approving %d pull requests\n", len(pulls))
	return nil
}
