// range_x081: Некорректный UTF-8
// Make the tests pass!
// I AM NOT DONE
//
// validRunes считает символы строки и число некорректных байтов.
// range заменяет некорректные байты на U+FFFD, и их надо распознать.
// Тренирует: при ошибке декодирования range выдаёт utf8.RuneError.
// Сложность: hard
package main_test

import (
	"testing"
	"unicode/utf8"
)

func validRunes(s string) (ok, bad int) {
	for range s {
		ok++
	}
	return
}

func TestValidRunes(t *testing.T) {
	_ = utf8.RuneError
	ok, bad := validRunes("a\xffб�\xfe")
	if ok != 3 || bad != 2 {
		t.Errorf("validRunes = %d, %d; want 3, 2", ok, bad)
	}
}
