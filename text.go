package humanize

import (
	"strings"

	"github.com/imbue11235/words"
)

// FuzzyText formats a fuzzy text as a common sentence,
// capitalizing the first letter of the first word, and
// lower-casing the rest
//
//		s := humanize.FuzzyText("some_random-sentence")
//		fmt.Print(s) => "Some random sentence"
//
func FuzzyText(input string) string {
	return FormatFuzzyText(input, func(index int, word string) string {
		if index == 0 {
			return strings.Title(word)
		}

		return strings.ToLower(word)
	})
}

// FormatFuzzyText extracts words from a fuzzy text and constructs
// a string from the words, using the provided formatter on every
// extracted word.
//
//		s:= humanize.FormatFuzzyText("some_random-sentence", strings.ToUpper)
//		fmt.Print(s) => "SOME RANDOM SENTENCE"
//
func FormatFuzzyText(input string, formatter interface{}) string {
	extractedWords := words.Extract(input)

	var sb strings.Builder
	for i, word := range extractedWords {
		if i != 0 {
			sb.WriteString(" ")
		}

		if formatter == nil {
			sb.WriteString(word)
			continue
		}

		sb.WriteString(useFormatter(formatter, word, i))
	}

	return sb.String()
}

// useFormatter tries to cast the formatter to one of the two function definitions accepted
// and calls it with given parameters. if the casting is unsuccessful, the parameter string value
// is returned
func useFormatter(formatter interface{}, value string, index int) string {
	if formatFunc, ok := formatter.(func(int, string) string); ok {
		return formatFunc(index, value)
	}

	if formatFunc, ok := formatter.(func(string) string); ok {
		return formatFunc(value)
	}

	return value
}
