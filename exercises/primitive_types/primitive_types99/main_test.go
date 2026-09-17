// primitive_types99
// Make the tests pass!

// I AM NOT DONE
//
// rawString должна собрать строку из «сырых» байтов (не текста).
// Для байтов больше 127 длина результата не совпадает с числом байтов.
// Тренирует: string(byte(200)) кодирует руну U+00C8 в UTF-8 (два байта).
// Сложность: hard
package main_test

import "testing"

func rawString(data []byte) string {
	s := ""
	for _, b := range data {
		s += string(rune(b))
	}
	return s
}

func TestRawString(t *testing.T) {
	data := []byte{0x41, 0xC8, 0xFF}
	got := rawString(data)
	if len(got) != 3 || got[1] != 0xC8 || got[2] != 0xFF {
		t.Errorf("rawString = % x, want 41 c8 ff", got)
	}
}
