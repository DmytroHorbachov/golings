// functions41
// Make the tests pass!

// I AM NOT DONE
//
// handler возвращает функцию-обработчик; очистка должна выполняться
// после обработки, при каждом вызове обработчика.
// Тренирует: defer относится к той функции, в теле которой он записан.
// Сложность: hard
package main_test

import (
	"reflect"
	"testing"
)

func handler(log *[]string) func() {
	*log = append(*log, "setup")
	defer func() { *log = append(*log, "teardown") }()
	return func() {
		*log = append(*log, "handle")
	}
}

func TestHandler(t *testing.T) {
	var log []string
	h := handler(&log)
	h()
	h()
	want := []string{"setup", "handle", "teardown", "handle", "teardown"}
	if !reflect.DeepEqual(log, want) {
		t.Errorf("log = %v, want %v", log, want)
	}
}
