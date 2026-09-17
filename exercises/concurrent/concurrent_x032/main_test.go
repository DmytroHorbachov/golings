// concurrent_x032: Параметр горутины
// Make the tests pass!
// I AM NOT DONE
//
// Каждая горутина должна записать свой номер. Номер передаётся параметром.
// Тренирует: передача значения в go func(v int){...}(v).
// Сложность: easy
package main_test

import (
	"sort"
	"sync"
	"testing"
)

func ids(n int) []int {
	var mu sync.Mutex
	var out []int
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			mu.Lock()
			out = append(out, id)
			mu.Unlock()
		}(n)
	}
	wg.Wait()
	sort.Ints(out)
	return out
}

func TestIDs(t *testing.T) {
	got := ids(3)
	if len(got) != 3 || got[0] != 0 || got[2] != 2 {
		t.Errorf("ids = %v", got)
	}
}
