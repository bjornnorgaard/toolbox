package github

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetRepos(t *testing.T) {
	repos, err := getRepos()
	require.NoErrorf(t, err, "failed to get repos: %v", err)
	assert.NotEmptyf(t, repos, "expected repos to not be empty")
	assert.Greaterf(t, len(repos), 50, "expected repos to be greater than 50")
}
