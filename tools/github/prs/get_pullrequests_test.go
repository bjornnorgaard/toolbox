package prs

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPullRequests(t *testing.T) {
	c := 7

	list, err := Get(
		WithLimit(uint(c)),
		WithStateClosed(),
		WithChecksSucceeded(),
		WithReviewAny(),
	)

	require.NoErrorf(t, err, "failed to get pull requests: %v", err)
	require.NotEmptyf(t, list, "expected pull requests to not be empty")
	require.Lenf(t, list, c, "expected pull requests len to be %d", c)

	for _, pr := range list {
		require.NotEmptyf(t, pr, "expected pull request to not be empty")
		require.NotEmptyf(t, pr.Author, "expected pull request author to not be empty")
		require.NotEmptyf(t, pr.CreatedAt, "expected pull request created at to not be empty")
		require.NotEmptyf(t, pr.Id, "expected pull request id to not be empty")
		require.NotEmptyf(t, pr.Number, "expected pull request number to not be empty")
		require.NotEmptyf(t, pr.Repository, "expected pull request repository to not be empty")
		require.NotEmptyf(t, pr.RepositoryWithOwner, "expected pull request repository with owner to not be empty")
		require.NotEmptyf(t, pr.State, "expected pull request state to not be empty")
		require.NotEmptyf(t, pr.Title, "expected pull request title to not be empty")
		require.NotEmptyf(t, pr.UpdatedAt, "expected pull request updated at to not be empty")
	}
}
