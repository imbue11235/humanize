package humanize

import (
	"math"
)

// truncFloat truncates the given float value to the given precision
// e.g. 1.23456 with precision 2 => 1.23
func truncFloat(value float64, precision int) float64 {
	power := math.Pow(10, float64(precision))
	return math.Round(value*power) / power
}
