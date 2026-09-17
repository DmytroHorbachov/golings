// maps_x069: Случайный порядок обхода
// Make the tests pass!
// I AM NOT DONE
//
// report форматирует остатки товаров строкой "apple=3,kiwi=1,pear=5".
// Результат каждый раз разный.
// Тренирует: порядок обхода map не определён и намеренно рандомизирован.
// Сложность: hard
package main_test

import (
	"sort"
	"strconv"
	"strings"
	"testing"
)

func report(stock map[string]int) string {
	var parts []string
	for k, v := range stock {
		parts = append(parts, k+"="+strconv.Itoa(v))
	}
	return strings.Join(parts, ",")
}

func TestReport(t *testing.T) {
	_ = sort.Strings
	stock := map[string]int{"pear": 5, "apple": 3, "kiwi": 1, "fig": 2, "lime": 4}
	for i := 0; i < 50; i++ {
		if got := report(stock); got != "apple=3,fig=2,kiwi=1,lime=4,pear=5" {
			t.Fatalf("report = %q", got)
		}
	}
}
