// concurrent_x069: Конкурентная запись в map
// Make the tests pass!
// I AM NOT DONE
//
// index заполняет map из нескольких горутин; программа падает с
// «concurrent map writes» или гонкой.
// Тренирует: map не безопасна для конкурентной записи.
// Сложность: hard
package main_test

import (
	"strconv"
	"sync"
	"testing"
)

func index(n int) map[string]int {
	m := map[string]int{}
	var mu sync.Mutex
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			m[strconv.Itoa(i)] = i
		}(i)
	}
	wg.Wait()
	_ = &mu
	return m
}

func TestIndex(t *testing.T) {
	if got := index(300); len(got) != 300 {
		t.Errorf("len = %d", len(got))
	}
}
