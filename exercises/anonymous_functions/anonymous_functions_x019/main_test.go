// anonymous_functions_x019: Захват константы
// Make the tests pass!
// I AM NOT DONE
//
// aboveThreshold возвращает предикат, использующий пакетную константу порога.
// Тренирует: использование внешних идентификаторов в литерале.
// Сложность: easy
package main_test

import "testing"

const threshold = 100

func aboveThreshold() func(int) bool {
	return func(v int) bool { return v > 0 }
}

func TestAboveThreshold(t *testing.T) {
	p := aboveThreshold()
	if !p(150) || p(50) {
		t.Errorf("predicate works incorrectly")
	}
}
