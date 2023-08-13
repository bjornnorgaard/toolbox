package github

import (
	"fmt"
	"log"
)

func Approve(dry bool) error {
	if dry {
		log.Printf("🐵 Dry run enabled")
	}

	log.Printf("🕓 Fetching pull requests...")
	prs, err := getprs()
	if err != nil {
		return fmt.Errorf("failed to fetch pull requests: %w", err)
	}

	if len(prs) == 0 {
		log.Printf("👍 No pull requests to approve")
		return nil
	}

	log.Printf("👀 Loaded %d pull requests", len(prs))

	for _, pr := range prs {
		if err = approve(pr); err != nil {
			return fmt.Errorf("failed to approve pull request: %w", err)
		}

		log.Printf("✅ Approved %s PR#%d '%s' created by %s", pr.RepositoryWithOwner, pr.Number, pr.Title, pr.Author)
	}

	log.Printf("🚀 Approved %d pull requests", len(prs))
	return nil
}
