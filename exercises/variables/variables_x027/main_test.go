// variables_x027: Затенение в цикле
// Make the tests pass!
// I AM NOT DONE
//
// Функция должна вернуть сумму всех положительных чисел и список промежуточных сумм.
// Промежуточные суммы растут, а итог почему-то равен нулю.
// Тренирует: затенение переменной внутри тела цикла.
// Сложность: hard
package main_test

import (
	"reflect"
	"testing"
)

func positiveSum(nums []int) (int, []int) {
	total := 0
	var steps []int
	for _, v := range nums {
		if v > 0 {
			total := total + v
			steps = append(steps, total)
		}
	}
	return total, steps
}

func TestPositiveSum(t *testing.T) {
	total, steps := positiveSum([]int{3, -1, 4, 5})
	if total != 12 {
		t.Errorf("total = %d, want 12", total)
	}
	if want := []int{3, 7, 12}; !reflect.DeepEqual(steps, want) {
		t.Errorf("steps = %v, want %v", steps, want)
	}
}
