package user

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetCurrentUser(t *testing.T) {
	user := Me()
	require.NotEmptyf(t, user, "expected user to not be empty")
}
