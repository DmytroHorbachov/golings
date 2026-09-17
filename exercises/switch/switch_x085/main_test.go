// switch_x085: Weekday и int
// Make the tests pass!
// I AM NOT DONE
//
// dayKind принимает номер дня как int и сравнивает его с константами time.Weekday.
// Код не компилируется из-за несовпадения типов.
// Тренирует: case-значения должны быть совместимы с типом тега.
// Сложность: hard
package main_test

import (
	"testing"
	"time"
)

func dayKind(n int) string {
	switch n {
	case time.Saturday, time.Sunday:
		return "weekend"
	}
	return "workday"
}

func TestDayKind(t *testing.T) {
	cases := map[int]string{0: "weekend", 6: "weekend", 1: "workday", 5: "workday"}
	for in, want := range cases {
		if got := dayKind(in); got != want {
			t.Errorf("dayKind(%d) = %s, want %s", in, got, want)
		}
	}
}
