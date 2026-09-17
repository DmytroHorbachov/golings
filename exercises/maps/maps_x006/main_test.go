// maps_x006: Счётчик
// Make the tests pass!
// I AM NOT DONE
//
// hit увеличивает счётчик посещений страницы.
// Тренирует: m[k]++ для отсутствующего ключа начинает с нуля.
// Сложность: easy
package main_test

import "testing"

func hit(stats map[string]int, page string) {
	stats[page] = 1
}

func TestHit(t *testing.T) {
	s := map[string]int{}
	hit(s, "/")
	hit(s, "/")
	hit(s, "/about")
	if s["/"] != 2 || s["/about"] != 1 {
		t.Errorf("stats = %v", s)
	}
}
