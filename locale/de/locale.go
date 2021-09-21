package de

import "github.com/imbue11235/humanize/locale"

// Code contains the locale code for german locale
var Code = "de"

// Locale contains the german locale data
var Locale = locale.Map{
	"slice": locale.Map{
		"connector": "und",
		"rest":      "ein anderer|%d andere",
	},
	"time": locale.Map{
		"now":    "gerade jetzt",
		"future": "in %s",
		"past":   "vor %s",
		"estimation": locale.Map{
			"s": "eine Sekunde|%d Sekunden",
			"m": "eine Minute|%d Minuten",
			"h": "eine Stunde|%d Stunden",
			"d": "ein Tag|%d Tage",
			"w": "eine Woche|%d Wochen",
			"M": "ein Monat|%d Monate",
			"y": "ein Jahr|%d Jahre",
			"D": "ein Jahrzehnt|%d Jahrzehnte",
			"l": "lange Zeit",
		},
		"precision": locale.Map{
			"s": "1 Sekunde|%d Sekunden",
			"m": "1 Minute|%d Minuten",
			"h": "1 Stunde|%d Stunden",
			"d": "1 Tag|%d Tage",
			"M": "1 Monat|%d Monate",
			"y": "1 Jahr|%d Jahre",
		},
	},
	"int": locale.Map{
		"K":  "tausend",
		"M":  "Million|Millionen",
		"B":  "Milliarde|Milliarden",
		"T":  "Billion|Billionen",
		"Q":  "Billiarde|Billiarden",
		"Qi": "Trillion|Trillionen",
		"Sx": "Sextillion|Sextillionen",
		"Sp": "Septillion|Septillionen",
	},
}
