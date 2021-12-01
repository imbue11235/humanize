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
		"rest":      "[1] one other|[2-*] %d others",
	},
	"time": locale.Map{
		"now":    "just now",
		"future": "in %s",
		"past":   "%s ago",
		"estimation": locale.Map{
			"s": "[0-5] a few seconds|[6-*] %d seconds",
			"m": "[1] a minute|[2-*] %d minutes",
			"h": "[1] an hour|[2-*] %d hours",
			"d": "[1] a day|[2-*] %d days",
			"w": "[1] a week|[2-*] %d weeks",
			"M": "[1] a month|[2-*] %d months",
			"y": "[1] a year|[2-*] %d years",
			"D": "[1] a decade|[2-*] %d decades",
			"l": "a long time",
		},
		"precision": locale.Map{
			"s": "[1] 1 second|[2-*] %d seconds",
			"m": "[1] 1 minute|[2-*] %d minutes",
			"h": "[1] 1 hour|[2-*] %d hours",
			"d": "[1] 1 day|[2-*] %d days",
			"M": "[1] 1 month|[2-*] %d months",
			"y": "[1] 1 year|[2-*] %d years",
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
