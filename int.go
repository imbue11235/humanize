package humanize

import (
	"fmt"
	"math"
	"strconv"
)

var suffixes = []string{"K", "M", "B", "T", "Q", "Qi", "Sx", "Sp", "O"}

// Int ...
func Int(value uint64) string {
	return formatInt(value, true)
}

// IntWithSymbol ...
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
	fValue := float64(value)
	whole, frac := math.Modf(math.Log10(fValue))

	// workaround: when numbers get really large, the precision gets worse
	// we end up with edge cases where 1 quadrillion resolves as 1000 trillion
	// instead, because of the fraction being 0.9999...
	if 1-frac < 0.0001 {
		whole++
	}

	power := math.Floor(whole / 3)
	volume := fValue / math.Pow(10, 3*power)

	return truncFloat(volume, 1), suffixes[int(power)-1]
}
