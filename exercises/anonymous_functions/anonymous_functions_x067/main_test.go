// anonymous_functions_x067: Обход с остановкой по ошибке
// Make the tests pass!
// I AM NOT DONE
//
// walkFiles вызывает литерал для каждого файла; если литерал вернул ошибку,
// обход прекращается и ошибка возвращается.
// Тренирует: колбэки, возвращающие error.
// Сложность: medium
package main_test

import (
	"errors"
	"reflect"
	"strings"
	"testing"
)

func walkFiles(files []string, fn func(string) error) error {
	for _, f := range files {
		fn(f)
	}
	return nil
}

func TestWalkFiles(t *testing.T) {
	var seen []string
	err := walkFiles([]string{"a.go", "b.tmp", "c.go"}, func(f string) error {
		if strings.HasSuffix(f, ".tmp") {
			return errors.New("temp file " + f)
		}
		seen = append(seen, f)
		return nil
	})
	if err == nil || !reflect.DeepEqual(seen, []string{"a.go"}) {
		t.Errorf("err=%v seen=%v", err, seen)
	}
}
