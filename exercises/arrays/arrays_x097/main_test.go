// arrays_x097: Защита от перезаписи
// Make the tests pass!
// I AM NOT DONE
//
// head возвращает срез первых двух элементов массива. Вызывающий код делает
// append к этому срезу — и перезаписывает третий элемент массива.
// Тренирует: полное выражение среза a[low:high:max] ограничивает ёмкость.
// Сложность: hard
package main_test

import "testing"

func head(a *[3]int) []int {
	return a[0:2]
}

func TestHead(t *testing.T) {
	a := [3]int{1, 2, 3}
	h := head(&a)
	h = append(h, 99)
	if a[2] != 3 {
		t.Errorf("array overwritten: %v", a)
	}
	if len(h) != 3 || h[2] != 99 {
		t.Errorf("h = %v", h)
	}
}
