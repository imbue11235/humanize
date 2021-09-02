package humanize

import (
	"fmt"
	"math"
	"strconv"
)

const defaultFormatPrecision = 0
const formatPrecisionLowerBound = 10

var (
	// multi-byte units
	// for reference: https://en.wikipedia.org/wiki/Byte#Multiple-byte_units
	decimal         = []string{"B", "kB", "MB", "GB", "TB", "PB", "EB", "ZB", "YB"}
	binary          = []string{"B", "KiB", "MiB", "GiB", "TiB", "PiB", "EiB", "ZiB", "YiB"}
	shortFormBinary = []string{"B", "K", "M", "G", "T", "P", "E", "Z", "Y"}
)

// Bytes ...
func Bytes(value uint64) string {
	return formatBytes(value, 1000, decimal, "%s %s")
}

// BinaryBytes ...
func BinaryBytes(value uint64) string {
	return formatBytes(value, 1024, binary, "%s %s")
}

// ShortFormBinaryBytes ...
func ShortFormBinaryBytes(value uint64) string {
	return formatBytes(value, 1024, shortFormBinary, "%s%s")
}

func formatBytes(bytes uint64, base float64, suffixes []string, format string) string {
	fBytes := float64(bytes)
	if fBytes < base {
		return fmt.Sprintf(format, strconv.Itoa(int(bytes)), suffixes[0])
	}

	formatPrecision := defaultFormatPrecision

	index := math.Floor(math.Log(fBytes) / math.Log(base))
	amount := base * fBytes / math.Pow(base, index+1)

	if amount < formatPrecisionLowerBound {
		formatPrecision++
	}

	return fmt.Sprintf(format, strconv.FormatFloat(amount, 'f', formatPrecision, 64), suffixes[int(index)])
}
