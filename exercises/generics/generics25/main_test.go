// generics25
// Make the tests pass!

// I AM NOT DONE
//
// NewAll создаёт n «сбрасываемых» объектов типа T. Для T = *Buffer
// нулевое значение — nil, и вызов Reset паникует.
// Тренирует: нулевое значение типа-указателя не является готовым объектом.
// Сложность: hard
package main_test

import "testing"

type Buffer struct{ data []byte }

func (b *Buffer) Reset() { b.data = b.data[:0] }

type Resetter interface{ Reset() }

func NewAll[T Resetter](n int) []T {
	out := make([]T, n)
	for i := range out {
		out[i].Reset()
	}
	return out
}

func buffers() []*Buffer { return NewAll[*Buffer](2) }

func TestBuffers(t *testing.T) {
	bs := buffers()
	if len(bs) != 2 || bs[0] == nil || bs[1] == nil || bs[0] == bs[1] {
		t.Errorf("buffers = %v", bs)
	}
}
