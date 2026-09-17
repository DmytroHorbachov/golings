// variables29
// Make the tests pass!

// I AM NOT DONE
//
// letter(i) должна вернуть i-ю строчную латинскую букву ("a" для 0),
// а index(c) — номер буквы c.
// Тренирует: преобразования между rune, byte, int и string.
// Сложность: medium
package main_test

import (
	"strconv"
	"testing"
)

func letter(i int) string {
	return strconv.Itoa('a' + i)
}

func index(c byte) int {
	return int(c)
}

func TestLetters(t *testing.T) {
	_ = strconv.Itoa
	if got := letter(2); got != "c" {
		t.Errorf("letter(2) = %q, want c", got)
	}
	if got := index('z'); got != 25 {
		t.Errorf("index('z') = %d, want 25", got)
	}
}
