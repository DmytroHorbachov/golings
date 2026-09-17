// concurrent_x075: Разблокировка без блокировки
// Make the tests pass!
// I AM NOT DONE
//
// update снимает блокировку в двух местах при раннем выходе, и программа
// аварийно завершается: «unlock of unlocked mutex».
// Тренирует: одна пара Lock/Unlock на путь выполнения.
// Сложность: hard
package main_test

import (
	"sync"
	"testing"
)

type Cache struct {
	mu   sync.Mutex
	data map[string]string
}

func (c *Cache) Update(k, v string) bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	if v == "" {
		c.mu.Unlock()
		return false
	}
	c.data[k] = v
	return true
}

func TestUpdate(t *testing.T) {
	c := &Cache{data: map[string]string{}}
	if !c.Update("a", "1") || c.Update("b", "") || !c.Update("c", "3") {
		t.Errorf("Update results are wrong")
	}
}
