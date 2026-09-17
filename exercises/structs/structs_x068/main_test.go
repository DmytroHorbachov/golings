// structs_x068: Набор методов значения
// Make the tests pass!
// I AM NOT DONE
//
// Counter должен удовлетворять интерфейсу Incrementer. Код не компилируется:
// метод Inc объявлен на указателе, а в интерфейс кладут значение.
// Тренирует: набор методов типа T не включает методы *T.
// Сложность: hard
package main_test

import "testing"

type Incrementer interface{ Inc() int }

type Counter struct{ n int }

func (c *Counter) Inc() int {
	c.n++
	return c.n
}

func newIncrementer() Incrementer {
	return Counter{}
}

func TestIncrementer(t *testing.T) {
	i := newIncrementer()
	i.Inc()
	if got := i.Inc(); got != 2 {
		t.Errorf("Inc = %d, want 2", got)
	}
}
