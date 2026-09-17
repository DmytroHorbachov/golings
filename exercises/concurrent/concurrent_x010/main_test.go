// concurrent_x010: sync.Once
// Make the tests pass!
// I AM NOT DONE
//
// Config загружается один раз, сколько бы горутин ни вызывали Get.
// Тренирует: sync.Once.Do.
// Сложность: easy
package main_test

import (
	"sync"
	"testing"
)

type Config struct {
	once  sync.Once
	loads int
	value string
}

func (c *Config) Get() string {
	func(f func()) { f() }(func() {
		c.loads++
		c.value = "ready"
	})
	return c.value
}

func TestConfigOnce(t *testing.T) {
	var c Config
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Get()
		}()
	}
	wg.Wait()
	if c.loads != 1 || c.Get() != "ready" {
		t.Errorf("loads = %d, value = %q", c.loads, c.value)
	}
}
