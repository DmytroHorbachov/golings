// structs_x064: Цепочка операций с векторами
// Make the tests pass!
// I AM NOT DONE
//
// Vec.Add и Vec.Scale возвращают новые векторы, позволяя писать цепочки.
// Тренирует: значимые получатели для неизменяемых операций.
// Сложность: medium
package main_test

import "testing"

type Vec struct{ X, Y int }

func (v Vec) Add(o Vec) Vec { return Vec{v.X + o.X, v.Y + o.Y} }

func (v Vec) Scale(k int) Vec {
	v.X *= k
	return Vec{v.X, v.Y + k}
}

func TestVec(t *testing.T) {
	a := Vec{1, 2}
	got := a.Add(Vec{1, 1}).Scale(3)
	if got != (Vec{6, 9}) || a != (Vec{1, 2}) {
		t.Errorf("got %v, a = %v", got, a)
	}
}
