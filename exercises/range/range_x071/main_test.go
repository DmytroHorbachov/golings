// range_x071: i++ внутри range
// Make the tests pass!
// I AM NOT DONE
//
// parsePairs читает аргументы вида ["-k", "v", "-x", "y"]: после флага
// значение нужно пропустить. Увеличение i внутри range не работает.
// Тренирует: переменная индекса range перезаписывается на каждой итерации.
// Сложность: hard
package main_test

import (
	"reflect"
	"strings"
	"testing"
)

func parsePairs(args []string) map[string]string {
	m := map[string]string{}
	for i, a := range args {
		if strings.HasPrefix(a, "-") && i+1 < len(args) {
			m[a[1:]] = args[i+1]
			i++
		} else {
			m[a] = ""
		}
	}
	return m
}

func TestParsePairs(t *testing.T) {
	got := parsePairs([]string{"-k", "v", "-x", "y", "file"})
	if !reflect.DeepEqual(got, map[string]string{"k": "v", "x": "y", "file": ""}) {
		t.Errorf("parsePairs = %v", got)
	}
}
