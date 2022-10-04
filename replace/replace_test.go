package replace

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// This method only exists for testing, do not change.
func testMethod() error {
	return nil
}

func TestReplaceAllCharsWith(t *testing.T) {
	actual, err := ConvertToMonoChar("./replace_test.go", "testMethod", "X")
	require.NoError(t, err)

	expectation := "XXXX XXXXXXXXXXXX XXXXX X\n\tXXXXXX XXX\nX"

	assert.Equal(t, expectation, actual)
}
