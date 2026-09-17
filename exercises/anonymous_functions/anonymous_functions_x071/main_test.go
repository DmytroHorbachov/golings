// anonymous_functions_x071: go без вызова
// Make the tests pass!
// I AM NOT DONE
//
// Фоновая задача должна отправить результат в канал. Код не компилируется.
// Тренирует: выражение в go должно быть вызовом функции.
// Сложность: hard
package main_test

import "testing"

func background() int {
	ch := make(chan int, 1)
	go func() {
		ch <- 42
	}
	return <-ch
}

func TestBackground(t *testing.T) {
	if background() != 42 {
		t.Errorf("background = %d", background())
	}
}
