package de

import "github.com/imbue11235/humanize/locale"

// Code contains the locale code for german locale
var Code = "de"

// Locale contains the german locale data
var Locale = locale.Map{
	"slice": locale.Map{
		"connector": "und",
		"rest":      "[1] ein anderer|[2-*] %d andere",
	},
	"time": locale.Map{
		"now":    "gerade jetzt",
		"future": "in %s",
		"past":   "vor %s",
		"estimation": locale.Map{
			"s": "[1] eine Sekunde|[2-*] %d Sekunden",
			"m": "[1] eine Minute|[2-*] %d Minuten",
			"h": "[1] eine Stunde|[2-*] %d Stunden",
			"d": "[1] ein Tag|[2-*] %d Tage",
			"w": "[1] eine Woche|[2-*] %d Wochen",
			"M": "[1] ein Monat|[2-*] %d Monate",
			"y": "[1] ein Jahr|[2-*] %d Jahre",
			"D": "[1] ein Jahrzehnt|[2-*] %d Jahrzehnte",
			"l": "lange Zeit",
		},
		"precision": locale.Map{
			"s": "[1] 1 Sekunde|[2-*] %d Sekunden",
			"m": "[1] 1 Minute|[2-*] %d Minuten",
			"h": "[1] 1 Stunde|[2-*] %d Stunden",
			"d": "[1] 1 Tag|[2-*] %d Tage",
			"M": "[1] 1 Monat|[2-*] %d Monate",
			"y": "[1] 1 Jahr|[2-*] %d Jahre",
		},
	},
	"int": locale.Map{
		"K":  "tausend",
		"M":  "[1] Million|[2-*] Millionen",
		"B":  "[1] Milliarde|[2-*] Milliarden",
		"T":  "[1] Billion|[2-*] Billionen",
		"Q":  "[1] Billiarde|[2-*] Billiarden",
		"Qi": "[1] Trillion|[2-*] Trillionen",
		"Sx": "[1] Sextillion|[2-*] Sextillionen",
		"Sp": "[1] Septillion|[2-*] Septillionen",
	},
}
