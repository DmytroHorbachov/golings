// structs76
// Make the tests pass!

// I AM NOT DONE
//
// grow увеличивает радиус круга, хранящегося в интерфейсе.
// Код не компилируется: результат утверждения типа не адресуем.
// Тренирует: значение внутри интерфейса нельзя изменить на месте.
// Сложность: hard
package main_test

import "testing"

type Circle struct{ R float64 }

func grow(s interface{}) {
	s.(Circle).R *= 2
}

func TestGrow(t *testing.T) {
	c := &Circle{R: 1.5}
	grow(c)
	if c.R != 3 {
		t.Errorf("R = %v, want 3", c.R)
	}
}
