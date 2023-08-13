package github

import (
	"fmt"
	"github.com/cli/go-gh"
	"strconv"
)

func enableAutoMergeForRepo(r repository) error {
	_, _, err := gh.Exec("repo", "edit", r.FullName, "--enable-auto-merge")
	if err != nil {
		return fmt.Errorf("❌ Failed to enabled auto-merge for %s, with error: %v\n", r.FullName, err)
	}
	return nil
}

func commentOnPR(repo string, pr int, msg string) error {
	_, _, err := gh.Exec("pr", "comment", strconv.Itoa(pr), "--repo", repo, "--body", msg)
	if err != nil {
		return fmt.Errorf("failed to comment on PR %d after settings auto-merge: %v", pr, err)
	}
	return nil
}

func setToAutoMerge(r string, pr PullRequestsV1) error {
	_, _, err := gh.Exec("pr", "merge", strconv.Itoa(pr.Number), "--auto", "--squash", "--delete-branch", "--repo", r)
	if err != nil {
		return fmt.Errorf("failed to set auto merge for repo %s with err: %v", r, err)
	}

	err = commentOnPR(r, pr.Number, "Set to merge automatically when requirements are meet.")
	if err != nil {
		return fmt.Errorf("failed to comment on PR after setting auto-merge: %v", err)
	}
	return nil
}
