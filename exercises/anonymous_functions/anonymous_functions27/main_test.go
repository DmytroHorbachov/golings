// anonymous_functions27
// Make the tests pass!

// I AM NOT DONE
//
// closeAll откладывает закрытие каждого ресурса литералом.
// Все отложенные литералы закрывают последний ресурс.
// Тренирует: отложенные литералы читают переменную в момент выполнения.
// Сложность: hard
package main_test

import (
	"reflect"
	"testing"
)

func closeAll(names []string) (closed []string) {
	var current string
	func() {
		for _, n := range names {
			current = n
			defer func() { closed = append(closed, current) }()
		}
	}()
	return closed
}

func TestCloseAll(t *testing.T) {
	if got := closeAll([]string{"db", "cache", "file"}); !reflect.DeepEqual(got, []string{"file", "cache", "db"}) {
		t.Errorf("closeAll = %v", got)
	}
}
