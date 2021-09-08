package era

import (
	"testing"
	"time"
)

func TestPositiveAbsDuration(t *testing.T) {
	positiveDuration := time.Duration(3500)
	absPositiveDuration := absDuration(positiveDuration)

	if positiveDuration != absPositiveDuration {
		t.Errorf("expected %d to be %d", absPositiveDuration, positiveDuration)
	}
}

func TestNegativeAbsDuration(t *testing.T) {
	expectedDuration := time.Duration(10000)
	negativeDuration := time.Duration(-10000)
	absNegativeDuration := absDuration(negativeDuration)

	if absNegativeDuration != expectedDuration {
		t.Errorf("expected %d to be %d", absNegativeDuration, negativeDuration)
	}
}
