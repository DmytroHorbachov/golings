// anonymous_functions44
// Make the tests pass!

// I AM NOT DONE
//
// timed выполняет функцию и записывает длительность в map по имени;
// часы передаются литералом.
// Тренирует: defer-литерал с доступом к захваченным значениям.
// Сложность: medium
package main_test

import (
	"testing"
	"time"
)

func timed(stats map[string]time.Duration, name string, now func() time.Time, f func()) {
	start := now()
	f()
	stats[name] = now().Sub(start)
}

func TestTimed(t *testing.T) {
	base := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	tick := 0
	clock := func() time.Time { tick++; return base.Add(time.Duration(tick) * time.Second) }
	stats := map[string]time.Duration{}
	timed(stats, "load", clock, func() {})
	func() {
		defer func() { _ = recover() }()
		timed(stats, "load", clock, func() { panic("x") })
	}()
	if stats["load"] != 2*time.Second {
		t.Errorf("load = %v, want 2s", stats["load"])
	}
}
