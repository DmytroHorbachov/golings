// structs102
// Make the tests pass!

// I AM NOT DONE
//
// Point.String форматирует точку как "(x, y)", и fmt использует этот метод.
// Тренирует: интерфейс fmt.Stringer.
// Сложность: easy
package main_test

import (
	"fmt"
	"testing"
)

type Point struct{ X, Y int }

func (p Point) String() string {
	return fmt.Sprintf("[%d; %d]", p.X, p.Y)
}

func TestPointString(t *testing.T) {
	if got := fmt.Sprint(Point{1, 2}); got != "(1, 2)" {
		t.Errorf("Sprint = %q", got)
	}
}
