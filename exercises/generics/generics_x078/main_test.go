// generics_x078: Ограничение с указателем
// Make the tests pass!
// I AM NOT DONE
//
// ParseAll создаёт значения типа T и вызывает у них метод Set с получателем-указателем.
// Код не компилируется: у T нет метода Set, он есть только у *T.
// Тренирует: паттерн ограничения «указатель на T с методом».
// Сложность: hard
package main_test

import (
	"reflect"
	"strings"
	"testing"
)

type Tag struct{ Name string }

func (t *Tag) Set(s string) { t.Name = strings.ToLower(s) }

type Setter interface{ Set(string) }

func ParseAll[T Setter](items []string) []T {
	out := make([]T, len(items))
	for i, s := range items {
		out[i].Set(s)
	}
	return out
}

func tags() []Tag { return ParseAll[Tag]([]string{"Go", "WEB"}) }

func TestTags(t *testing.T) {
	if got := tags(); !reflect.DeepEqual(got, []Tag{{"go"}, {"web"}}) {
		t.Errorf("tags = %v", got)
	}
}
