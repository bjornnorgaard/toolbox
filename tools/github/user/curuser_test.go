package user

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetCurrentUser(t *testing.T) {
	user := Me()
	require.NotEmptyf(t, user, "expected user to not be empty")
}
