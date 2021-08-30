package humanize_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/imbue11235/humanize"
	_ "github.com/imbue11235/humanize/locales/da"
)

func TestFromNow(t *testing.T) {
	from := time.Date(2011, 11, 12, 14, 0, 0, 0, time.UTC)
	to := time.Date(2011, 11, 12, 14, 30, 22, 0, time.UTC)

	fmt.Println(humanize.SetLanguage("en"))
	fmt.Println(humanize.Time(from).From(to))

	fmt.Println(humanize.SetLanguage("da"))
	fmt.Println(humanize.Time(from).From(to))
}
