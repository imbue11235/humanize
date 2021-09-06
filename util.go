package humanize

import "math"

func truncFloat(value float64, precision int) float64 {
	power := math.Pow(10, float64(precision))
	return math.Round(value*power) / power
}
