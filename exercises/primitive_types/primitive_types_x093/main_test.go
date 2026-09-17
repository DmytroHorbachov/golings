// primitive_types_x093: Split пустой строки
// Make the tests pass!
// I AM NOT DONE
//
// tags разбирает строку тегов через запятую. Для пустой строки должно
// получиться ноль тегов, а получается один пустой тег.
// Тренирует: strings.Split("", ",") возвращает [""], а не пустой срез.
// Сложность: hard
package main_test

import (
	"strings"
	"testing"
)

func tags(s string) []string {
	return strings.Split(s, ",")
}

func TestTags(t *testing.T) {
	if got := tags(""); len(got) != 0 {
		t.Errorf("tags(\"\") = %q, want empty", got)
	}
	if got := tags("go,web"); len(got) != 2 || got[0] != "go" || got[1] != "web" {
		t.Errorf("tags(go,web) = %q", got)
	}
	if got := tags("go,,web,"); len(got) != 2 {
		t.Errorf("tags(go,,web,) = %q, want 2 tags", got)
	}
}
