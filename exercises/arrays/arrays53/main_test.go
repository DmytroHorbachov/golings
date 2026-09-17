// arrays53
// Make the tests pass!

// I AM NOT DONE
//
// backup должна вернуть исходное содержимое массива, даже если после
// копирования оригинал изменили.
// Тренирует: присваивание массива создаёт полную копию.
// Сложность: easy
package main_test

import "testing"

func backup() ([3]int, [3]int) {
	data := [3]int{1, 2, 3}
	saved := data
	data[0] = 100
	return data, data
}

func TestBackup(t *testing.T) {
	data, saved := backup()
	if data[0] != 100 || saved[0] != 1 {
		t.Errorf("data = %v, saved = %v", data, saved)
	}
}
