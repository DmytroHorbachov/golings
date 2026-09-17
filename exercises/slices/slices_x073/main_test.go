// slices_x073: DeepEqual и nil
// Make the tests pass!
// I AM NOT DONE
//
// sameTags сравнивает наборы тегов; nil и пустой срез должны считаться равными.
// Тренирует: reflect.DeepEqual различает nil и пустой срез.
// Сложность: hard
package main_test

import (
	"reflect"
	"testing"
)

func sameTags(a, b []string) bool {
	return reflect.DeepEqual(a, b)
}

func TestSameTags(t *testing.T) {
	_ = reflect.DeepEqual
	if !sameTags(nil, []string{}) {
		t.Errorf("nil and empty tags should be equal")
	}
	if !sameTags([]string{"go"}, []string{"go"}) || sameTags([]string{"go"}, nil) {
		t.Errorf("sameTags works incorrectly")
	}
}
