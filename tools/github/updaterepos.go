package github

import (
	"fmt"
	"sync"

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

	wg := sync.WaitGroup{}
	errCh := make(chan error, len(repositories))

	for i, r := range repositories {
		var (
			number = i + 1
			repo   = r
		)

		go func() {
			wg.Add(1)
			defer wg.Done()

			err = repoedit.Update(repo,
				repoedit.WithEnableAutoMerge(),
				repoedit.WithEnableSquashMerge(),
				repoedit.WithShowUpdateBranch(),
				repoedit.WithDeleteBranchOnMerge())

			if err != nil {
				err = fmt.Errorf("🔥 Failed to update repo '%s': %w", repo.FullName, err)
				errCh <- err
				return
			}

			fmt.Printf("✅ Updated repo %s %d of %d\n", repo.FullName, number, len(repositories))
		}()
	}

	wg.Wait()

	close(errCh)
	if 0 < len(errCh) {
		for err = range errCh {
			fmt.Println(err)
		}
		return fmt.Errorf("💀 Failed to update %d repos", len(errCh))
	}

	fmt.Printf("🚀 Successfully updated %d repos\n", len(repositories))
	return nil
}
