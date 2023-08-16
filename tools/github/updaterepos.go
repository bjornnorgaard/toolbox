package github

import (
	"fmt"
	"github.com/bjornnorgaard/toolbox/tools/github/repoedit"
	"github.com/bjornnorgaard/toolbox/tools/github/repos"
	"sync"
)

func UpdateRepos() error {
	fmt.Printf("🔍 Fetching repos\n")
	repositories, err := repos.GetRepos()
	if err != nil {
		return err
	}

	fmt.Printf("🔧 Updating %d repos\n", len(repositories))

	wg := sync.WaitGroup{}
	errCh := make(chan error, len(repositories)*2)

	for _, doNotUseRepo := range repositories {
		repo := doNotUseRepo

		go func() {
			wg.Add(1)
			defer wg.Done()

			if updateErr := repoedit.Update(repo); updateErr != nil {
				updateErr = fmt.Errorf("🔥 Failed to update '%s': %w", repo.FullName, updateErr)
				fmt.Println(updateErr)
				errCh <- updateErr
				return
			}

			fmt.Printf("✅ Updated repo %s\n", repo.FullName)
		}()
	}

	wg.Wait()
	close(errCh)

	if _, hasErrors := <-errCh; hasErrors {
		for err = range errCh {
			fmt.Println(err)
		}
		return fmt.Errorf("💀 Failed to update one or more repos")
	}

	fmt.Printf("🚀 Successfully updated %d repos\n", len(repositories))
	return nil
}
