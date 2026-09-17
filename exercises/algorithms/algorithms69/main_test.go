// algorithms69
// Make the tests pass!

// I AM NOT DONE
//
// Паттерн: поиск в глубину с возвратом. Проверьте, можно ли собрать слово,
// двигаясь по соседним клеткам без повторного использования клетки.
// Сложность: medium. Ожидаемая асимптотика: O(r·c·4^L) по времени, O(L) по памяти
package main_test

import "testing"

func exist(board [][]byte, word string) bool {
	return false
}

func TestExist(t *testing.T) {
	board := [][]byte{
		[]byte("ABCE"),
		[]byte("SFCS"),
		[]byte("ADEE"),
	}
	cases := map[string]bool{"ABCCED": true, "SEE": true, "ABCB": false, "": true}
	for word, want := range cases {
		if got := exist(board, word); got != want {
			t.Errorf("exist(%q) = %v, want %v", word, got, want)
		}
	}
	if exist(nil, "a") {
		t.Errorf("empty board should not contain words")
	}
}
