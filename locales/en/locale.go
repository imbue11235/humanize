package en

import (
	"github.com/imbue11235/humanize/language"
)

// Language ...
var Language = language.Map{
	"slice": language.Map{
		"connector": "and",
		"rest":      "one other|%d others",
	},
	"time": language.Map{
		"now":    "just now",
		"future": "in %s",
		"past":   "%s ago",
		"approximation": language.Map{
			"second": "a second|%d seconds",
			"minute": "a minute|%d minutes",
			"hour":   "an hour|%d hours",
			"day":    "a day|%d days",
			"week":   "a week|%d weeks",
			"month":  "a month|%d months",
			"year":   "a year|%d years",
			"decade": "a decade|%d decades",
			"long":   "a long time",
		},
		"precision": language.Map{
			"second": "1 second|%d seconds",
			"minute": "1 minute|%d minutes",
			"hour":   "1 hour|%d hours",
			"day":    "1 day|%d days",
			"month":  "1 month|%d months",
			"year":   "1 year|%d year",
		},
	},
	"int": language.Map{
		"K": "thousand",
		"M": "million",
		"B": "billion",
		"t": "trillion",
		"q": "quadrillion",
		"Q": "quintillion",
		"s": "sextillion",
		"S": "septillion",
	},
}
