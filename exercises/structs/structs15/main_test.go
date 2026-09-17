// structs15
// Make the tests pass!

// I AM NOT DONE
//
// У структуры Timer есть поле Elapsed и метод Elapsed. Код не компилируется.
// Тренирует: поле и метод типа не могут иметь одно имя.
// Сложность: hard
package main_test

import (
	"testing"
	"time"
)

type Timer struct {
	Elapsed time.Duration
}

func (t *Timer) Add(d time.Duration) { t.Elapsed += d }

func (t Timer) Elapsed() time.Duration { return t.Elapsed }

func TestTimer(t *testing.T) {
	var tm Timer
	tm.Add(time.Second)
	tm.Add(time.Second)
	if tm.Elapsed() != 2*time.Second {
		t.Errorf("Elapsed = %v", tm.Elapsed())
	}
}
