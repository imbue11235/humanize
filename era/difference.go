package era

import (
	"time"
)

var (
	differenceUnits   = []int{0, 12, 32, 24, 60, 60}
	differenceSymbols = []Symbol{SymbolYear, SymbolMonth, SymbolDay, SymbolHour, SymbolMinute, SymbolSecond}
)

// Difference
func Difference(from, to time.Time) []*Result {
	return convertToResults(difference(from, to))
}

func difference(from, to time.Time) []int {
	// reverse times if necessary
	if from.After(to) {
		from, to = to, from
	}

	fromMetrics := deriveTimeMetrics(from)
	toMetrics := deriveTimeMetrics(to)

	differences := make([]int, len(fromMetrics))
	for i := range differences {
		differences[i] = toMetrics[i] - fromMetrics[i]

		if differences[i] < 0 {
			differences[i] += differenceUnits[i]
			differences[i-1] -= 1

			// if it's days, we need to
			// calculate how many days are in the month
			if i == 2 {
				date := time.Date(from.Year(), from.Month(), differenceUnits[i], 0, 0, 0, 0, from.Location())
				differences[i] -= date.Day()
			}
		}
	}

	return differences
}

func convertToResults(differences []int) []*Result {
	var results []*Result
	for i, difference := range differences {
		if difference == 0 {
			continue
		}

		results = append(results, &Result{
			Symbol: differenceSymbols[i],
			Volume: difference,
		})
	}

	return results
}

func deriveTimeMetrics(t time.Time) []int {
	y, m, d := t.Date()
	h, mm, s := t.Clock()

	return []int{y, int(m), d, h, mm, s}
}
