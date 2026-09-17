// concurrent_x054: Слияние одинаковых запросов
// Make the tests pass!
// I AM NOT DONE
//
// Group выполняет функцию один раз для одновременных запросов с одним ключом
// (как singleflight); остальные ждут и получают тот же результат.
// Тренирует: map ожиданий под мьютексом и WaitGroup на запрос.
// Сложность: medium
package main_test

import (
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

type call struct {
	wg  sync.WaitGroup
	val string
}

type Group struct {
	mu    sync.Mutex
	calls map[string]*call
}

func (g *Group) Do(key string, fn func() string) string {
	g.mu.Lock()
	if g.calls == nil {
		g.calls = map[string]*call{}
	}
	g.mu.Unlock()
	return fn()
}

func TestGroup(t *testing.T) {
	var g Group
	var runs int32
	release := make(chan struct{})
	fn := func() string {
		atomic.AddInt32(&runs, 1)
		<-release
		return "data"
	}
	var wg sync.WaitGroup
	results := make([]string, 5)
	for i := range results {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			results[i] = g.Do("k", fn)
		}(i)
	}
	time.Sleep(50 * time.Millisecond)
	close(release)
	wg.Wait()
	if atomic.LoadInt32(&runs) != 1 {
		t.Errorf("fn ran %d times", runs)
	}
	for _, r := range results {
		if r != "data" {
			t.Errorf("result = %q", r)
		}
	}
}
