package humanize

import (
	"strings"
	"testing"
)

func TestFormatFuzzyText(t *testing.T) {
	tests := []struct {
		input, expected string
		formatter       interface{}
	}{
		{
			"formatting_test",
			"FORMATTING TEST",
			strings.ToUpper,
		},
		{
			"@@@@@a.b.c!!!!1.2.3",
			"a b c 1 2 3",
			nil,
		},
		{
			"some-----text-----here",
			"Some Text Here",
			strings.Title,
		},
		{
			"1_2_3",
			"1 2 3",
			"invalid formatter",
		},
	}

	for _, test := range tests {
		if actual := FormatFuzzyText(test.input, test.formatter); actual != test.expected {
			t.Errorf("expected `%s` but got `%s`", test.expected, actual)
		}
	}
}

func TestFuzzyText(t *testing.T) {
	tests := []struct {
		input, expected string
	}{
		{"some_random-Sentence", "Some random sentence"},
		{"App_Id", "App id"},
		{"Field-name", "Field name"},
		{"PascalCase", "Pascal case"},
		{"some-!!@@----Wierd_____format", "Some wierd format"},
	}

	for _, test := range tests {
		if actual := FuzzyText(test.input); actual != test.expected {
			t.Errorf("expected `%s` but got `%s`", test.expected, actual)
		}
	}

}
