// range58
// Make the tests pass!

// I AM NOT DONE
//
// chars возвращает символы строки как срез строк.
// Для кириллицы получаются обрывки байтов.
// Тренирует: s[i:i+1] берёт один байт, а не символ.
// Сложность: hard
package main_test

import (
	"reflect"
	"testing"
)

func chars(s string) []string {
	var out []string
	for i, r := range s {
		out = append(out, s[i:i+1])
		_ = r
	}
	return out
}

func TestChars(t *testing.T) {
	if got := chars("да!"); !reflect.DeepEqual(got, []string{"д", "а", "!"}) {
		t.Errorf("chars = %q", got)
	}
}
