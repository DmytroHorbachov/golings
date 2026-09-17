// generics50
// Make the tests pass!

// I AM NOT DONE
//
// JoinAll вызывает String у всех элементов. Для []Point код не компилируется:
// String объявлен на *Point, и Point не удовлетворяет ограничению.
// Тренирует: набор методов типа-аргумента должен содержать методы ограничения.
// Сложность: hard
package main_test

import (
	"fmt"
	"strings"
	"testing"
)

type Stringer interface{ String() string }

type Point struct{ X, Y int }

func (p *Point) String() string { return fmt.Sprintf("(%d,%d)", p.X, p.Y) }

func JoinAll[T Stringer](items []T) string {
	parts := make([]string, len(items))
	for i, it := range items {
		parts[i] = it.String()
	}
	return strings.Join(parts, " ")
}

func TestJoinAll(t *testing.T) {
	if got := JoinAll([]Point{{1, 2}, {3, 4}}); got != "(1,2) (3,4)" {
		t.Errorf("JoinAll = %q", got)
	}
}
