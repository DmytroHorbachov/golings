// concurrent49
// Make the tests pass!

// I AM NOT DONE
//
// process раздаёт задачи n воркерам и собирает результаты (квадраты чисел).
// Тренирует: каналы jobs/results и закрытие results после завершения воркеров.
// Сложность: medium
package main_test

import (
	"sort"
	"sync"
	"testing"
	"time"
)

func process(nums []int, workers int) []int {
	jobs := make(chan int)
	results := make(chan int)
	var wg sync.WaitGroup
	for w := 0; w < workers; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for n := range jobs {
				results <- n * n
			}
		}()
	}
	go func() {
		for _, n := range nums {
			jobs <- n
		}
	}()
	var out []int
	for r := range results {
		out = append(out, r)
	}
	sort.Ints(out)
	return out
}

func TestProcess(t *testing.T) {
	done := make(chan []int, 1)
	go func() { done <- process([]int{3, 1, 2}, 2) }()
	select {
	case got := <-done:
		if len(got) != 3 || got[0] != 1 || got[2] != 9 {
			t.Errorf("process = %v", got)
		}
	case <-time.After(2 * time.Second):
		t.Fatal("process never finished")
	}
}
