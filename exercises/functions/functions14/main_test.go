// functions14
// Make the tests pass!

// I AM NOT DONE
//
// missingKey должна достать ключ из ошибки NotFound в цепочке.
// Ошибка есть, но errors.As её не находит.
// Тренирует: тип цели в errors.As должен точно совпадать с типом ошибки в цепочке.
// Сложность: hard
package main_test

import (
	"errors"
	"fmt"
	"testing"
)

type NotFound struct{ Key string }

func (e NotFound) Error() string { return "not found: " + e.Key }

func find(key string) error {
	return fmt.Errorf("db: %w", NotFound{Key: key})
}

func missingKey(err error) string {
	var nf *NotFound
	if errors.As(err, &nf) {
		return nf.Key
	}
	return ""
}

func TestMissingKey(t *testing.T) {
	if got := missingKey(find("user:42")); got != "user:42" {
		t.Errorf("missingKey = %q, want user:42", got)
	}
	if got := missingKey(errors.New("other")); got != "" {
		t.Errorf("missingKey(other) = %q, want empty", got)
	}
}
