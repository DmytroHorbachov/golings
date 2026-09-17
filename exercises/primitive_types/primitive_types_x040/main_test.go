// primitive_types_x040: Пары ключ=значение
// Make the tests pass!
// I AM NOT DONE
//
// parsePairs разбирает строку "a=1;b=20" в map[string]int, пропуская
// пустые и некорректные пары.
// Тренирует: strings.Split, strings.Cut и strconv.Atoi вместе.
// Сложность: medium
package main_test

import (
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func parsePairs(s string) map[string]int {
	out := map[string]int{}
	for _, part := range strings.Split(s, ";") {
		k, v, _ := strings.Cut(part, ":")
		n, _ := strconv.Atoi(v)
		out[k] = n
	}
	return out
}

func TestParsePairs(t *testing.T) {
	got := parsePairs("a=1;b=20;;bad;c=x;d=-4")
	want := map[string]int{"a": 1, "b": 20, "d": -4}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("parsePairs = %v, want %v", got, want)
	}
}
