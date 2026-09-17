// variables63
// Make the tests pass!

// I AM NOT DONE
//
// Функции должны вернуть значения с типами int и float64 соответственно.
// Тип переменной выводится из литерала, и сейчас он не тот.
// Тренирует: вывод типа при := из нетипизированных констант.
// Сложность: medium
package main_test

import (
	"fmt"
	"testing"
)

func answer() interface{} {
	v := 42.0
	return v
}

func ratio() interface{} {
	v := 1 / 2
	return v
}

func TestTypes(t *testing.T) {
	if got := fmt.Sprintf("%T %v", answer(), answer()); got != "int 42" {
		t.Errorf("answer() = %s, want int 42", got)
	}
	if got := fmt.Sprintf("%T %v", ratio(), ratio()); got != "float64 0.5" {
		t.Errorf("ratio() = %s, want float64 0.5", got)
	}
}
