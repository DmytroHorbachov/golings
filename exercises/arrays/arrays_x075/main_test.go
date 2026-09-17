// arrays_x075: nil-указатель на массив
// Make the tests pass!
// I AM NOT DONE
//
// firstOr возвращает первый элемент массива по указателю или def, если указатель nil.
// len от nil-указателя на массив не паникует, а вот индексация — да.
// Тренирует: len(p) для *[N]T — константа N, даже если p == nil.
// Сложность: hard
package main_test

import "testing"

func firstOr(p *[3]int, def int) int {
	if len(p) > 0 {
		return p[0]
	}
	return def
}

func TestFirstOr(t *testing.T) {
	a := [3]int{7, 8, 9}
	if got := firstOr(&a, -1); got != 7 {
		t.Errorf("firstOr(&a) = %d", got)
	}
	if got := firstOr(nil, -1); got != -1 {
		t.Errorf("firstOr(nil) = %d, want -1", got)
	}
}
