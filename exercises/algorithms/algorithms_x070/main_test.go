// algorithms_x070: Word Search II (поиск слов на доске)
// Make the tests pass!
// I AM NOT DONE
//
// Паттерн: префиксное дерево и поиск с возвратом. Найдите все слова из словаря,
// которые можно собрать на доске, двигаясь по соседним клеткам (без повторного
// использования клетки). Ответ отсортирован по алфавиту.
// Сложность: hard. Ожидаемая асимптотика: O(r·c·4^L) по времени, O(общая длина слов) по памяти
package main_test

import (
	"reflect"
	"sort"
	"testing"
)

func findWords(board [][]byte, words []string) []string {
	return nil
}

func TestFindWords(t *testing.T) {
	_ = sort.Strings
	board := [][]byte{
		[]byte("oaan"),
		[]byte("etae"),
		[]byte("ihkr"),
		[]byte("iflv"),
	}
	got := findWords(board, []string{"oath", "pea", "eat", "rain"})
	if !reflect.DeepEqual(got, []string{"eat", "oath"}) {
		t.Errorf("findWords = %v", got)
	}
	if got := findWords([][]byte{[]byte("ab")}, []string{"ba", "ab", "abc"}); !reflect.DeepEqual(got, []string{"ab", "ba"}) {
		t.Errorf("small board = %v", got)
	}
	if got := findWords(nil, []string{"a"}); len(got) != 0 {
		t.Errorf("empty board = %v", got)
	}
}
