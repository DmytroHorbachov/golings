// switch_x055: Вложенный switch
// Make the tests pass!
// I AM NOT DONE
//
// vat возвращает ставку НДС: для категории "food" — 10, но для подкатегории
// "delicacy" — 20; для "books" — 0; для остальных — 20.
// Тренирует: switch внутри ветки switch.
// Сложность: medium
package main_test

import "testing"

func vat(category, sub string) int {
	switch category {
	case "food":
		return 10
	case "books":
		return 10
	}
	return 20
}

func TestVAT(t *testing.T) {
	cases := []struct {
		cat, sub string
		want     int
	}{{"food", "bread", 10}, {"food", "delicacy", 20}, {"books", "", 0}, {"toys", "", 20}}
	for _, c := range cases {
		if got := vat(c.cat, c.sub); got != c.want {
			t.Errorf("vat(%s, %s) = %d, want %d", c.cat, c.sub, got, c.want)
		}
	}
}
