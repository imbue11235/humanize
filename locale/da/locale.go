package da

import "github.com/imbue11235/humanize/locale"

// Code contains the locale code for danish locale
var Code = "da"

// Locale contains the danish locale data
var Locale = locale.Map{
	"slice": locale.Map{
		"connector": "og",
		"rest":      "[1] èn anden|[2-*] %d andre",
	},
	"time": locale.Map{
		"now":    "lige nu",
		"future": "om %s",
		"past":   "%s siden",
		"estimation": locale.Map{
			"s": "[0-5] et par sekunder|[6-*] %d sekunder",
			"m": "[1] et minut|[2-*] %d minutter",
			"h": "[1] en time|[2-*] %d timer",
			"d": "[1] en dag|[2-*] %d dage",
			"w": "[1] en uge|[2-*] %d uger",
			"M": "[1] en måned|[2-*] %d måneder",
			"y": "[1] et år|[2-*] %d år",
			"D": "[1] et årti|[2-*] %d årtier",
			"l": "lang tid",
		},
		"precision": locale.Map{
			"s": "[1] 1 sekund|[2-*] %d sekunder",
			"m": "[1] 1 minut|[2-*] %d minutter",
			"h": "[1] 1 time|[2-*] %d timer",
			"d": "[1] 1 dag|[2-*] %d dage",
			"M": "[1] 1 måned|[2-*] %d måneder",
			"y": "[1] 1 år|[2-*] %d år",
		},
	},
	"int": locale.Map{
		"K":  "[1] tusind|[2-*] tusinde",
		"M":  "[1] million|[2-*] millioner",
		"B":  "[1] milliard|[2-*] milliarder",
		"T":  "[1] billion|[2-*] billioner",
		"Q":  "[1] billiard|[2-*] billiarder",
		"Qi": "[1] trillion|[2-*] trillioner",
		"Sx": "[1] trilliard|[2-*] trilliarder",
		"Sp": "[1] kvadrillion|[2-*] kvadrillioner",
	},
}
