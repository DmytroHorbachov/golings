// algorithms_x144: Integer to English Words (число словами)
// Make the tests pass!
// I AM NOT DONE
//
// Паттерн: разбор по группам разрядов. Запишите неотрицательное число
// английскими словами, как принято: группы по три цифры с суффиксами
// Thousand, Million, Billion.
// Сложность: hard. Ожидаемая асимптотика: O(log n) по времени, O(1) по памяти
package main_test

import (
	"strings"
	"testing"
)

func numberToWords(n int) string {
	return ""
}

func TestNumberToWords(t *testing.T) {
	_ = strings.Join
	cases := map[int]string{
		0:          "Zero",
		123:        "One Hundred Twenty Three",
		12345:      "Twelve Thousand Three Hundred Forty Five",
		1234567:    "One Million Two Hundred Thirty Four Thousand Five Hundred Sixty Seven",
		1000000:    "One Million",
		20:         "Twenty",
		1000010:    "One Million Ten",
		2147483647: "Two Billion One Hundred Forty Seven Million Four Hundred Eighty Three Thousand Six Hundred Forty Seven",
	}
	for in, want := range cases {
		if got := numberToWords(in); got != want {
			t.Errorf("numberToWords(%d) = %q, want %q", in, got, want)
		}
	}
}
