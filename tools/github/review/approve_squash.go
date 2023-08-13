package review

import (
	"fmt"
	"github.com/bjornnorgaard/toolbox/tools/github/types"
	"github.com/cli/go-gh"
)

const (
	dependabotSquash = "@dependabot squash and merge"
)

func ApproveSquash(pr types.PR) error {
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
