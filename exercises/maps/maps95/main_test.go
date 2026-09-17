// maps95
// Make the tests pass!

// I AM NOT DONE
//
// NewIndex создаёт индекс «язык -> слово -> частота». Добавление слова
// для нового языка паникует.
// Тренирует: make внешней map не создаёт внутренние.
// Сложность: hard
package main_test

import "testing"

type Index struct {
	data map[string]map[string]int
}

func NewIndex() *Index {
	return &Index{data: make(map[string]map[string]int)}
}

func (ix *Index) Add(lang, word string) {
	ix.data[lang][word]++
}

func TestIndex(t *testing.T) {
	ix := NewIndex()
	ix.Add("go", "func")
	ix.Add("go", "func")
	ix.Add("py", "def")
	if ix.data["go"]["func"] != 2 || ix.data["py"]["def"] != 1 {
		t.Errorf("index = %v", ix.data)
	}
}
