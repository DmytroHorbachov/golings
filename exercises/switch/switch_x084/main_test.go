// switch_x084: switch по срезу
// Make the tests pass!
// I AM NOT DONE
//
// chord называет аккорд по списку нот. Код не компилируется: срезы
// нельзя использовать в switch.
// Тренирует: тег switch должен быть сравнимым типом.
// Сложность: hard
package main_test

import (
	"strings"
	"testing"
)

func chord(notes []string) string {
	switch notes {
	case []string{"C", "E", "G"}:
		return "C major"
	case []string{"A", "C", "E"}:
		return "A minor"
	}
	return "unknown"
}

func TestChord(t *testing.T) {
	_ = strings.Join
	if got := chord([]string{"C", "E", "G"}); got != "C major" {
		t.Errorf("chord(CEG) = %s", got)
	}
	if got := chord([]string{"A", "C", "E"}); got != "A minor" {
		t.Errorf("chord(ACE) = %s", got)
	}
	if got := chord([]string{"C", "G"}); got != "unknown" {
		t.Errorf("chord(CG) = %s", got)
	}
}
