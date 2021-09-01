package humanize

import (
	"strings"
	"testing"
)

func TestSlice(t *testing.T) {
	tests := []struct {
		input    []string
		expected string
		limit    uint
	}{
		{
			input:    []string{},
			expected: "",
		},
		{
			input:    []string{"Joachim"},
			expected: "Joachim",
		},
		{
			input:    []string{"Joe", "Jennifer", "Anna"},
			expected: "Joe, Jennifer and Anna",
		},
		{
			input:    []string{"Hans", "Grethel"},
			expected: "Hans and Grethel",
		},
		{
			input:    []string{"Hans", "Joe", "Rob"},
			limit:    2,
			expected: "Hans, Joe and one other",
		},
		{
			input:    []string{"Zanday", "Toumas", "Beate", "Nicolaj"},
			limit:    2,
			expected: "Zanday, Toumas and 2 others",
		},
		{
			input:    []string{"George"},
			limit:    3,
			expected: "George",
		},
		{
			input:    []string{"Donald", "Goofy"},
			limit:    4,
			expected: "Donald and Goofy",
		},
		{
			input:    []string{"A", "B", "C"},
			limit:    1,
			expected: "A and 2 others",
		},
		{
			input:    strings.Split("ABCDEFGHIJKLMNOPQRSTUVXYZ", ""),
			limit:    5,
			expected: "A, B, C, D, E and 20 others",
		},
	}

	for _, test := range tests {
		output := Slice(test.input, test.limit)
		if output != test.expected {
			t.Errorf("expected '%s', but got '%s'", test.expected, output)
		}
	}
}