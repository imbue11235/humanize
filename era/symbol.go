package era

// Symbol type of time units
type Symbol string

// SymbolSecond ...
const (
	SymbolSecond Symbol = "s"
	SymbolMinute Symbol = "m"
	SymbolHour   Symbol = "h"
	SymbolDay    Symbol = "d"
	SymbolWeek   Symbol = "w"
	SymbolMonth  Symbol = "M"
	SymbolYear   Symbol = "y"
	SymbolDecade Symbol = "D"
	SymbolLong   Symbol = "l"
)

// String returns the string
// representation of the string
func (s Symbol) String() string {
	return string(s)
}
