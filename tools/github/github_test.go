package github

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGithub(t *testing.T) {
	t.Run("run github cmd", func(t *testing.T) {
		err := Github()
		require.NoErrorf(t, err, "failed to run github cmd, received error: %v", err)
	})
}
