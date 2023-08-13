package github

import (
	"fmt"
	"github.com/cli/go-gh"
)

func approveDependabotPRV1(repo string, pr int) error {
	_, _, err := gh.Exec("pr", "review", fmt.Sprintf("%d", pr), "--repo", repo, "--approve", "--body", "@dependabot squash and merge")
	if err != nil {
		return fmt.Errorf("failed to approve %s %d - %v", repo, pr, err)
	}
	return nil
}

const (
	dependabotSquash = "@dependabot squash and merge"
)

func approve(pr PullRequest) error {
	_, _, err := gh.Exec("pr", "review",
		fmt.Sprintf("%d", pr.Number),
		fmt.Sprintf("--repo=%s", pr.RepositoryWithOwner),
		fmt.Sprintf("--body=%s", dependabotSquash),
		"--approve")

	if err != nil {
		return err
	}

	return nil
}
