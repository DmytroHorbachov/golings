// primitive_types_x030: Вывод bool
// Make the tests pass!
// I AM NOT DONE
//
// status должна вернуть "enabled=true" или "enabled=false".
// Тренирует: глагол %t для bool.
// Сложность: easy
package main_test

import (
	"fmt"
	"testing"
)

func status(on bool) string {
	return fmt.Sprintf("enabled=%d", on)
}

func TestStatus(t *testing.T) {
	if status(true) != "enabled=true" || status(false) != "enabled=false" {
		t.Errorf("status = %q, %q", status(true), status(false))
	}
}
