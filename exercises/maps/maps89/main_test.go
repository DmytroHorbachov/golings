// maps89
// Make the tests pass!

// I AM NOT DONE
//
// mergeInto копирует все пары из src в dst.
// Тренирует: range по map и запись в другую map.
// Сложность: easy
package main_test

import (
	"reflect"
	"testing"
)

func mergeInto(dst, src map[string]int) {
	for k, v := range src {
		src[k] = v
	}
}

func TestMergeInto(t *testing.T) {
	dst := map[string]int{"a": 1}
	mergeInto(dst, map[string]int{"b": 2, "a": 3})
	if !reflect.DeepEqual(dst, map[string]int{"a": 3, "b": 2}) {
		t.Errorf("dst = %v", dst)
	}
}
