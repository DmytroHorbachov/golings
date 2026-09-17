// range66
// Make the tests pass!

// I AM NOT DONE
//
// everyKth возвращает каждый k-й элемент, начиная с k-го (позиции k, 2k, ... при счёте с 1).
// Тренирует: использование индекса range в условии.
// Сложность: medium
package main_test

import (
	"reflect"
	"testing"
)

func everyKth(s []string, k int) []string {
	var out []string
	for i, v := range s {
		if i%k == 0 {
			out = append(out, s[i/k])
		}
		_ = v
	}
	return out
}

func TestEveryKth(t *testing.T) {
	got := everyKth([]string{"a", "b", "c", "d", "e", "f", "g"}, 3)
	if !reflect.DeepEqual(got, []string{"c", "f"}) {
		t.Errorf("everyKth = %v", got)
	}
}
