package da

import (
	"github.com/imbue11235/humanize/language"
)

// Language ...
var Language = language.Map{
	"slice": language.Map{
		"connector": "og",
		"rest":      "èn anden|%d andre",
	},
	"time": language.Map{
		"now":    "lige nu",
		"future": "om %s",
		"past":   "%s siden",
		"estimation": language.Map{
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
		"precision": language.Map{
			"s": "1 sekund|%d sekunder",
			"m": "1 minut|%d minutter",
			"h": "1 time|%d timer",
			"d": "1 dag|%d dage",
			"M": "1 måned|%d måneder",
			"y": "1 år|%d år",
		},
	},
	"int": language.Map{
		"K": "tusind|tusinde",
		"M": "million|millioner",
		"B": "milliard|milliarder",
		"t": "billion|billioner",
		"q": "billiard|billiarder",
		"Q": "trillion|trillioner",
		"s": "trilliard|trilliarder",
		"S": "kvadrillion|kvadrillioner",
	},
}
