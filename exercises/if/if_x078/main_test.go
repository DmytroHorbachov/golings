// if_x078: Index и нулевая позиция
// Make the tests pass!
// I AM NOT DONE
//
// mentions должна вернуть true, если в тексте есть имя.
// Когда имя стоит в самом начале, функция возвращает false.
// Тренирует: strings.Index возвращает 0 для совпадения в начале и -1 при отсутствии.
// Сложность: hard
package main_test

import (
	"strings"
	"testing"
)

func mentions(text, name string) bool {
	if strings.Index(text, name) > 0 {
		return true
	}
	return false
}

func TestMentions(t *testing.T) {
	cases := []struct {
		text, name string
		want       bool
	}{{"ann is here", "ann", true}, {"hi ann", "ann", true}, {"bob", "ann", false}}
	for _, c := range cases {
		if got := mentions(c.text, c.name); got != c.want {
			t.Errorf("mentions(%q, %q) = %v, want %v", c.text, c.name, got, c.want)
		}
	}
}
