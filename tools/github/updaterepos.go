package github

import (
	"fmt"
	"github.com/bjornnorgaard/toolbox/tools/github/repoedit"
	"github.com/bjornnorgaard/toolbox/tools/github/repos"
)

func UpdateRepos() error {
	repositories, err := repos.GetRepos()
	if err != nil {
		return err
	}

	for _, repo := range repositories {
		err = repoedit.Update(repo,
			repoedit.WithEnableAutoMerge(),
			repoedit.WithEnableSquashMerge(),
			repoedit.WithDeleteBranchOnMerge())

		if err != nil {
			return fmt.Errorf("failed to update repo '%s': %w", repo.FullName, err)
		}
	}

	return nil
}
