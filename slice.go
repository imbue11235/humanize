package humanize

import (
	"fmt"
	"strings"
)

func Slice(items []string) string {
	if len(items) == 0 {
		return ""
	}

	// pop last element from output
	lastItem, items := items[len(items)-1], items[:len(items)-1]

	if len(items) == 0 {
		return lastItem
	}

	// construct the output string with format as "Joe, Joseph and Joanna"
	return fmt.Sprintf("%s %s %s", strings.Join(items, ", "), currentLocale.Common.Connector, lastItem)
}