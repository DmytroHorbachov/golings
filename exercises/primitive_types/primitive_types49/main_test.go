// primitive_types49
// Make the tests pass!

// I AM NOT DONE
//
// countRunes должна вернуть число символов или ошибку, если строка содержит
// некорректные байты UTF-8. Сейчас мусорные байты считаются символами.
// Тренирует: range по невалидной строке выдаёт utf8.RuneError.
// Сложность: hard
package main_test

import (
	"errors"
	"testing"
	"unicode/utf8"
)

func countRunes(s string) (int, error) {
	return utf8.RuneCountInString(s), nil
}

func TestCountRunes(t *testing.T) {
	_ = errors.New
	if n, err := countRunes("héllo"); err != nil || n != 5 {
		t.Errorf("countRunes(héllo) = %d, %v", n, err)
	}
	if _, err := countRunes("bad\xff\xfe"); err == nil {
		t.Errorf("countRunes with invalid bytes should fail")
	}
}
