package humanize

import (
	"fmt"
	"strings"
)

// Slice takes a slice of items and returns a string
// representation of that slice.
// Example:
//
//		s := humanize.Slice([]string{"One", "Two", "Three"})
//		fmt.Print(s) => "One, Two and Three"
//
// Optionally a limit can be given.
// Example:
//
//		s := humanize.Slice([]string{"One", "Two", "Three"}, 2)
// 		fmt.Print(s) => "One, Two and one other"
//
func Slice(items []string, limits ...uint) string {
	if len(items) == 0 {
		return ""
	}

	// if only one item exists,
	// we should not continue with the formatting,
	// and just return the single item instead
	if len(items) == 1 {
		return items[0]
	}

	// if a limit is set and that limit does not exceed the length of the items,
	// we will format the slice with a limit
	if len(limits) > 0 && limits[0] > 0 && int(limits[0]) <= len(items) {
		return formatSliceWithLimit(items, limits[0])
	}

	// otherwise, format slice normally
	return formatSlice(items)
}

// splitSlice splits a slice into two at a given position
func splitSlice(items []string, at int) ([]string, []string) {
	if len(items) <= at {
		return items, []string{}
	}

	return items[:at], items[at:]
}

// formatSliceWithLimit converts a slice of items into a string
// representation, but limited to the given limit as a positive int
// e.g.: [A, B, C, D, E] with limit 2 => "A, B and 3 others"
func formatSliceWithLimit(items []string, limit uint) string {
	items, rest := splitSlice(items, int(limit))

	return fmt.Sprintf(
		"%s %s %s",
		strings.Join(items, ", "),
		translate("slice.connector"),
		pluralize("slice.rest", len(rest)),
	)
}

// formatSlice converts a slice of items into a string
// representation, separated by commas
// e.g.: [A, B, C] => "A, B and C"
func formatSlice(items []string) string {
	items, rest := splitSlice(items, len(items)-1)

	return fmt.Sprintf("%s %s %s", strings.Join(items, ", "), translate("slice.connector"), rest[0])
}
