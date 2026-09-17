// variables_x028: Переполнение int8
// Make the tests pass!
// I AM NOT DONE
//
// Функция должна сложить показания датчиков, каждое в пределах int8.
// Для больших сумм результат внезапно становится отрицательным.
// Тренирует: переполнение целых чисел фиксированного размера.
// Сложность: hard
package main_test

import "testing"

func totalReading(values []int8) int {
	var sum int8
	for _, v := range values {
		sum += v
	}
	return int(sum)
}

func TestTotalReading(t *testing.T) {
	if got := totalReading([]int8{100, 100, 100}); got != 300 {
		t.Errorf("totalReading(100,100,100) = %d, want 300", got)
	}
	if got := totalReading([]int8{-128, -128}); got != -256 {
		t.Errorf("totalReading(-128,-128) = %d, want -256", got)
	}
}
