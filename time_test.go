package humanize

import (
	"testing"
	"time"
)

func TestTimeFrom(t *testing.T) {
	/*
		tests := []struct {
			from     string
			to       string
			expected string
		}{
			{"2021-01-01 22:00:00", "2021-01-02 22:00:00", ""},
		}


		for _, test := range tests {
			//to := parseTime(t, test.to)
			//from := parseTime(t, test.from)

			//humanize.Time(to).From(from)
		}*/

	/*
		fmt.Println(humanize.SetLanguage("en"))
		fmt.Println(humanize.Time(from).From(to))

		fmt.Println(humanize.SetLanguage("da"))
		fmt.Println(humanize.Time(from).From(to))*/
}

func parseTime(t *testing.T, value string) time.Time {
	parsed, err := time.Parse("yyyy-mm-dd hh:mm:ss", value)
	if err != nil {
		t.Errorf("could not parse time '%s'", value)
	}

	return parsed
}
