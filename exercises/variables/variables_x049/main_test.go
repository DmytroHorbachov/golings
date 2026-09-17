// variables_x049: Порядок в кортежном присваивании
// Make the tests pass!
// I AM NOT DONE
//
// Функция pushAfter сдвигает курсор и записывает значение в НОВУЮ позицию.
// Сейчас значение пишется в старую позицию.
// Тренирует: индексные выражения слева вычисляются до присваивания.
// Сложность: hard
package main_test

import (
	"reflect"
	"testing"
)

func pushAfter(buf []int, pos, v int) ([]int, int) {
	pos, buf[pos] = pos+1, v
	return buf, pos
}

func TestPushAfter(t *testing.T) {
	buf := make([]int, 4)
	pos := 0
	buf, pos = pushAfter(buf, pos, 7)
	buf, pos = pushAfter(buf, pos, 9)
	if pos != 2 {
		t.Errorf("pos = %d, want 2", pos)
	}
	if want := []int{0, 7, 9, 0}; !reflect.DeepEqual(buf, want) {
		t.Errorf("buf = %v, want %v", buf, want)
	}
}
