package humanize

import (
	"fmt"
	"strings"

	"github.com/imbue11235/humanize/language"
)

func Slice(items []string, limits ...uint) string {
	if len(items) == 0 {
		return ""
	}

	// if only one item exists
	// we should not continue with the formatting
	// and just return the single item instead
	if len(items) == 1 {
		return items[0]
	}

	// get translations
	translations := manager.Locale().Slice

	if len(limits) > 0 && limits[0] > 0 {
		return formatSliceWithLimit(items, limits[0], translations)
	}

	return formatSlice(items, translations.Connector)
}

func splitSlice(items []string, at int) ([]string, []string) {
	if len(items) <= at {
		return items, []string{}
	}

	return items[:at], items[at:]
}

func formatSliceWithLimit(items []string, limit uint, translations language.Slice) string {
	items, rest := splitSlice(items, int(limit))

	return fmt.Sprintf(
		"%s %s %s",
		strings.Join(items, ", "),
		translations.Connector,
		translations.Rest.Pluralize(len(rest)),
	)
}

func formatSlice(items []string, connector string) string {
	items, rest := splitSlice(items, len(items)-1)

	return fmt.Sprintf("%s %s %s", strings.Join(items, ", "), connector, rest[0])
}
