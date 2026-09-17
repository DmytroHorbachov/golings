// anonymous_functions_x062: Параллельное отображение
// Make the tests pass!
// I AM NOT DONE
//
// parallelMap применяет f к каждому элементу в отдельной горутине и сохраняет
// порядок результатов.
// Тренирует: горутины-литералы с передачей индекса аргументом.
// Сложность: medium
package main_test

import (
	"reflect"
	"sync"
	"testing"
)

func parallelMap(in []int, f func(int) int) []int {
	out := make([]int, len(in))
	var wg sync.WaitGroup
	for i, v := range in {
		wg.Add(1)
		go func() {
			defer wg.Done()
			out = append(out, f(v))
		}()
	}
	wg.Wait()
	return out
}

func TestParallelMap(t *testing.T) {
	got := parallelMap([]int{1, 2, 3, 4}, func(x int) int { return x * 10 })
	if !reflect.DeepEqual(got, []int{10, 20, 30, 40}) {
		t.Errorf("parallelMap = %v", got)
	}
}
