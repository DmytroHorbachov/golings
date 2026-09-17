// functions27
// Make the tests pass!

// I AM NOT DONE
//
// processItems должна в конце сообщить, сколько элементов обработано.
// В отчёте всегда 0.
// Тренирует: при defer s.Method() получатель-значение копируется сразу.
// Сложность: hard
package main_test

import (
	"fmt"
	"testing"
)

type Stats struct{ Count int }

func (s Stats) Report(out *string) {
	*out = fmt.Sprintf("processed %d", s.Count)
}

func processItems(items []string, out *string) {
	var s Stats
	defer s.Report(out)
	for range items {
		s.Count++
	}
}

func TestProcessItems(t *testing.T) {
	var report string
	processItems([]string{"a", "b", "c"}, &report)
	if report != "processed 3" {
		t.Errorf("report = %q, want %q", report, "processed 3")
	}
}
