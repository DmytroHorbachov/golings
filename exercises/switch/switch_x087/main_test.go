// switch_x087: Неиспользуемая переменная type switch
// Make the tests pass!
// I AM NOT DONE
//
// category определяет категорию значения по типу. Код не компилируется:
// переменная type switch объявлена, но ни в одной ветке не используется.
// Тренирует: форма switch v := x.(type) требует использовать v.
// Сложность: hard
package main_test

import "testing"

func category(x interface{}) string {
	switch v := x.(type) {
	case int, float64:
		return "number"
	case string:
		return "text"
	}
	return "other"
}

func TestCategory(t *testing.T) {
	cases := []struct {
		in   interface{}
		want string
	}{{1, "number"}, {2.5, "number"}, {"a", "text"}, {true, "other"}}
	for _, c := range cases {
		if got := category(c.in); got != c.want {
			t.Errorf("category(%v) = %s, want %s", c.in, got, c.want)
		}
	}
}
