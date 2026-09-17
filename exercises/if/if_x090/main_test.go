// if_x090: Длительность и число
// Make the tests pass!
// I AM NOT DONE
//
// tooSlow должна вернуть true, если запрос длился дольше 5 секунд.
// Сейчас «медленным» считается почти любой запрос.
// Тренирует: нетипизированная константа 5 в сравнении с Duration — это 5 наносекунд.
// Сложность: hard
package main_test

import (
	"testing"
	"time"
)

func tooSlow(d time.Duration) bool {
	if d > 5 {
		return true
	}
	return false
}

func TestTooSlow(t *testing.T) {
	if tooSlow(300 * time.Millisecond) {
		t.Errorf("300ms is not slow")
	}
	if tooSlow(5 * time.Second) {
		t.Errorf("exactly 5s is not slow")
	}
	if !tooSlow(6 * time.Second) {
		t.Errorf("6s is slow")
	}
}
