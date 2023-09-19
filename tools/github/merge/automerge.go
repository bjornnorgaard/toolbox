package merge

import (
	"fmt"
	"github.com/bjornnorgaard/toolbox/tools/github/types"
	"github.com/cli/go-gh"
)

func SetToAutoMerge(pr types.PR) error {
	_, _, err := gh.Exec("pr", "merge",
		fmt.Sprintf("%d", pr.Number),
		fmt.Sprintf("%s", autoMerge),
		fmt.Sprintf("%s", useSquash),
		fmt.Sprintf("%s", deleteBranchAfterMerge),
		fmt.Sprintf("--repo=%s", pr.RepositoryWithOwner),
	)

	return err
}
