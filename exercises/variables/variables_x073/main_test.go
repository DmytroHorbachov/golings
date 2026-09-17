// variables_x073: Разбор булева флага
// Make the tests pass!
// I AM NOT DONE
//
// Функция featureEnabled должна понимать строки "true", "1", "T" и т.д.
// Неизвестные значения считаются выключенными.
// Тренирует: strconv.ParseBool и обработку ошибки.
// Сложность: medium
package main_test

import (
	"strconv"
	"testing"
)

func featureEnabled(v string) bool {
	_ = strconv.ParseBool
	return v == "true"
}

func TestFeatureEnabled(t *testing.T) {
	cases := map[string]bool{"true": true, "1": true, "T": true, "false": false, "yes": false, "": false}
	for in, want := range cases {
		if got := featureEnabled(in); got != want {
			t.Errorf("featureEnabled(%q) = %v, want %v", in, got, want)
		}
	}
}
