// primitive_types96
// Make the tests pass!

// I AM NOT DONE
//
// parseLevel разбирает уровень громкости в int8 и должна отклонять значения
// вне диапазона. Сейчас 300 молча превращается в 44.
// Тренирует: параметр bitSize в strconv.ParseInt.
// Сложность: hard
package main_test

import (
	"strconv"
	"testing"
)

func parseLevel(s string) (int8, error) {
	n, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0, err
	}
	return int8(n), nil
}

func TestParseLevel(t *testing.T) {
	if v, err := parseLevel("-128"); err != nil || v != -128 {
		t.Errorf("parseLevel(-128) = %d, %v", v, err)
	}
	for _, bad := range []string{"300", "128", "-129"} {
		if _, err := parseLevel(bad); err == nil {
			t.Errorf("parseLevel(%s) should fail", bad)
		}
	}
}
