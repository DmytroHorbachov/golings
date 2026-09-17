// switch68
// Make the tests pass!

// I AM NOT DONE
//
// shouldLog возвращает true, если уровень сообщения не ниже минимального.
// Уровни: debug < info < warn < error; неизвестный уровень считается error.
// Тренирует: switch для преобразования строки в число.
// Сложность: medium
package main_test

import "testing"

func rank(level string) int {
	switch level {
	case "debug":
		return 0
	case "info":
		return 1
	case "warn":
		return 3
	}
	return 0
}

func shouldLog(level, min string) bool {
	return rank(level) >= rank(min)
}

func TestShouldLog(t *testing.T) {
	cases := []struct {
		level, min string
		want       bool
	}{{"debug", "info", false}, {"warn", "info", true}, {"info", "warn", false}, {"error", "warn", true}, {"fatal", "error", true}}
	for _, c := range cases {
		if got := shouldLog(c.level, c.min); got != c.want {
			t.Errorf("shouldLog(%s, %s) = %v, want %v", c.level, c.min, got, c.want)
		}
	}
}
