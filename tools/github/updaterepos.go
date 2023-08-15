package github

import (
	"fmt"

	"github.com/bjornnorgaard/toolbox/tools/github/repoedit"
	"github.com/bjornnorgaard/toolbox/tools/github/repos"
)

func UpdateRepos() error {
	fmt.Printf("🔍 Fetching repos\n")
	repositories, err := repos.GetRepos()
	if err != nil {
		return err
	}

	fmt.Printf("🔧 Updating %d repos\n", len(repositories))

	for i, repo := range repositories {
		err = repoedit.Update(repo,
			repoedit.WithEnableAutoMerge(),
			repoedit.WithEnableSquashMerge(),
			repoedit.WithAllowUpdateBranch(),
			repoedit.WithDeleteBranchOnMerge())

		if err != nil {
			return fmt.Errorf("🔥 Failed to update repo '%s': %w", repo.FullName, err)
		}

		fmt.Printf("✅ Updated repo %s %d of %d\n", repo.FullName, i+1, len(repositories))
	}

	fmt.Printf("🚀 Successfully updatd %d repos\n", len(repositories))
	return nil
}
