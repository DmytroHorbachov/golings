// concurrent_x053: Параллельная сумма по частям
// Make the tests pass!
// I AM NOT DONE
//
// parallelSum делит срез на части, суммирует их в горутинах и складывает итог.
// Тренирует: разбиение работы и сбор результатов через канал.
// Сложность: medium
package main_test

import "testing"

func parallelSum(nums []int, parts int) int {
	size := (len(nums) + parts - 1) / parts
	ch := make(chan int, parts)
	count := 0
	for start := 0; start < len(nums); start += size {
		end := start + size
		go func(part []int) {
			s := 0
			for _, v := range part {
				s += v
			}
			ch <- s
		}(nums[start:end])
		count++
	}
	return <-ch
}

func TestParallelSum(t *testing.T) {
	nums := make([]int, 101)
	for i := range nums {
		nums[i] = i
	}
	if got := parallelSum(nums, 4); got != 5050 {
		t.Errorf("parallelSum = %d", got)
	}
}
