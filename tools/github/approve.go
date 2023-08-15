package github

import (
	"fmt"
	"github.com/bjornnorgaard/toolbox/tools/github/pullrequests"
	"github.com/bjornnorgaard/toolbox/tools/github/review"
	"strings"
	"time"
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

	for i, pr := range prs {
		if err = review.ApproveSquash(pr); err != nil {
			messages := []string{
				fmt.Sprintf("❌ Failed to approve pull request #%d of %d", i+1, len(prs)),
				fmt.Sprintf("   #%s %s", pr.Title, pr.RepositoryWithOwner),
				fmt.Sprintf("   %d authored by %s created %v", pr.Number, pr.Author, pr.CreatedAt.Format(time.DateTime)),
			}

			return fmt.Errorf(strings.Join(messages, "\n"))
		}

		fmt.Printf("✅ Approved %s PR#%d '%s' created by %s\n", pr.RepositoryWithOwner, pr.Number, pr.Title, pr.Author)
	}

	fmt.Printf("🚀 Approved %d pull requests\n", len(prs))

	return nil
}
