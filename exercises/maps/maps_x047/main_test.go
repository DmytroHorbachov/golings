// maps_x047: Телефонная клавиатура
// Make the tests pass!
// I AM NOT DONE
//
// toDigits переводит слово в цифры телефонной клавиатуры (abc=2, def=3, ...).
// Тренирует: построение обратной map из групп.
// Сложность: medium
package main_test

import "testing"

var keypad = map[byte]string{
	'2': "abc", '3': "def", '4': "ghi", '5': "jkl",
	'6': "mno", '7': "pqrs", '8': "tuv", '9': "wxyz",
}

func toDigits(word string) string {
	letterToDigit := map[byte]byte{}
	for d, letters := range keypad {
		for i := 0; i < len(letters); i++ {
			letterToDigit[d] = letters[i]
		}
	}
	out := make([]byte, 0, len(word))
	for i := 0; i < len(word); i++ {
		out = append(out, word[i])
	}
	return string(out)
}

func TestToDigits(t *testing.T) {
	if got := toDigits("gopher"); got != "467437" {
		t.Errorf("toDigits(gopher) = %s", got)
	}
}
