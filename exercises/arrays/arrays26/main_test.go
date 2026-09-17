// arrays26
// Make the tests pass!

// I AM NOT DONE
//
// Ring хранит последние 3 значения. Push добавляет значение, вытесняя самое старое;
// Items возвращает значения от старого к новому.
// Тренирует: фиксированный массив и индекс по модулю.
// Сложность: medium
package main_test

import (
	"reflect"
	"testing"
)

type Ring struct {
	buf   [3]int
	pos   int
	count int
}

func (r *Ring) Push(v int) {
	r.buf[r.pos] = v
	r.pos++
}

func (r *Ring) Items() []int {
	out := make([]int, 0, r.count)
	start := (r.pos - r.count + len(r.buf)) % len(r.buf)
	for i := 0; i < r.count; i++ {
		out = append(out, r.buf[(start+i)%len(r.buf)])
	}
	return out
}

func TestRing(t *testing.T) {
	var r Ring
	for _, v := range []int{1, 2, 3, 4, 5} {
		r.Push(v)
	}
	if got := r.Items(); !reflect.DeepEqual(got, []int{3, 4, 5}) {
		t.Errorf("Items = %v, want [3 4 5]", got)
	}
}
