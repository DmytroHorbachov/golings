// structs_x082: Интерфейсное поле и ==
// Make the tests pass!
// I AM NOT DONE
//
// Две структуры Event с полем Payload interface{} сравниваются через ==.
// Если в Payload лежит срез, сравнение паникует.
// Тренирует: == для структур с интерфейсными полями может паниковать.
// Сложность: hard
package main_test

import (
	"reflect"
	"testing"
)

type Event struct {
	Kind    string
	Payload interface{}
}

func sameEvent(a, b Event) bool {
	return a == b
}

func TestSameEvent(t *testing.T) {
	_ = reflect.DeepEqual
	if !sameEvent(Event{"x", 1}, Event{"x", 1}) {
		t.Errorf("scalar payloads should match")
	}
	if !sameEvent(Event{"batch", []int{1, 2}}, Event{"batch", []int{1, 2}}) {
		t.Errorf("slice payloads should match")
	}
}
