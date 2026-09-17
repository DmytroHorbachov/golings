// generics_x075: Специализация метода
// Make the tests pass!
// I AM NOT DONE
//
// Container[T].Format должен форматировать int-значения особым образом.
// Код не компилируется: нельзя объявить метод только для Container[int].
// Тренирует: методы объявляются для всех инстанциаций сразу.
// Сложность: hard
package main_test

import (
	"fmt"
	"testing"
)

type Container[T any] struct{ V T }

func (c Container[T]) Format() string { return fmt.Sprint(c.V) }

func (c Container[int]) Format() string { return fmt.Sprintf("#%d", c.V) }

func TestFormat(t *testing.T) {
	if (Container[int]{7}).Format() != "#7" || (Container[string]{"x"}).Format() != "x" {
		t.Errorf("Format works incorrectly")
	}
}
