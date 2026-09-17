// switch_x013: Дубликат case
// Make the tests pass!
// I AM NOT DONE
//
// sizeLabel возвращает метку размера. Код не компилируется: два одинаковых case.
// Тренирует: константы в case должны быть уникальными.
// Сложность: easy
package main_test

import "testing"

func sizeLabel(size string) string {
	switch size {
	case "S":
		return "small"
	case "M":
		return "medium"
	case "S":
		return "large"
	}
	return "custom"
}

func TestSizeLabel(t *testing.T) {
	cases := map[string]string{"S": "small", "M": "medium", "L": "large", "XL": "custom"}
	for in, want := range cases {
		if got := sizeLabel(in); got != want {
			t.Errorf("sizeLabel(%s) = %s, want %s", in, got, want)
		}
	}
}
