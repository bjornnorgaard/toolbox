package github

import (
	"fmt"
	"github.com/bjornnorgaard/toolbox/tools/github/repoedit"
	"github.com/bjornnorgaard/toolbox/tools/github/repos"
	"github.com/bjornnorgaard/toolbox/tools/github/types"
	"sync"
)

func UpdateRepos(dryRun bool) error {
	fmt.Printf("🔍 Fetching repositories...\n")
	repositories, err := repos.GetRepos()
	if err != nil {
		return err
	}
	fmt.Printf("👍 Found %d repositories\n", len(repositories))

	wg := sync.WaitGroup{}
	for _, repo := range repositories {
		wg.Add(1)
		go func(repo types.Repo) {
			defer wg.Done()

			err = repoedit.Update(repo, repoedit.WithDebug(dryRun))
			if err != nil {
				fmt.Printf("🔥 Error for '%s': %v\n", repo.FullName, err)
				return
			}

			fmt.Printf("✅ Updated %s\n", repo.FullName)
		}(repo)
	}

	wg.Wait()
	fmt.Printf("🚀 Finished updating %d repositories\n", len(repositories))
	return nil
}
