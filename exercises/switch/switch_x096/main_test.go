// switch_x096: Несравнимый тег
// Make the tests pass!
// I AM NOT DONE
//
// tariff выбирает тариф по ключу-структуре. Код не компилируется:
// структура содержит срез и поэтому несравнима.
// Тренирует: в switch можно использовать только сравнимые типы.
// Сложность: hard
package main_test

import "testing"

type key struct {
	Country string
	Tags    []string
}

func tariff(k key) int {
	switch k {
	case key{"RU", "vip"}:
		return 50
	case key{"RU", ""}:
		return 100
	}
	return 200
}

func TestTariff(t *testing.T) {
	if got := tariff(key{"RU", "vip"}); got != 50 {
		t.Errorf("tariff(RU vip) = %d", got)
	}
	if got := tariff(key{"RU", ""}); got != 100 {
		t.Errorf("tariff(RU) = %d", got)
	}
	if got := tariff(key{"US", ""}); got != 200 {
		t.Errorf("tariff(US) = %d", got)
	}
}
