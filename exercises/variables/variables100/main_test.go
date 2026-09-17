// variables100
// Make the tests pass!

// I AM NOT DONE
//
// Номера кабинетов записаны в таблице с ведущими нулями для красоты.
// Кабинет "010" должен иметь номер 10, но получается 8.
// Тренирует: целочисленный литерал с ведущим нулём — восьмеричный.
// Сложность: hard
package main_test

import "testing"

func roomNumbers() []int {
	return []int{007, 010, 012}
}

func TestRoomNumbers(t *testing.T) {
	want := []int{7, 10, 12}
	got := roomNumbers()
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("room %d = %d, want %d", i, got[i], want[i])
		}
	}
}
