// generics_x072: Тип только в результате
// Make the tests pass!
// I AM NOT DONE
//
// Parse[T] создаёт значение нужного типа из строки. Вызов без аргумента типа
// не компилируется: компилятор не выводит T по возвращаемому значению.
// Тренирует: вывод типов работает только по аргументам.
// Сложность: hard
package main_test

import (
	"fmt"
	"testing"
)

func Parse[T any](s string) (T, error) {
	var v T
	_, err := fmt.Sscan(s, &v)
	return v, err
}

func port() int {
	p, _ := Parse("8080")
	return p
}

func TestPort(t *testing.T) {
	if port() != 8080 {
		t.Errorf("port = %d", port())
	}
}
