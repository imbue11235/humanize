package humanize

import "testing"

func TestSentence(t *testing.T) {
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
		if actual := Sentence(test.input); actual != test.expected {
			t.Errorf("expected `%s` but got `%s`", test.expected, actual)
		}
	}

}
