// concurrent_x002: Ожидание горутин
// Make the tests pass!
// I AM NOT DONE
//
// collect запускает горутины, заполняющие срез, и сразу возвращает результат.
// Тренирует: wg.Wait перед чтением результатов.
// Сложность: easy
package main_test

import (
	"reflect"
	"sync"
	"testing"
)

func collect(n int) []int {
	out := make([]int, n)
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			out[i] = i * 10
		}(i)
	}
	_ = &wg
	return out
}

func TestCollect(t *testing.T) {
	if got := collect(3); !reflect.DeepEqual(got, []int{0, 10, 20}) {
		t.Errorf("collect = %v", got)
	}
}
