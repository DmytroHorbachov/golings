// maps_x040: Вложенная map по требованию
// Make the tests pass!
// I AM NOT DONE
//
// addScore записывает оценку ученика по предмету в map[ученик]map[предмет]оценка.
// Для нового ученика запись паникует.
// Тренирует: инициализацию вложенной map.
// Сложность: medium
package main_test

import "testing"

func addScore(book map[string]map[string]int, student, subject string, score int) {
	book[student][subject] = score
}

func TestAddScore(t *testing.T) {
	book := map[string]map[string]int{}
	addScore(book, "ann", "math", 5)
	addScore(book, "ann", "art", 4)
	addScore(book, "bob", "math", 3)
	if book["ann"]["math"] != 5 || book["ann"]["art"] != 4 || book["bob"]["math"] != 3 {
		t.Errorf("book = %v", book)
	}
}
