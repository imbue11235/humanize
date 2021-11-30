package era

import (
	"math"
	"time"
)

const proximityUpperBound = 0.2

var estimationThresholds = []threshold{
	{SymbolLong, long},
	{SymbolDecade, decade},
	{SymbolYear, year},
	{SymbolMonth, month},
	{SymbolWeek, week},
	{SymbolDay, day},
	{SymbolHour, hour},
	{SymbolMinute, minute},
	{SymbolSecond, second},
}

// DurationToEstimation calculates a loose estimate
// of given duration
func DurationToEstimation(duration time.Duration) *Result {
	delta := absDuration(duration).Seconds()

	for _, threshold := range estimationThresholds {
		// calculate the proximity time distance
		proximity := delta / threshold.duration.Seconds()
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
				Volume: int(next),
				Symbol: threshold.symbol,
			}
		}

		if proximity >= 1 {
			return &Result{
				Volume: int(proximity),
				Symbol: threshold.symbol,
			}
		}
	}

	// if no result, it's assumed that it's less
	// than a second
	return &Result{
		Volume: 0,
		Symbol: SymbolSecond,
	}
}
