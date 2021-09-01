package humantime

import (
	"math"
	"time"
)

const proximityUpperBound = 0.2

var approximationThresholds = []threshold{
	{"long", long},
	{"decade", decade},
	{"year", year},
	{"month", month},
	{"week", week},
	{"day", day},
	{"hour", hour},
	{"minute", minute},
	{"second", second},
}

// CalculateApproximateDuration ...
func CalculateApproximateDuration(duration time.Duration) *Result {
	seconds := absDuration(duration).Seconds()
	for _, threshold := range approximationThresholds {
		// calculate the proximity time distance
		proximity := seconds / threshold.duration
		next := math.Ceil(proximity)

		// if the current proximity factor is
		// close to the next whole unit of time
		// we will round it up, sort of like math.Round()
		// but with a tighter requirement than half up/half down
		// e.g.:
		// 1 hour 50 minutes => 2 hours
		// 1 hour 30 minutes => 1 hour
		if next-proximity < proximityUpperBound {
			return &Result{
				Count:  int(next),
				Symbol: threshold.symbol,
			}
		}

		if proximity >= 1 {
			return &Result{
				Count:  int(proximity),
				Symbol: threshold.symbol,
			}
		}
	}

	// if no result, it's assumed that it's less
	// than a second
	return &Result{
		Count:  1,
		Symbol: "second",
	}
}
