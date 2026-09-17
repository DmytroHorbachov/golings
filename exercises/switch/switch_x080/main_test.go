// switch_x080: Пустой case
// Make the tests pass!
// I AM NOT DONE
//
// isWeekend должна вернуть true для субботы и воскресенья.
// Для субботы сейчас возвращается false.
// Тренирует: в Go нет неявного «проваливания» в следующий case.
// Сложность: hard
package main_test

import "testing"

func isWeekend(day string) bool {
	switch day {
	case "sat":
	case "sun":
		return true
	}
	return false
}

func TestIsWeekend(t *testing.T) {
	if !isWeekend("sat") || !isWeekend("sun") {
		t.Errorf("sat and sun are weekend days")
	}
	if isWeekend("mon") {
		t.Errorf("mon is not a weekend day")
	}
}
