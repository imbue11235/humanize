package era

import (
	"testing"
	"time"
)

func TestTimeInterval(t *testing.T) {
	now := time.Now()
	tests := []struct {
		from, to time.Time
		expected []int
	}{
		{now, now.Add(24 * time.Hour), []int{0, 0, 1, 0, 0, 0}},
		{now, now.Add(35 * day), []int{0, 1, 5, 0, 0, 0}},
		{now, now.Add(28 * day), []int{0, 0, 28, 0, 0, 0}},
		{now, now.Add(-28 * day), []int{0, 0, 28, 0, 0, 0}},
		{now, now.Add(-32 * day), []int{0, 1, 1, 0, 0, 0}},
	}

	for _, test := range tests {
		differences := difference(test.from, test.to)

		if len(test.expected) != len(differences) {
			t.Errorf("expected and actual does not have same length (%d != %d)", len(test.expected), len(differences))
			continue
		}

		for i, v := range test.expected {
			if v != differences[i] {
				t.Errorf("expected %+v, but got %+v", test.expected, differences)
				return
			}
		}
	}
}
