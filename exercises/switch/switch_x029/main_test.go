// switch_x029: Перевод температуры
// Make the tests pass!
// I AM NOT DONE
//
// convert переводит температуру из Цельсия в указанную шкалу ("F" или "K").
// Тренирует: switch с вычислениями в ветках.
// Сложность: easy
package main_test

import "testing"

func convert(c float64, scale string) float64 {
	switch scale {
	case "F":
		return c*9/5 + 32
	case "K":
		return c - 273.15
	}
	return c
}

func TestConvert(t *testing.T) {
	cases := []struct {
		c     float64
		scale string
		want  float64
	}{{100, "F", 212}, {0, "K", 273.15}, {25, "C", 25}}
	for _, cs := range cases {
		if got := convert(cs.c, cs.scale); got != cs.want {
			t.Errorf("convert(%v, %s) = %v, want %v", cs.c, cs.scale, got, cs.want)
		}
	}
}
