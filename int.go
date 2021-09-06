package humanize

import (
	"fmt"
	"math"
	"strconv"
)

var suffixes = []string{"K", "M", "B", "t", "q", "Q", "s", "S", "o"}

// Int ...
func Int(value int64) string {
	volume, suffix := deriveVolumeAndSuffix(value)

	formatPrecision := 0
	if volume != math.Floor(volume) {
		formatPrecision++
	}

	return fmt.Sprintf(
		"%s %s",
		strconv.FormatFloat(volume, 'f', formatPrecision, 64),
		pluralize(fmt.Sprintf("int.%s", suffix), int(volume)),
	)
}

func deriveVolumeAndSuffix(value int64) (float64, string) {
	fValue := float64(value)
	decimal := math.RoundToEven(math.Log10(fValue))
	power := math.Floor(decimal / 3)
	volume := fValue / math.Pow(10, 3*power)

	return truncFloat(volume, 1), suffixes[int(power)-1]
}
