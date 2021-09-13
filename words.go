package humanize

import (
	"strings"

	"github.com/imbue11235/words"
)

// FormatWords extracts words from a fuzzy text and constructs
// a string from the words, using the provided formatter on every
// extracted word.
//
//		s:= humanize.FormatWords("some_random-sentence", func(i int, word string) {
//			return strings.ToUpper(word)
//		})
//		fmt.Print(s) => "SOME RANDOM SENTENCE"
//
func FormatWords(input string, formatter func(index int, word string) string) string {
	extractedWords := words.Extract(input)

	var sb strings.Builder
	for i, word := range extractedWords {
		if i != 0 {
			sb.WriteString(" ")
		}

		sb.WriteString(formatter(i, word))
	}

	return sb.String()
}

// Title formats a fuzzy text as a sentence in title
// case, capitalizing the first letter of every word
//
//		s := humanize.Title("some_random-sentence")
//		fmt.Print(s) => "Some Random Sentence"
//
func Title(input string) string {
	return FormatWords(input, func(index int, word string) string {
		return strings.Title(input)
	})
}

// Sentence formats a fuzzy text as a common sentence,
// capitalizing the first letter of the first word, and
// lower-casing the rest
//
//		s := humanize.Sentence("some_random-sentence")
//		fmt.Print(s) => "Some random sentence"
//
func Sentence(input string) string {
	return FormatWords(input, func(index int, word string) string {
		if index == 0 {
			return strings.Title(word)
		}

		return strings.ToLower(word)
	})
}
