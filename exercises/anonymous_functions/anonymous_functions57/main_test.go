// anonymous_functions57
// Make the tests pass!

// I AM NOT DONE
//
// save должна откатывать изменения, если шаг завершился ошибкой.
// Отложенный литерал проверяет err, но ошибка записана в затенённую переменную.
// Тренирует: литерал видит внешнюю переменную, а не одноимённую внутреннюю.
// Сложность: hard
package main_test

import (
	"errors"
	"testing"
)

func save(fail bool, log *[]string) error {
	var err error
	defer func() {
		if err != nil {
			*log = append(*log, "rollback")
		}
	}()
	if fail {
		err := errors.New("disk full")
		return err
	}
	*log = append(*log, "commit")
	return nil
}

func TestSave(t *testing.T) {
	var log []string
	if save(true, &log) == nil || len(log) != 1 || log[0] != "rollback" {
		t.Errorf("failed save log = %v", log)
	}
}
