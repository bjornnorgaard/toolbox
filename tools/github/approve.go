package github

import (
	"fmt"
	"github.com/bjornnorgaard/toolbox/tools/github/pullrequests"
	"github.com/bjornnorgaard/toolbox/tools/github/review"
	"sync"
)

func Approve() error {
	fmt.Println("🕓 Fetching pull requests...")
	prs, err := pullrequests.Get()
	if err != nil {
		return fmt.Errorf("🔥 Failed to fetch pull requests: %w", err)
	}
	if len(prs) == 0 {
		fmt.Println("👍 No pull requests to approve")
		return nil
	}

	fmt.Printf("👀 Loaded %d pull requests\n", len(prs))

	wg := sync.WaitGroup{}
	for _, doNotUse := range prs {
		pr := doNotUse

		wg.Add(1)
		go func() {
			defer wg.Done()

			if err = review.ApproveSquash(pr); err != nil {
				fmt.Printf("❗️Failed to approve %s PR#%d '%s': %v\n", pr.Repository, pr.Number, pr.Title, err)
				return
			}

			fmt.Printf("✅ Approved %s PR#%d '%s'\n", pr.Repository, pr.Number, pr.Title)
		}()
	}

	wg.Wait()
	fmt.Printf("🚀 Finished approving %d pull requests\n", len(prs))
	return nil
}
