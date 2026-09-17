// variables_x097: Статистика за один проход
// Make the tests pass!
// I AM NOT DONE
//
// Функция stats должна вернуть количество, сумму и среднее значение.
// Для пустого среза среднее должно быть 0.
// Тренирует: несколько переменных-аккумуляторов и деление без паники.
// Сложность: medium
package main_test

import "testing"

func stats(nums []int) (count, sum int, mean float64) {
	for _, v := range nums {
		count++
		sum += v
	}
	mean = float64(sum / count)
	return
}

func TestStats(t *testing.T) {
	c, s, m := stats([]int{1, 2, 4})
	if c != 3 || s != 7 || m < 2.33 || m > 2.34 {
		t.Errorf("stats(1,2,4) = %d, %d, %v", c, s, m)
	}
	c, s, m = stats(nil)
	if c != 0 || s != 0 || m != 0 {
		t.Errorf("stats(nil) = %d, %d, %v; want 0, 0, 0", c, s, m)
	}
}
