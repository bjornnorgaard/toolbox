package github

import (
	"fmt"
	"github.com/bjornnorgaard/toolbox/tools/github/repoedit"
	"github.com/bjornnorgaard/toolbox/tools/github/repos"
	"sync"
	"sync/atomic"
)

func UpdateRepos(dryRun bool) error {
	fmt.Printf("🔍 Fetching repositories...\n")
	repositories, err := repos.GetRepos()
	if err != nil {
		return err
	}

	fmt.Printf("👍 Found %d repositories\n", len(repositories))

	wg := sync.WaitGroup{}
	successCount := uint32(len(repositories))

	for _, doNotUseRepo := range repositories {
		repo := doNotUseRepo

		wg.Add(1)
		go func() {
			defer wg.Done()

			err = repoedit.Update(repo, repoedit.WithDebug(dryRun))
			if err != nil {
				fmt.Printf("🔥 Error for '%s': %v\n", repo.FullName, err)
				atomic.AddUint32(&successCount, -1)
				return
			}

			fmt.Printf("✅ Updated %s\n", repo.FullName)
		}()
	}

	wg.Wait()

	if int(successCount) != len(repositories) {
		fmt.Printf("🔥 Failed to update %d repos\n", len(repositories)-int(successCount))
	}
	if 0 < successCount {
		fmt.Printf("🚀 Updated %d repos\n", successCount)
	}

	return nil
}
