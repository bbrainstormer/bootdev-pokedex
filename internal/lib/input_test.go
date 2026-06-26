package lib

import (
	"reflect"
	"testing"
)

func TestCleanInput(t *testing.T) {
	type testCase struct {
		input    string
		expected []string
	}

	cases := []testCase{
		{
			input:    "hello world",
			expected: []string{"hello", "world"},
		},
		{
			input:    "line\nbreak",
			expected: []string{"line", "break"},
		},
		{
			input:    " leading trailing   ",
			expected: []string{"leading", "trailing"},
		},
		{
			input:    "",
			expected: []string{},
		},
		{
			input:    "",
			expected: []string{},
		},
		{
			input:    "UPPER",
			expected: []string{"upper"},
		},
	}

	for _, c := range cases {
		t.Run("\""+c.input+"\"", func(t *testing.T) {
			output := CleanInput(c.input)
			if !reflect.DeepEqual(c.expected, output) {
				t.Fatalf("expected: %v, got: %v", c.expected, output)
			}
		})
	}
}
