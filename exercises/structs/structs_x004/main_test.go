// structs_x004: Указатель на структуру
// Make the tests pass!
// I AM NOT DONE
//
// moveRight сдвигает точку по указателю на dx.
// Тренирует: автоматическое разыменование при доступе к полям.
// Сложность: easy
package main_test

import "testing"

type Point struct{ X, Y int }

func moveRight(p *Point, dx int) {
	p.Y += dx
}

func TestMoveRight(t *testing.T) {
	p := Point{1, 1}
	moveRight(&p, 4)
	if p != (Point{5, 1}) {
		t.Errorf("p = %+v", p)
	}
}
