// anonymous_functions66
// Make the tests pass!

// I AM NOT DONE
//
// transaction выполняет шаги; каждый шаг может зарегистрировать литерал отката.
// При ошибке откаты выполняются в обратном порядке.
// Тренирует: накопление литералов и обратный порядок выполнения.
// Сложность: medium
package main_test

import (
	"errors"
	"reflect"
	"testing"
)

type step func(onUndo func(func())) error

func transaction(steps ...step) error {
	var undos []func()
	register := func(u func()) { undos = append(undos, u) }
	for _, s := range steps {
		if err := s(register); err != nil {
			for _, u := range undos {
				u()
			}
			return err
		}
	}
	return nil
}

func TestTransaction(t *testing.T) {
	var log []string
	mk := func(name string, fail bool) step {
		return func(onUndo func(func())) error {
			if fail {
				return errors.New(name + " failed")
			}
			log = append(log, "do "+name)
			onUndo(func() { log = append(log, "undo "+name) })
			return nil
		}
	}
	err := transaction(mk("a", false), mk("b", false), mk("c", true))
	want := []string{"do a", "do b", "undo b", "undo a"}
	if err == nil || !reflect.DeepEqual(log, want) {
		t.Errorf("err=%v log=%v", err, log)
	}
}
