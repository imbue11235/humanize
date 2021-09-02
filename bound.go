package humanize

import (
	"fmt"
)

// Bound ...
func Bound(min, max int64) string {
	if min > max {
		return fmt.Sprintf("%d+", max)
	}

	return fmt.Sprintf("%d", max)
}
