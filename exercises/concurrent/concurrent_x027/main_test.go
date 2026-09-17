// concurrent_x027: Защита map
// Make the tests pass!
// I AM NOT DONE
//
// Store записывает в map из многих горутин, но блокировка забыта.
// Тренирует: мьютекс для записи в map.
// Сложность: medium
package main_test

import (
	"strconv"
	"sync"
	"testing"
)

type Store struct {
	mu   sync.Mutex
	data map[string]int
}

func (s *Store) Set(k string, v int) {
	s.data[k] = v
}

func TestStore(t *testing.T) {
	s := &Store{data: map[string]int{}}
	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			s.Set(strconv.Itoa(i), i)
		}(i)
	}
	wg.Wait()
	if len(s.data) != 50 {
		t.Errorf("len = %d", len(s.data))
	}
}
