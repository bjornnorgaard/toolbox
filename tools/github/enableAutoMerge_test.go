package github

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestEnableAutoMerge(t *testing.T) {
	err := EnableAutoMerge()
	require.NoError(t, err)
}
