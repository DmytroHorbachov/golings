// slices67
// Make the tests pass!

// I AM NOT DONE
//
// restore копирует сохранённое состояние в рабочий буфер и должна сообщать
// об ошибке, если буфер слишком мал. Сейчас данные молча обрезаются.
// Тренирует: copy не сообщает об обрезке — нужно проверять результат.
// Сложность: hard
package main_test

import (
	"errors"
	"testing"
)

func restore(dst, saved []int) error {
	copy(dst, saved)
	return nil
}

func TestRestore(t *testing.T) {
	_ = errors.New
	dst := make([]int, 3)
	if err := restore(dst, []int{1, 2, 3}); err != nil || dst[2] != 3 {
		t.Errorf("restore into big buffer: %v, %v", err, dst)
	}
	if err := restore(make([]int, 2), []int{1, 2, 3}); err == nil {
		t.Errorf("restore into small buffer should fail")
	}
}
