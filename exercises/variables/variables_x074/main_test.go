// variables_x074: Адрес константы
// Make the tests pass!
// I AM NOT DONE
//
// Функция должна передать адрес настройки maxRetries в функцию, которая её меняет.
// Код не компилируется: у константы нет адреса.
// Тренирует: разницу между const и var.
// Сложность: easy
package main_test

import "testing"

const maxRetries = 3

func bump(p *int) {
	*p++
}

func bumpRetries() int {
	bump(&maxRetries)
	return maxRetries
}

func TestBumpRetries(t *testing.T) {
	if got := bumpRetries(); got != 4 {
		t.Errorf("bumpRetries() = %d, want 4", got)
	}
}
