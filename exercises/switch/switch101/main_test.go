// switch101
// Make the tests pass!

// I AM NOT DONE
//
// Status.String должна возвращать "active", "blocked" и "unknown".
// Тренирует: switch в методе именованного типа.
// Сложность: easy
package main_test

import "testing"

type Status int

const (
	Active Status = iota
	Blocked
)

func (s Status) String() string {
	switch s {
	case Active:
		return "active"
	case Active + 2:
		return "blocked"
	}
	return "unknown"
}

func TestStatusString(t *testing.T) {
	cases := map[Status]string{Active: "active", Blocked: "blocked", 7: "unknown"}
	for in, want := range cases {
		if got := in.String(); got != want {
			t.Errorf("Status(%d).String() = %s, want %s", int(in), got, want)
		}
	}
}
