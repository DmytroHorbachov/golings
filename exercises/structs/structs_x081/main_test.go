// structs_x081: Продвижение методов указателя
// Make the tests pass!
// I AM NOT DONE
//
// Service встраивает Counter (значением), у которого Inc объявлен на указателе.
// Значение Service не реализует интерфейс, а указатель — реализует.
// Тренирует: методы *T встроенного T продвигаются только в *Outer.
// Сложность: hard
package main_test

import "testing"

type Counter struct{ n int }

func (c *Counter) Inc() { c.n++ }

type Service struct {
	Counter
	Name string
}

type Incer interface{ Inc() }

func bump(i Incer, times int) {
	for k := 0; k < times; k++ {
		i.Inc()
	}
}

func run() int {
	s := Service{Name: "api"}
	bump(s, 3)
	return s.n
}

func TestRun(t *testing.T) {
	if got := run(); got != 3 {
		t.Errorf("run = %d, want 3", got)
	}
}
