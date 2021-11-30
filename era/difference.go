package era

import (
	"time"
)

var (
	differenceUnits   = []int{60, 60, 24, 32, 12, 0}
	differenceSymbols = []Symbol{SymbolYear, SymbolMonth, SymbolDay, SymbolHour, SymbolMinute, SymbolSecond}
)

// Difference finds the difference
// between two given times
func Difference(from, to time.Time) []*Result {
	return convertToResults(difference(from, to))
}

func difference(a, b time.Time) []int {
	// reverse times if necessary
	if a.After(b) {
		a, b = b, a
	}

	aMetrics := deriveTimeMetrics(a)
	bMetrics := deriveTimeMetrics(b)

	differences := make([]int, len(aMetrics))
	for i := range differences {
		differences[i] += bMetrics[i] - aMetrics[i]

		if differences[i] < 0 {
			differences[i] += differenceUnits[i]
			differences[i+1]--

			// if it's days, we need to
			// calculate how many days are in the month
			if i == 3 {
				date := time.Date(a.Year(), a.Month(), differenceUnits[i], 0, 0, 0, 0, a.Location())
				differences[i] -= date.Day()
			}
		}
	}

	// reversing slice, as we are calculating them
	// in order of 'secs, mins, hours etc.', but
	// want them returned as 'years, months, days etc.'
	return reverse(differences)
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

func reverse(a []int) []int {
	for i := len(a)/2 - 1; i >= 0; i-- {
		o := len(a) - 1 - i
		a[i], a[o] = a[o], a[i]
	}

	return a
}

func deriveTimeMetrics(from time.Time) []int {
	y, m, d := from.Date()
	h, mm, s := from.Clock()

	return []int{s, mm, h, d, int(m), y}
}
