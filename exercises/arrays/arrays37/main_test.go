// arrays37
// Make the tests pass!

// I AM NOT DONE
//
// fill записывает значение во все элементы массива по указателю.
// Тренирует: индексацию через указатель на массив.
// Сложность: easy
package main_test

import "testing"

func fill(p *[4]int, v int) {
	for i := range p {
		p[i] = i
	}
}

func TestFill(t *testing.T) {
	var a [4]int
	fill(&a, 7)
	if a != [4]int{7, 7, 7, 7} {
		t.Errorf("fill = %v", a)
	}
}
