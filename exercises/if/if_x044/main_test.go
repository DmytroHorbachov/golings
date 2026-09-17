// if_x044: Вид треугольника
// Make the tests pass!
// I AM NOT DONE
//
// triangleKind: "equilateral" — все стороны равны, "isosceles" — две равны,
// "scalene" — все разные.
// Тренирует: порядок условий, когда одно является частным случаем другого.
// Сложность: medium
package main_test

import "testing"

func triangleKind(a, b, c int) string {
	if a == b {
		return "isosceles"
	} else if a == b && b == c {
		return "equilateral"
	}
	return "scalene"
}

func TestTriangleKind(t *testing.T) {
	cases := []struct {
		a, b, c int
		want    string
	}{{2, 2, 2, "equilateral"}, {2, 2, 3, "isosceles"}, {3, 2, 2, "isosceles"}, {2, 3, 2, "isosceles"}, {3, 4, 5, "scalene"}}
	for _, c := range cases {
		if got := triangleKind(c.a, c.b, c.c); got != c.want {
			t.Errorf("triangleKind(%d,%d,%d) = %s, want %s", c.a, c.b, c.c, got, c.want)
		}
	}
}
