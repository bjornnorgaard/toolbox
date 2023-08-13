package github

import (
	"encoding/json"
	"fmt"
	"github.com/cli/go-gh"
)

// getRepos fetches all repos owned by current gh cli user.
func getRepos() ([]repository, error) {
	fmt.Println("🕓 Fetching list of repositories")
	buffer, _, err := gh.Exec("repo", "list", "--no-archived", "--limit", "1000", "--source", "--json", "name,owner")
	if err != nil {
		return nil, fmt.Errorf("failed to fetch list of repos - %v", err)
	}

	var repos []repository
	err = json.Unmarshal(buffer.Bytes(), &repos)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal repos - %v", err)
	}

	for i, r := range repos {
		repos[i].FullName = fmt.Sprintf("%s/%s", r.Owner.Login, r.Name)
	}

	fmt.Println("👀 Found", len(repos), "repositories")
	return repos, nil
}
