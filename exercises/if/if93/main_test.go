// if93
// Make the tests pass!

// I AM NOT DONE
//
// stockStatus должна различать товар с нулевым остатком ("out of stock")
// и товар, которого нет в каталоге ("unknown").
// Тренирует: чтение отсутствующего ключа map возвращает нулевое значение.
// Сложность: hard
package main_test

import "testing"

var stock = map[string]int{"apple": 5, "pear": 0}

func stockStatus(item string) string {
	qty := stock[item]
	if qty == 0 {
		return "unknown"
	}
	if qty == 0 {
		return "out of stock"
	}
	return "available"
}

func TestStockStatus(t *testing.T) {
	cases := map[string]string{"apple": "available", "pear": "out of stock", "kiwi": "unknown"}
	for in, want := range cases {
		if got := stockStatus(in); got != want {
			t.Errorf("stockStatus(%s) = %q, want %q", in, got, want)
		}
	}
}
