// anonymous_functions_x038: Хуки до и после
// Make the tests pass!
// I AM NOT DONE
//
// run выполняет действие, вызывая хуки before и after (если заданы),
// и хук after вызывается даже при ошибке действия.
// Тренирует: необязательные литералы и defer.
// Сложность: medium
package main_test

import (
	"errors"
	"reflect"
	"testing"
)

type Hooks struct {
	Before, After func()
}

func run(h Hooks, action func() error) error {
	h.Before()
	err := action()
	h.After()
	return err
}

func TestRun(t *testing.T) {
	var log []string
	h := Hooks{After: func() { log = append(log, "after") }}
	err := run(h, func() error { log = append(log, "act"); return errors.New("x") })
	if err == nil || !reflect.DeepEqual(log, []string{"act", "after"}) {
		t.Errorf("err=%v log=%v", err, log)
	}
}
