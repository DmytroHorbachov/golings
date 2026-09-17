// slices71
// Make the tests pass!

// I AM NOT DONE
//
// isEmpty должна вернуть true и для nil, и для пустого среза.
// Тренирует: len(nil) == 0.
// Сложность: easy
package main_test

import "testing"

func isEmpty(s []int) bool {
	return s == nil
}

func TestIsEmpty(t *testing.T) {
	if !isEmpty(nil) || !isEmpty([]int{}) || isEmpty([]int{1}) {
		t.Errorf("isEmpty works incorrectly")
	}
}
