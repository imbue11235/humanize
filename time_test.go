package humanize

import (
	"fmt"
	"testing"
	"time"
)

func TestFromNow(t *testing.T) {
	from := time.Date(2011, 11, 12, 14, 0, 0, 0, time.UTC)
	to := time.Date(2011, 11, 12, 14, 30, 22, 0, time.UTC)

	fmt.Println(SetLocale("da"))
	fmt.Println(ExactTime(from).From(to))

	fmt.Println(SetLocale("en"))
	fmt.Println(ExactTime(from).From(to))
}