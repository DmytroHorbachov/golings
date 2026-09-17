// slices_x070: Удаление внутри функции
// Make the tests pass!
// I AM NOT DONE
//
// removeFirst удаляет первый элемент со значением x. Вызывающий код продолжает
// использовать старый срез и видит «хвост» из дубликатов.
// Тренирует: функция может изменить элементы, но не длину среза вызывающего.
// Сложность: hard
package main_test

import (
	"reflect"
	"testing"
)

func removeFirst(s []int, x int) {
	for i, v := range s {
		if v == x {
			s = append(s[:i], s[i+1:]...)
			return
		}
	}
}

func cleanup(s []int) []int {
	removeFirst(s, 0)
	return s
}

func TestCleanup(t *testing.T) {
	if got := cleanup([]int{1, 0, 2, 3}); !reflect.DeepEqual(got, []int{1, 2, 3}) {
		t.Errorf("cleanup = %v, want [1 2 3]", got)
	}
}
