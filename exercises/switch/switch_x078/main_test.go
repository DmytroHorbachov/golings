// switch_x078: Вызов в каждом case
// Make the tests pass!
// I AM NOT DONE
//
// nextAction читает очередной токен и возвращает действие.
// Функция вызывает next() в нескольких ветках и теряет токены.
// Тренирует: выражения в case вычисляются последовательно до первого совпадения.
// Сложность: hard
package main_test

import (
	"reflect"
	"testing"
)

type lexer struct{ toks []string }

func (l *lexer) next() string {
	if len(l.toks) == 0 {
		return ""
	}
	t := l.toks[0]
	l.toks = l.toks[1:]
	return t
}

func nextAction(l *lexer) string {
	switch {
	case l.next() == "go":
		return "run"
	case l.next() == "stop":
		return "halt"
	}
	return "skip"
}

func TestNextAction(t *testing.T) {
	l := &lexer{toks: []string{"stop", "x", "go", "stop"}}
	var got []string
	for i := 0; i < 4; i++ {
		got = append(got, nextAction(l))
	}
	if want := []string{"halt", "skip", "run", "halt"}; !reflect.DeepEqual(got, want) {
		t.Errorf("actions = %v, want %v", got, want)
	}
}
