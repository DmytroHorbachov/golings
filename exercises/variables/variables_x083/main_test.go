// variables_x083: Шифр Цезаря
// Make the tests pass!
// I AM NOT DONE
//
// Функция shift должна сдвинуть строчную латинскую букву на k позиций по кругу.
// Сейчас буквы в конце алфавита превращаются в посторонние символы.
// Тренирует: арифметику над byte и остаток от деления.
// Сложность: medium
package main_test

import "testing"

func shift(c byte, k int) byte {
	return c + byte(k)
}

func TestShift(t *testing.T) {
	cases := []struct {
		c    byte
		k    int
		want byte
	}{{'a', 1, 'b'}, {'x', 3, 'a'}, {'z', 27, 'a'}, {'m', 0, 'm'}}
	for _, c := range cases {
		if got := shift(c.c, c.k); got != c.want {
			t.Errorf("shift(%c, %d) = %c, want %c", c.c, c.k, got, c.want)
		}
	}
}
