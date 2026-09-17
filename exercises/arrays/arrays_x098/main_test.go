// arrays_x098: Снимок или ссылка
// Make the tests pass!
// I AM NOT DONE
//
// track запоминает состояние массива до изменений и сообщает, изменилось ли оно.
// Снимок хранится как указатель и поэтому изменяется вместе с оригиналом.
// Тренирует: присваивание массива копирует, а взятие адреса — нет.
// Сложность: hard
package main_test

import "testing"

func track(state *[3]int, mutate func(*[3]int)) bool {
	before := state
	mutate(state)
	return *before != *state
}

func TestTrack(t *testing.T) {
	s := [3]int{1, 2, 3}
	if !track(&s, func(p *[3]int) { p[0] = 9 }) {
		t.Errorf("change not detected")
	}
	if track(&s, func(p *[3]int) {}) {
		t.Errorf("false change detected")
	}
}
