package era

import (
	"time"
)

var (
	differenceUnits   = []int{60, 60, 24, 32, 12, 0}
	differenceSymbols = []Symbol{SymbolYear, SymbolMonth, SymbolDay, SymbolHour, SymbolMinute, SymbolSecond}
)

// Difference ...
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
			differences[i+1] -= 1

			// if it's days, we need to
			// calculate how many days are in the month
			if i == 3 {
				date := time.Date(from.Year(), from.Month(), differenceUnits[i], 0, 0, 0, 0, from.Location())
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

func deriveTimeMetrics(t time.Time) []int {
	y, m, d := t.Date()
	h, mm, s := t.Clock()

	return []int{s, mm, h, d, int(m), y}
}
