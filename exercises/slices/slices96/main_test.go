// slices96
// Make the tests pass!

// I AM NOT DONE
//
// page возвращает элементы страницы n (с единицы) размером size.
// Для страницы за пределами данных — пустой срез.
// Тренирует: вычисление границ среза и их ограничение.
// Сложность: medium
package main_test

import (
	"reflect"
	"testing"
)

func page(s []string, n, size int) []string {
	start := n * size
	end := start + size
	return s[start:end]
}

func TestPage(t *testing.T) {
	s := []string{"a", "b", "c", "d", "e"}
	if got := page(s, 1, 2); !reflect.DeepEqual(got, []string{"a", "b"}) {
		t.Errorf("page 1 = %v", got)
	}
	if got := page(s, 3, 2); !reflect.DeepEqual(got, []string{"e"}) {
		t.Errorf("page 3 = %v", got)
	}
	if got := page(s, 4, 2); len(got) != 0 {
		t.Errorf("page 4 = %v", got)
	}
}
