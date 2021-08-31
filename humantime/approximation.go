package humantime

import (
	"math"
	"time"
)

var approximationThresholds = []Threshold{
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
		proximity := math.Floor(seconds / threshold.Duration)
		if proximity == 0 {
			continue
		}

		return &Result{
			Count:     int(proximity),
			Threshold: threshold,
		}
	}

	return nil
}
