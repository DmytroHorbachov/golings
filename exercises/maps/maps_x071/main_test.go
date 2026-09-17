// maps_x071: Ключи разных типов
// Make the tests pass!
// I AM NOT DONE
//
// Реестр хранит обработчики в map[interface{}]string. Регистрация идёт по int,
// а поиск — по int64 из внешнего источника, и обработчик не находится.
// Тренирует: ключи interface{} равны, только если совпадают и тип, и значение.
// Сложность: hard
package main_test

import "testing"

type Registry map[interface{}]string

func (r Registry) Register(code int, name string) {
	r[code] = name
}

func (r Registry) Lookup(code int64) string {
	return r[code]
}

func TestRegistry(t *testing.T) {
	r := Registry{}
	r.Register(200, "ok")
	r.Register(404, "not found")
	var fromWire int64 = 404
	if got := r.Lookup(fromWire); got != "not found" {
		t.Errorf("Lookup(404) = %q, want not found", got)
	}
}
