package github

import (
	"fmt"
	"sync"

	"github.com/bjornnorgaard/toolbox/tools/github/merge"
	"github.com/bjornnorgaard/toolbox/tools/github/prs"
	"github.com/bjornnorgaard/toolbox/tools/github/types"
)

func SetAutoMerge() error {
	fmt.Println("🔎 Fetching PRs to auto merge")

	list, err := prs.Get(
		prs.WithAuthorBot(),
		prs.WithChecksSucceeded(),
		prs.WithReviewNone(),
	)

	if err != nil {
		return err
	}
	if len(list) == 0 {
		fmt.Println("👍 No PRs to auto merge")
		return nil
	}

	fmt.Printf("👀 Loaded %d PRs\n", len(list))

	wg := sync.WaitGroup{}
	for _, pr := range list {
		wg.Add(1)
		go func(pr types.PR) {
			defer wg.Done()
			if err = merge.Auto(pr); err != nil {
				fmt.Printf("❗️Failed to set %s PR#%d '%s' to auto merge: %v\n", pr.Repository, pr.Number, pr.Title, err)
				return
			}
			fmt.Printf("✅ Set %s PR#%d '%s' to auto merge\n", pr.Repository, pr.Number, pr.Title)
		}(pr)
	}
	wg.Wait()

	fmt.Printf("🚀 Finished setting %d PRs to auto merge\n", len(list))
	return nil
}
