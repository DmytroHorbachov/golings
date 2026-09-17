// switch_x052: Зоны доставки
// Make the tests pass!
// I AM NOT DONE
//
// zone возвращает зону доставки по коду страны: "domestic" для RU,
// "near" для BY и KZ, "far" для остальных; пустой код — ошибка.
// Тренирует: switch с проверкой ошибочного значения.
// Сложность: medium
package main_test

import (
	"errors"
	"testing"
)

func zone(country string) (string, error) {
	switch country {
	case "RU", "BY":
		return "domestic", nil
	case "KZ":
		return "near", nil
	}
	return "far", nil
}

func TestZone(t *testing.T) {
	_ = errors.New
	cases := map[string]string{"RU": "domestic", "BY": "near", "KZ": "near", "US": "far"}
	for in, want := range cases {
		if got, err := zone(in); err != nil || got != want {
			t.Errorf("zone(%s) = %s, %v; want %s", in, got, err, want)
		}
	}
	if _, err := zone(""); err == nil {
		t.Errorf("zone(\"\") should fail")
	}
}
