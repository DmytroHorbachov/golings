// variables_x006: Размеры через сдвиг
// Make the tests pass!
// I AM NOT DONE
//
// Константы KB, MB, GB должны быть степенями 1024.
// Тренирует: iota вместе с побитовым сдвигом.
// Сложность: easy
package main_test

import "testing"

const (
	_  = iota
	KB = 1 << iota
	MB
	GB
)

func TestSizes(t *testing.T) {
	if KB != 1024 || MB != 1024*1024 || GB != 1024*1024*1024 {
		t.Errorf("KB=%d MB=%d GB=%d, want powers of 1024", KB, MB, GB)
	}
}
