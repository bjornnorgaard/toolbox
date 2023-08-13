package jqexp

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestReplacer(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"empty string", "", ""},
		{"whitespace", " ", " "},
		{"newline", "hello \nworld", "hello world"},
		{"tab", "hello \tworld", "hello world"},
		{"both", "hello \n\tworld", "hello world"},
		{"many mixed", "\n\thello \n\tworld\n", "hello world"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := New(tt.input)
			require.Equalf(t, tt.expected, out, "expected func(‰s) = %s, got %s", tt.input, tt.expected, out)
		})
	}
}
