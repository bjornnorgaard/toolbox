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
	errCh := make(chan error, len(prs))

	for _, pr := range prs {
		capPR := pr

		go func() {
			wg.Add(1)
			defer wg.Done()

			if err = review.ApproveSquash(pr); err != nil {
				errCh <- fmt.Errorf("❗️Failed to approve %s PR#%d '%s'", capPR.Repository, capPR.Number, capPR.Title)
				return
			}

			fmt.Printf("✅ Approved %s PR#%d '%s' created by %s\n", pr.RepositoryWithOwner, pr.Number, pr.Title, pr.Author)
		}()
	}

	wg.Wait()

	close(errCh)
	if 0 < len(errCh) {
		for err = range errCh {
			fmt.Println(err)
		}
		return fmt.Errorf("💀 Failed to approve %d pull requests", len(errCh))
	}

	fmt.Printf("🚀 Approved %d pull requests\n", len(prs))
	return nil
}
