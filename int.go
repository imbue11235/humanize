package humanize

import (
	"fmt"
	"math"
	"strconv"
)

type intRepresentation struct {
	zeroes float64
	text   string
}

var ints = []intRepresentation{
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

func Int(value int) string {
	length := math.Log10(float64(value))
	if length < ints[0].zeroes {
		return strconv.Itoa(value)
	}

	for i, in := range ints {
		amount := length / in.zeroes

		if amount < 1 {
			best := i - 1
			max := float64(value) / math.Pow(10, ints[best].zeroes)
			return fmt.Sprintf("%s %s", strconv.FormatFloat(max, 'f', -1, 64), ints[best].text)
			break
		}
	}

	return ""
}
