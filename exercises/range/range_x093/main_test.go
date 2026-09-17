// range_x093: Переименование переменной ключа
// Make the tests pass!
// I AM NOT DONE
//
// lowerKeys должна вернуть map с ключами в нижнем регистре.
// Изменение переменной k внутри цикла не меняет ключи map.
// Тренирует: переменная ключа в range — копия, не ссылка на ключ.
// Сложность: hard
package main_test

import (
	"reflect"
	"strings"
	"testing"
)

func lowerKeys(m map[string]int) map[string]int {
	for k := range m {
		k = strings.ToLower(k)
		_ = k
	}
	return m
}

func TestLowerKeys(t *testing.T) {
	got := lowerKeys(map[string]int{"Go": 1, "RUST": 2})
	if !reflect.DeepEqual(got, map[string]int{"go": 1, "rust": 2}) {
		t.Errorf("lowerKeys = %v", got)
	}
}
