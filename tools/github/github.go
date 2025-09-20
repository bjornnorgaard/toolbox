package github

import (
	"fmt"

	"github.com/cli/go-gh"
)

func Github() error {
	fmt.Println("👀 Checking if GitHub CLI is correctly configured...")

	if _, _, err := gh.Exec("auth", "status"); err != nil {
		return fmt.Errorf("🔥 Something definetly went wrong: %w", err)
	}

	fmt.Println("✅ GitHub CLI is correctly configured")

	if _, _, err := gh.Exec("status"); err != nil {
		return fmt.Errorf("❌ Failed to fetch status, something might be wrong after all: %w", err)
	}

	fmt.Println("👍 Use any of the nested commands to get shit done")

	return nil
}
