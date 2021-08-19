package humantime

import (
	"math"
	"time"
)

var exactThresholds = []Threshold{
	{"year", year},
	{"month", month},
	{"day", day},
	{"hour", hour},
	{"minute", minute},
	{"second", second},
}

func CalculateExactDuration(duration time.Duration) []*Result {
	delta := absDuration(duration).Seconds()
	var results []*Result

	for _, threshold := range exactThresholds {
		if delta <= 0 {
			break
		}

		// calculate amount of whole days/hours/months/units
		count := math.Floor(delta / threshold.Duration)
		delta -= count * threshold.Duration

		if count == 0 {
			continue
		}

		results = append(results, &Result{
			Threshold: threshold,
			Count: int(count),
		})
	}

	return results
}