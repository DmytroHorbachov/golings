// variables_x048: Цикл затеняет счётчик
// Make the tests pass!
// I AM NOT DONE
//
// Функция indexOf должна вернуть индекс первого вхождения символа или -1.
// Сейчас она всегда возвращает начальное значение i.
// Тренирует: := в заголовке for создаёт новую переменную.
// Сложность: hard
package main_test

import "testing"

func indexOf(s string, c byte) int {
	i := 0
	for i := 0; i < len(s); i++ {
		if s[i] == c {
			break
		}
	}
	if i == len(s) {
		return -1
	}
	return i
}

func TestIndexOf(t *testing.T) {
	cases := []struct {
		s    string
		c    byte
		want int
	}{{"golang", 'l', 2}, {"golang", 'g', 0}, {"golang", 'x', -1}, {"", 'a', -1}}
	for _, c := range cases {
		if got := indexOf(c.s, c.c); got != c.want {
			t.Errorf("indexOf(%q, %q) = %d, want %d", c.s, c.c, got, c.want)
		}
	}
}
