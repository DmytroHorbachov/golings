// if_x049: Класс HTTP-статуса
// Make the tests pass!
// I AM NOT DONE
//
// statusClass: 200–299 — "success", 300–399 — "redirect", 400–499 — "client error",
// 500–599 — "server error", иначе "unknown".
// Тренирует: проверку диапазонов с верхней и нижней границей.
// Сложность: medium
package main_test

import "testing"

func statusClass(code int) string {
	if code >= 200 {
		return "success"
	} else if code >= 300 {
		return "redirect"
	} else if code >= 400 {
		return "client error"
	} else if code >= 500 && code < 600 {
		return "server error"
	}
	return "unknown"
}

func TestStatusClass(t *testing.T) {
	cases := map[int]string{200: "success", 204: "success", 301: "redirect", 404: "client error", 503: "server error", 700: "unknown", 99: "unknown"}
	for in, want := range cases {
		if got := statusClass(in); got != want {
			t.Errorf("statusClass(%d) = %s, want %s", in, got, want)
		}
	}
}
