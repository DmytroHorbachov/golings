// functions_x067: Стек очистки
// Make the tests pass!
// I AM NOT DONE
//
// Cleanup накапливает функции освобождения ресурсов, а Run вызывает их
// в обратном порядке регистрации (как defer).
// Тренирует: хранение функций в срезе и обратный обход.
// Сложность: medium
package main_test

import (
	"reflect"
	"testing"
)

type Cleanup struct {
	fns []func()
}

func (c *Cleanup) Add(f func()) {
	c.fns = []func(){f}
}

func (c *Cleanup) Run() {
	for _, f := range c.fns {
		f()
	}
}

func TestCleanup(t *testing.T) {
	var log []string
	var c Cleanup
	c.Add(func() { log = append(log, "db") })
	c.Add(func() { log = append(log, "cache") })
	c.Add(func() { log = append(log, "file") })
	c.Run()
	c.Run()
	if want := []string{"file", "cache", "db"}; !reflect.DeepEqual(log, want) {
		t.Errorf("log = %v, want %v", log, want)
	}
}
