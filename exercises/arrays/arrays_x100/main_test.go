// arrays_x100: Несравнимое внутри interface{}
// Make the tests pass!
// I AM NOT DONE
//
// sameArgs сравнивает два набора аргументов [2]interface{}.
// Если среди аргументов есть срезы, сравнение == паникует.
// Тренирует: массив интерфейсов сравним на этапе компиляции, но может паниковать.
// Сложность: hard
package main_test

import (
	"reflect"
	"testing"
)

func sameArgs(a, b [2]interface{}) bool {
	return a == b
}

func TestSameArgs(t *testing.T) {
	_ = reflect.DeepEqual
	if !sameArgs([2]interface{}{1, "x"}, [2]interface{}{1, "x"}) {
		t.Errorf("scalars should match")
	}
	if !sameArgs([2]interface{}{[]int{1}, 2}, [2]interface{}{[]int{1}, 2}) {
		t.Errorf("equal slices should match")
	}
}
