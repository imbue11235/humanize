package humanize

import (
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
	}

	for _, test := range tests {
		output := Slice(test.input, test.limit)

		if output != test.expected {
			t.Errorf("expected '%s', but got '%s'", test.expected, output)
		}
	}
}
