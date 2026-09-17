// maps51
// Make the tests pass!

// I AM NOT DONE
//
// charCounts считает символы строки. Для кириллицы в map попадают
// отдельные байты UTF-8, и символы считаются неверно.
// Тренирует: индексирование строки даёт байты, а не руны.
// Сложность: hard
package main_test

import "testing"

func charCounts(s string) map[byte]int {
	m := map[byte]int{}
	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}
	return m
}

func TestCharCounts(t *testing.T) {
	m := charCounts("ёжик")
	if len(m) != 4 || m['ж'] != 1 {
		t.Errorf("charCounts = %v", m)
	}
}
