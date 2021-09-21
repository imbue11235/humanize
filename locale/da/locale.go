package da

import "github.com/imbue11235/humanize/locale"

// Code contains the locale code for danish locale
var Code = "da"

// Locale contains the danish locale data
var Locale = locale.Map{
	"slice": locale.Map{
		"connector": "og",
		"rest":      "èn anden|%d andre",
	},
	"time": locale.Map{
		"now":    "lige nu",
		"future": "om %s",
		"past":   "%s siden",
		"estimation": locale.Map{
			"s": "et sekund|%d sekunder",
			"m": "et minut|%d minutter",
			"h": "en time|%d timer",
			"d": "en dag|%d dage",
			"w": "en uge|%d uger",
			"M": "en måned|%d måneder",
			"y": "et år|%d år",
			"D": "et årti|%d årtier",
			"l": "lang tid",
		},
		"precision": locale.Map{
			"s": "1 sekund|%d sekunder",
			"m": "1 minut|%d minutter",
			"h": "1 time|%d timer",
			"d": "1 dag|%d dage",
			"M": "1 måned|%d måneder",
			"y": "1 år|%d år",
		},
	},
	"int": locale.Map{
		"K":  "tusind|tusinde",
		"M":  "million|millioner",
		"B":  "milliard|milliarder",
		"T":  "billion|billioner",
		"Q":  "billiard|billiarder",
		"Qi": "trillion|trillioner",
		"Sx": "trilliard|trilliarder",
		"Sp": "kvadrillion|kvadrillioner",
	},
}
