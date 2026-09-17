// anonymous_functions55
// Make the tests pass!

// I AM NOT DONE
//
// withDefaults возвращает функцию, дополняющую запрос значениями по умолчанию.
// Литерал меняет общий объект настроек, и последующие запросы получают чужие значения.
// Тренирует: захват указателя даёт доступ к общему изменяемому объекту.
// Сложность: hard
package main_test

import "testing"

type Req struct {
	Timeout int
	Path    string
}

func withDefaults(def *Req) func(path string) Req {
	return func(path string) Req {
		r := def
		r.Path = path
		return *r
	}
}

func TestWithDefaults(t *testing.T) {
	def := &Req{Timeout: 30}
	mk := withDefaults(def)
	a := mk("/a")
	b := mk("/b")
	if a.Path != "/a" || b.Path != "/b" || def.Path != "" {
		t.Errorf("a=%+v b=%+v def=%+v", a, b, *def)
	}
}
