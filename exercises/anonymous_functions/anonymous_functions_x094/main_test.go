// anonymous_functions_x094: Время создания и вызова
// Make the tests pass!
// I AM NOT DONE
//
// stamper возвращает литерал, который должен ставить метку текущего времени
// в момент вызова. Сейчас время берётся один раз при создании.
// Тренирует: вычисление внутри литерала происходит при каждом вызове.
// Сложность: hard
package main_test

import (
	"testing"
	"time"
)

func stamper(now func() time.Time) func(msg string) string {
	ts := now().Format("15:04:05")
	return func(msg string) string {
		return ts + " " + msg
	}
}

func TestStamper(t *testing.T) {
	base := time.Date(2024, 1, 1, 10, 0, 0, 0, time.UTC)
	tick := 0
	clock := func() time.Time { tick++; return base.Add(time.Duration(tick) * time.Second) }
	stamp := stamper(clock)
	a := stamp("a")
	b := stamp("b")
	if a != "10:00:01 a" || b != "10:00:02 b" {
		t.Errorf("a = %q, b = %q", a, b)
	}
}
