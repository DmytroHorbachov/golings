// arrays_x050: Самая частая буква
// Make the tests pass!
// I AM NOT DONE
//
// topLetter возвращает самую частую латинскую букву строки (при равенстве — первую по алфавиту).
// Тренирует: подсчёт в [26]int и поиск максимума.
// Сложность: medium
package main_test

import "testing"

func topLetter(s string) byte {
	var counts [26]int
	for i := 0; i < len(s); i++ {
		if s[i] >= 'a' && s[i] <= 'z' {
			counts[s[i]-'a']++
		}
	}
	best := 0
	for i, c := range counts {
		if c >= best {
			best = i
		}
	}
	return byte(best)
}

func TestTopLetter(t *testing.T) {
	cases := map[string]byte{"banana": 'a', "zzyy": 'y', "hello world": 'l'}
	for in, want := range cases {
		if got := topLetter(in); got != want {
			t.Errorf("topLetter(%q) = %c, want %c", in, got, want)
		}
	}
}
