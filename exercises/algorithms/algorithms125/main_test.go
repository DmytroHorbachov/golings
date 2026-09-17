// algorithms125
// Make the tests pass!

// I AM NOT DONE
//
// Паттерн: словарь и динамическое программирование по строке. Найдите слова,
// которые целиком составляются из двух или более других слов словаря.
// Ответ отсортирован по алфавиту.
// Сложность: hard. Ожидаемая асимптотика: O(n·L²) по времени, O(n·L) по памяти
package main_test

import (
	"reflect"
	"sort"
	"testing"
)

func concatenatedWords(words []string) []string {
	return nil
}

func TestConcatenatedWords(t *testing.T) {
	_ = sort.Strings
	got := concatenatedWords([]string{"cat", "cats", "catsdogcats", "dog", "dogcatsdog", "hippopotamuses", "rat", "ratcatdogcat"})
	want := []string{"catsdogcats", "dogcatsdog", "ratcatdogcat"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("concatenatedWords = %v, want %v", got, want)
	}
	if got := concatenatedWords([]string{"cat", "dog"}); len(got) != 0 {
		t.Errorf("no concatenated words expected, got %v", got)
	}
	if got := concatenatedWords([]string{"a", "aa", "aaa"}); !reflect.DeepEqual(got, []string{"aa", "aaa"}) {
		t.Errorf("repeated letters = %v", got)
	}
}
