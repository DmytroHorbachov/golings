// functions_x063: Замер времени через defer
// Make the tests pass!
// I AM NOT DONE
//
// timed должна выполнить f и записать в *elapsed разницу показаний часов now,
// даже если f паникует. Часы передаются функцией для детерминированности.
// Тренирует: defer с замыканием для измерения длительности.
// Сложность: medium
package main_test

import (
	"testing"
	"time"
)

func timed(now func() time.Time, elapsed *time.Duration, f func()) {
	start := now()
	f()
	*elapsed = now().Sub(start)
}

func TestTimed(t *testing.T) {
	base := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	tick := 0
	clock := func() time.Time {
		tick++
		return base.Add(time.Duration(tick) * time.Second)
	}
	var d time.Duration
	timed(clock, &d, func() { clock(); clock() })
	if d != 3*time.Second {
		t.Errorf("elapsed = %v, want 3s", d)
	}
	var p time.Duration
	func() {
		defer func() { _ = recover() }()
		timed(clock, &p, func() { clock(); panic("fail") })
	}()
	if p != 2*time.Second {
		t.Errorf("elapsed after panic = %v, want 2s", p)
	}
}
