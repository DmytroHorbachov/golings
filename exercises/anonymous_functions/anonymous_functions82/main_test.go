// anonymous_functions82
// Make the tests pass!

// I AM NOT DONE
//
// Функцию greet декорируют, переприсваивая переменную литералом, который
// вызывает greet. Литерал вызывает сам себя, а не исходную функцию.
// Тренирует: литерал захватывает переменную, а не значение, которое в ней было.
// Сложность: hard
package main_test

import (
	"strings"
	"testing"
)

func decorated() func(string) string {
	greet := func(name string) string { return "hello, " + name }
	depth := 0
	greet = func(name string) string {
		depth++
		if depth > 3 {
			return "recursion"
		}
		return strings.ToUpper(greet(name))
	}
	return greet
}

func TestDecorated(t *testing.T) {
	if got := decorated()("go"); got != "HELLO, GO" {
		t.Errorf("decorated = %q", got)
	}
}
