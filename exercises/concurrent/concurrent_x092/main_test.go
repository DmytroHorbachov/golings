// concurrent_x092: append из горутин
// Make the tests pass!
// I AM NOT DONE
//
// gather добавляет результаты в общий срез из многих горутин;
// часть результатов теряется.
// Тренирует: append не потокобезопасен.
// Сложность: hard
package main_test

import (
	"sync"
	"testing"
)

func gather(n int) []int {
	var out []int
	var mu sync.Mutex
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			out = append(out, i)
		}(i)
	}
	wg.Wait()
	_ = &mu
	return out
}

func TestGather(t *testing.T) {
	if got := gather(300); len(got) != 300 {
		t.Errorf("len = %d, want 300", len(got))
	}
}
