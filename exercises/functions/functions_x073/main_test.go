// functions_x073: Значение метода
// Make the tests pass!
// I AM NOT DONE
//
// Функция snapshotter возвращает метод-значение, который должен
// показывать АКТУАЛЬНОЕ значение счётчика.
// Тренирует: метод-значение с value-receiver копирует получатель в момент вычисления.
// Сложность: hard
package main_test

import "testing"

type Counter struct{ n int }

func (c *Counter) Inc() { c.n++ }

func (c Counter) Value() int { return c.n }

func TestMethodValue(t *testing.T) {
	c := &Counter{}
	value := c.Value
	c.Inc()
	c.Inc()
	if got := value(); got != 2 {
		t.Errorf("value() = %d, want 2", got)
	}
}
