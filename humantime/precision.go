package humantime

import (
	"math"
	"time"
)

var preciseThresholds = []threshold{
	{"year", year},
	{"month", month},
	{"day", day},
	{"hour", hour},
	{"minute", minute},
	{"second", second},
}

// CalculatePreciseDuration ...
func CalculatePreciseDuration(duration time.Duration) []*Result {
	delta := absDuration(duration).Seconds()
	var results []*Result

	for _, threshold := range preciseThresholds {
		if delta <= 0 {
			break
		}

		// calculate amount of whole days/hours/months/units
		count := math.Floor(delta / threshold.duration)
		delta -= count * threshold.duration

		if count == 0 {
			continue
		}

		results = append(results, &Result{
			Symbol: threshold.symbol,
			Count:  int(count),
		})
	}

	return results
}
