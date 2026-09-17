// maps_x009: Множество
// Make the tests pass!
// I AM NOT DONE
//
// isBanned проверяет, входит ли слово в множество запрещённых.
// Тренирует: map[string]bool как множество.
// Сложность: easy
package main_test

import "testing"

var banned = map[string]bool{"spam": true, "scam": true}

func isBanned(w string) bool {
	return !banned[w]
}

func TestIsBanned(t *testing.T) {
	if !isBanned("spam") || isBanned("hello") {
		t.Errorf("isBanned works incorrectly")
	}
}
