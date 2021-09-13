package era

// Symbol ...
type Symbol string

// Second ...
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

// String ...
func (s Symbol) String() string {
	return string(s)
}
