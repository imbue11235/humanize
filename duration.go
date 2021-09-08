package humanize

import (
	"fmt"
	"time"

	"github.com/imbue11235/humanize/era"
)

// Duration ...
func Duration(duration time.Duration) string {
	// find the closest estimated "time distance" from given difference
	estimation := era.DurationToEstimation(duration)

	path := fmt.Sprintf("time.estimation.%s", estimation.Symbol)
	return pluralize(path, estimation.Volume)
}

// ExactDuration ...
func ExactDuration(duration time.Duration) string {
	return concatResults("time.precision", era.DurationToPreciseTimeUnits(duration))
}

func concatResults(path string, results []*era.Result) string {
	var output []string
	for _, result := range results {
		if result.Volume == 0 {
			continue
		}

		path := fmt.Sprintf("%s.%s", path, result.Symbol)
		output = append(output, pluralize(path, result.Volume))
	}

	return Slice(output)
}
