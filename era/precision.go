package era

import (
	"math"
	"time"
)

var preciseThresholds = []threshold{
	{SymbolYear, year},
	{SymbolMonth, month},
	{SymbolDay, day},
	{SymbolHour, hour},
	{SymbolMinute, minute},
	{SymbolSecond, second},
}

// CalculatePreciseDuration ...
func DurationToPreciseTimeUnits(duration time.Duration) []*Result {
	delta := absDuration(duration).Seconds()
	var results []*Result

	for _, threshold := range preciseThresholds {
		if delta <= 0 {
			break
		}

		// calculate amount of whole days/hours/months/units
		volume := math.Floor(delta / threshold.duration.Seconds())
		delta -= volume * threshold.duration.Seconds()

		if volume == 0 {
			continue
		}

		results = append(results, &Result{
			Symbol: threshold.symbol,
			Volume: int(volume),
		})
	}

	return results
}
