package en

import (
	"github.com/imbue11235/humanize/locale"
)

// Code contains the locale code for english locale
var Code = "en"

// Locale contains the english locale data
var Locale = locale.Map{
	"slice": locale.Map{
		"connector": "and",
		"rest":      "one other|%d others",
	},
	"time": locale.Map{
		"now":    "just now",
		"future": "in %s",
		"past":   "%s ago",
		"estimation": locale.Map{
			"s": "a second|%d seconds",
			"m": "a minute|%d minutes",
			"h": "an hour|%d hours",
			"d": "a day|%d days",
			"w": "a week|%d weeks",
			"M": "a month|%d months",
			"y": "a year|%d years",
			"D": "a decade|%d decades",
			"l": "a long time",
		},
		"precision": locale.Map{
			"s": "1 second|%d seconds",
			"m": "1 minute|%d minutes",
			"h": "1 hour|%d hours",
			"d": "1 day|%d days",
			"M": "1 month|%d months",
			"y": "1 year|%d years",
		},
	},
	"int": locale.Map{
		"K":  "thousand",
		"M":  "million",
		"B":  "billion",
		"T":  "trillion",
		"Q":  "quadrillion",
		"Qi": "quintillion",
		"Sx": "sextillion",
		"Sp": "septillion",
	},
}
