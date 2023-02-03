package github

import (
	"fmt"
	"sync"
)

func EnableAutoMerge() error {
	repos, err := getRepositories()
	if err != nil {
		return err
	}

	wg := sync.WaitGroup{}
	enabledCount := 0

	for _, repo := range repos {
		wg.Add(1)
		r := repo
		go func() {
			defer wg.Done()
			err = enableAutoMergeForRepo(r)
			if err != nil {
				fmt.Println(err)
				return
			}
			enabledCount++
		}()
	}

	wg.Wait()

	if enabledCount == len(repos) {
		fmt.Printf("🚀 Enabled auto-merge for all %d repositories\n", len(repos))
	} else {
		fmt.Printf("🚀 Enabled auto-merge for %d of %d total repositories\n", enabledCount, len(repos))
	}

	return nil
}
