package era

// Symbol ...
type Symbol string

// Second ...
const (
	SymbolSecond Symbol = "s"
	SymbolMinute        = "m"
	SymbolHour          = "h"
	SymbolDay           = "d"
	SymbolWeek          = "w"
	SymbolMonth         = "M"
	SymbolYear          = "y"
	SymbolDecade        = "D"
	SymbolLong          = "l"
)

// String ...
func (s Symbol) String() string {
	return string(s)
}
