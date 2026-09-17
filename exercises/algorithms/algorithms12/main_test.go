// algorithms12
// Make the tests pass!

// I AM NOT DONE
//
// Паттерн: хэш-таблицы с вычисляемым ключом. Сгруппируйте слова-анаграммы.
// Порядок групп — по первому появлению, порядок слов в группе — как во входе.
// Сложность: medium. Ожидаемая асимптотика: O(n·k·log k) по времени, O(n·k) по памяти
package main_test

import (
	"reflect"
	"sort"
	"testing"
)

func groupAnagrams(words []string) [][]string {
	return nil
}

func TestGroupAnagrams(t *testing.T) {
	_ = sort.Ints
	got := groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})
	want := [][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("groupAnagrams = %v, want %v", got, want)
	}
	if got := groupAnagrams([]string{""}); !reflect.DeepEqual(got, [][]string{{""}}) {
		t.Errorf("groupAnagrams([\"\"]) = %v", got)
	}
	if got := groupAnagrams(nil); len(got) != 0 {
		t.Errorf("groupAnagrams(nil) = %v", got)
	}
}
