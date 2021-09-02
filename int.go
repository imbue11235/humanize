package humanize

import (
	"fmt"
	"math"
	"math/big"
	"strconv"
)

type intBreakpoint struct {
	length float64
	key    string
}

var breakpoints = []intBreakpoint{
	{3, "thousand"},
	{6, "million"},
	{9, "billion"},
	{12, "trillion"},
	{15, "quadrillion"},
	{18, "quintillion"},
	{21, "sextillion"},
	{24, "septillion"},
	{27, "octillion"},
	{30, "nonillion"},
	{33, "decillion"},
	{36, "undecillion"},
	{39, "duodecillion"},
	{42, "tredecillion"},
	{45, "quattuordecillion"},
	{48, "quindecillion"},
	{51, "sexdecillion"},
	{54, "septendecillion"},
	{57, "octodecillion"},
	{60, "novendecillion"},
	{100, "googol"},
}

// Int ...
func Int(value int) string {
	length := math.Log10(float64(value))
	if length < breakpoints[0].length {
		return strconv.Itoa(value)
	}

	for i, breakpoint := range breakpoints {
		amount := length / breakpoint.length

		if amount < 1 {
			max := float64(value) / math.Pow(10, breakpoints[i-1].length)
			return fmt.Sprintf("%s %s", strconv.FormatFloat(max, 'f', -1, 64), breakpoints[i-1].key)
		}
	}

	return ""
}

// BigInt ...
func BigInt(value *big.Int) string {
	length := float64(len(value.String()) - 1)

	f := new(big.Float).SetInt(value)
	for i, breakpoint := range breakpoints {
		amount := length / breakpoint.length

		if amount < 1 {
			max := new(big.Float).Quo(f, big.NewFloat(math.Pow(10, breakpoints[i-1].length)))
			return fmt.Sprintf("%.2f %s", max, breakpoints[i-1].key)
		}
	}

	return ""
}
