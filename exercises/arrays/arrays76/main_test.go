// arrays76
// Make the tests pass!

// I AM NOT DONE
//
// pascalRow возвращает n-ю строку треугольника Паскаля в массиве [10]int
// (неиспользуемые элементы — нули).
// Тренирует: обновление массива на месте справа налево.
// Сложность: medium
package main_test

import "testing"

func pascalRow(n int) [10]int {
	var row [10]int
	row[1] = 1
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			row[j] += row[j-1]
		}
	}
	return row
}

func TestPascalRow(t *testing.T) {
	if got := pascalRow(4); got != [10]int{1, 4, 6, 4, 1} {
		t.Errorf("pascalRow(4) = %v", got)
	}
	if got := pascalRow(0); got != [10]int{1} {
		t.Errorf("pascalRow(0) = %v", got)
	}
}
