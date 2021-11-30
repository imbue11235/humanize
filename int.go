package humanize

import (
	"fmt"
	"math"
	"strconv"
)

var suffixes = []string{"K", "M", "B", "T", "Q", "Qi", "Sx", "Sp", "O"}

// Int converts an integer into a readable string representation,
// rounding the volume and adding the "name" of the number as a suffix
// e.g. 1000000 = 1 million
func Int(value uint64) string {
	return formatInt(value, true)
}

// IntWithSymbol converts an integer into a readable string representation,
// rounding the volume and adding the "symbol" of the number as a suffix
// e.g. 1000000 = 1M
func IntWithSymbol(value uint64) string {
	return formatInt(value, false)
}

func formatInt(value uint64, translateSymbol bool) string {
	if value < 1000 {
		return strconv.FormatUint(value, 10)
	}

	volume, suffix := deriveVolumeAndSuffix(value)

	formatPrecision := 0
	if volume != math.Floor(volume) {
		formatPrecision++
	}

	if translateSymbol {
		return fmt.Sprintf(
			"%s %s",
			strconv.FormatFloat(volume, 'f', formatPrecision, 64),
			pluralize(fmt.Sprintf("int.%s", suffix), int(volume)),
		)
	}

	return fmt.Sprintf(
		"%s%s",
		strconv.FormatFloat(volume, 'f', formatPrecision, 64),
		suffix,
	)
}

func deriveVolumeAndSuffix(value uint64) (float64, string) {
	floatValue := float64(value)

	// we have to truncate the result of log10 to a smaller precision
	// because when numbers get really large, like 1 quadrillion,
	// the math.Log10 function becomes imprecise.
	// e.g. math.Log10(1000000000000000) => 14.999999999999998,
	// where it should be 15
	decimal := truncFloat(math.Log10(floatValue), 5)

	// dividing with 3, because every step between definitions of
	// large numbers, is 3 extra zeroes.
	// e.g. million = 10^6, billion = 10^9
	power := math.Floor(decimal / 3)
	volume := floatValue / math.Pow(10, 3*power)

	return truncFloat(volume, 1), suffixes[int(power)-1]
}
