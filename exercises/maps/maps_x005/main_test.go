// maps_x005: Размер map
// Make the tests pass!
// I AM NOT DONE
//
// userCount возвращает количество пользователей.
// Тренирует: len для map.
// Сложность: easy
package main_test

import "testing"

func userCount(m map[int]string) int {
	return len(m) - 1
}

func TestUserCount(t *testing.T) {
	if got := userCount(map[int]string{1: "a", 2: "b", 3: "c"}); got != 3 {
		t.Errorf("userCount = %d", got)
	}
}
