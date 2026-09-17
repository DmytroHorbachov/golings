// range52
// Make the tests pass!

// I AM NOT DONE
//
// normalizeAll приводит строки к нижнему регистру по указателю на срез.
// Код не компилируется: range по *[]string недопустим.
// Тренирует: range работает с указателем на массив, но не на срез.
// Сложность: hard
package main_test

import (
	"reflect"
	"strings"
	"testing"
)

func normalizeAll(p *[]string) {
	for i, s := range p {
		p[i] = strings.ToLower(s)
	}
}

func TestNormalizeAll(t *testing.T) {
	s := []string{"Go", "RUST"}
	normalizeAll(&s)
	if !reflect.DeepEqual(s, []string{"go", "rust"}) {
		t.Errorf("normalizeAll = %v", s)
	}
}
