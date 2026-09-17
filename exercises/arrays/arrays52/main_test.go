// arrays52
// Make the tests pass!

// I AM NOT DONE
//
// buildTable строит таблицу шифра Цезаря [26]byte для сдвига k,
// а encrypt кодирует ей строчные буквы.
// Тренирует: предвычисленный массив подстановки.
// Сложность: medium
package main_test

import "testing"

func buildTable(k int) [26]byte {
	var t [26]byte
	for i := range t {
		t[i] = byte(i + k)
	}
	return t
}

func encrypt(s string, t [26]byte) string {
	b := []byte(s)
	for i, c := range b {
		if c >= 'a' && c <= 'z' {
			b[i] = t[c]
		}
	}
	return string(b)
}

func TestEncrypt(t *testing.T) {
	table := buildTable(3)
	if got := encrypt("xyz abc", table); got != "abc def" {
		t.Errorf("encrypt = %q, want %q", got, "abc def")
	}
}
