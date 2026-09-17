// primitive_types34
// Make the tests pass!

// I AM NOT DONE
//
// digitSum должна сложить все цифры в строке, игнорируя прочие символы.
// Тренирует: преобразование байта-цифры в число.
// Сложность: medium
package main_test

import "testing"

func digitSum(s string) int {
	sum := 0
	for i := 0; i < len(s); i++ {
		sum += int(s[i])
	}
	return sum
}

func TestDigitSum(t *testing.T) {
	cases := map[string]int{"123": 6, "a1b2c3": 6, "": 0, "9-9": 18}
	for in, want := range cases {
		if got := digitSum(in); got != want {
			t.Errorf("digitSum(%q) = %d, want %d", in, got, want)
		}
	}
}
