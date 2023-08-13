package repos

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetRepos(t *testing.T) {
	c := 6

	repos, err := GetRepos(withLimit(uint(c)))
	require.NoErrorf(t, err, "failed to get repos: %v", err)
	require.NotEmptyf(t, repos, "expected repos to not be empty")
	assert.Lenf(t, repos, c, "expected repos len to be %d", c)

	for _, repo := range repos {
		require.NotEmptyf(t, repo, "expected repo to not be empty")
		require.NotEmptyf(t, repo.CreatedAt, "expected repo created at to not be empty")
		require.NotNilf(t, repo.Description, "expected repo description to not be nil")
		require.NotEmptyf(t, repo.FullName, "expected repo full name to not be empty")
		require.NotEmptyf(t, repo.ID, "expected repo id to not be empty")
		require.NotEmptyf(t, repo.Name, "expected repo name to not be empty")
		require.NotEmptyf(t, repo.Owner, "expected repo owner to not be empty")
		require.NotEmptyf(t, repo.OwnerID, "expected repo owner id to not be empty")
		require.NotEmptyf(t, repo.URL, "expected repo url to not be empty")
		require.NotEmptyf(t, repo.Visibility, "expected repo visibility to not be empty")
	}
}
