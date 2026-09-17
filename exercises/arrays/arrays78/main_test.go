// arrays78
// Make the tests pass!

// I AM NOT DONE
//
// average принимает массив [3]int, но её вызывают для массива из 4 элементов.
// Код не компилируется.
// Тренирует: длина — часть типа массива.
// Сложность: hard
package main_test

import "testing"

func average(a [3]int) float64 {
	s := 0
	for _, v := range a {
		s += v
	}
	return float64(s) / float64(len(a))
}

func averages() (float64, float64) {
	three := [3]int{1, 2, 3}
	four := [4]int{1, 2, 3, 6}
	return average(three), average(four)
}

func TestAverages(t *testing.T) {
	a, b := averages()
	if a != 2 || b != 3 {
		t.Errorf("averages = %v, %v", a, b)
	}
}
