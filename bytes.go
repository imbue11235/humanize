package humanize

import (
	"fmt"
	"math"
	"strconv"
)

// defines the starting precision for formatted bytes
const defaultBytesFormatPrecision = 0

// defines the lower bound that the volume of bytes would
// have to be lower than, for a more precise definition.
// e.g. "9.5MB" => "9.5MB", "10.5MB" => "10MB"
const bytesFormatPrecisionMinVolume = 10

var (
	// multi-byte units
	// for reference: https://en.wikipedia.org/wiki/Byte#Multiple-byte_units
	decimal         = []string{"B", "kB", "MB", "GB", "TB", "PB", "EB", "ZB", "YB"}
	binary          = []string{"B", "KiB", "MiB", "GiB", "TiB", "PiB", "EiB", "ZiB", "YiB"}
	shortFormBinary = []string{"B", "K", "M", "G", "T", "P", "E", "Z", "Y"}
)

// Bytes uses the SI prefixes (powers of 10, e.g. 1000b = 1kB)
// for converting the bytes into their human-readable representation.
//
//		s := humanize.Bytes(2500000)
//		fmt.Println(s) => "2.5 MB"
//
func Bytes(value uint64) string {
	return formatBytes(value, 1000, decimal, "%s %s")
}

// BinaryBytes uses the binary system (powers of 2, e.g. 1024b = 1KiB)
// for converting the bytes into their human-readable representation.
//
//		s := humanize.BinaryBytes(2500000)
//		fmt.Println(s) => "2.4 MiB"
//
func BinaryBytes(value uint64) string {
	return formatBytes(value, 1024, binary, "%s %s")
}

// ShortFormBinaryBytes uses the binary system (powers of 2, e.g. 1024b = 1KiB)
// for converting the bytes into a GNU-like format.
//
//		s := humanize.ShortFormBinaryBytes(2500000)
//		fmt.Println(s) => "2.4M"
//
func ShortFormBinaryBytes(value uint64) string {
	return formatBytes(value, 1024, shortFormBinary, "%s%s")
}

func formatBytes(bytes uint64, base float64, suffixes []string, format string) string {
	fBytes := float64(bytes)
	if fBytes < base {
		return fmt.Sprintf(format, strconv.Itoa(int(bytes)), suffixes[0])
	}

	index := math.Floor(math.Log(fBytes) / math.Log(base))
	volume := base * fBytes / math.Pow(base, index+1)

	formatPrecision := defaultBytesFormatPrecision
	if volume < bytesFormatPrecisionMinVolume {
		formatPrecision++
	}

	return fmt.Sprintf(format, strconv.FormatFloat(volume, 'f', formatPrecision, 64), suffixes[int(index)])
}
