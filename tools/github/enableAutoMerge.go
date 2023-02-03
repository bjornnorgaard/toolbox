package github

import (
	"fmt"
	"github.com/cli/go-gh"
	"sync"
)

func EnableAutoMerge() error {
	repos, err := getRepositories()
	if err != nil {
		return err
	}

	wg := sync.WaitGroup{}

	for _, repo := range repos {
		wg.Add(1)
		r := repo
		go func() {
			defer wg.Done()
			enableAutoMergeForRepo(r)
		}()
	}

	wg.Wait()

	return nil
}

func enableAutoMergeForRepo(r repository) {
	_, _, err := gh.Exec("repo", "edit", r.FullName, "--enable-auto-merge")
	if err != nil {
		fmt.Println("❌ Failed to enabled auto-merge for", r.FullName, "with error:", err)
		return
	}

	fmt.Printf("✅ Enabled auto-merge for %s\n", r.FullName)
}
