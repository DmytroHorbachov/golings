// range_x051: Чтение с проверкой закрытия
// Make the tests pass!
// I AM NOT DONE
//
// firstN читает из канала не больше n значений и сообщает, был ли канал
// закрыт раньше.
// Тренирует: получение v, ok := <-ch в цикле как альтернатива range.
// Сложность: medium
package main_test

import (
	"reflect"
	"testing"
)

func firstN(ch <-chan int, n int) ([]int, bool) {
	var out []int
	for v := range ch {
		out = append(out, v)
	}
	return out, true
}

func TestFirstN(t *testing.T) {
	ch := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		ch <- i
	}
	got, closed := firstN(ch, 3)
	if !reflect.DeepEqual(got, []int{1, 2, 3}) || closed {
		t.Errorf("firstN = %v, %v", got, closed)
	}
	close(ch)
	got, closed = firstN(ch, 5)
	if !reflect.DeepEqual(got, []int{4, 5}) || !closed {
		t.Errorf("firstN after close = %v, %v", got, closed)
	}
}
