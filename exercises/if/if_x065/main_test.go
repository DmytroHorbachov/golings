// if_x065: Стоимость отправки
// Make the tests pass!
// I AM NOT DONE
//
// shippingCost: до 1 кг — 200, до 5 кг — 400, тяжелее — 400 + 100 за каждый
// полный килограмм сверх 5. Экспресс-доставка удваивает цену.
// Тренирует: сочетание ветвления и модификатора.
// Сложность: medium
package main_test

import "testing"

func shippingCost(kg int, express bool) int {
	var cost int
	if kg <= 1 {
		cost = 200
	} else if kg <= 5 {
		cost = 400
	} else {
		cost = 100 * (kg - 5)
	}
	if express {
		cost += 2
	}
	return cost
}

func TestShippingCost(t *testing.T) {
	cases := []struct {
		kg      int
		express bool
		want    int
	}{{1, false, 200}, {3, false, 400}, {8, false, 700}, {1, true, 400}, {8, true, 1400}}
	for _, c := range cases {
		if got := shippingCost(c.kg, c.express); got != c.want {
			t.Errorf("shippingCost(%d, %v) = %d, want %d", c.kg, c.express, got, c.want)
		}
	}
}
