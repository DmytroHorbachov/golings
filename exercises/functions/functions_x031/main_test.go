// functions_x031: Необязательный аргумент
// Make the tests pass!
// I AM NOT DONE
//
// greet(name, punct...) использует первый необязательный аргумент как знак препинания,
// а по умолчанию ставит ".".
// Тренирует: вариативный параметр как способ задать необязательный аргумент.
// Сложность: easy
package main_test

import "testing"

func greet(name string, punct ...string) string {
	p := "."
	if len(punct) > 1 {
		p = punct[0]
	}
	return "Hi, " + name + p
}

func TestGreet(t *testing.T) {
	if got := greet("Ann"); got != "Hi, Ann." {
		t.Errorf("greet(Ann) = %q", got)
	}
	if got := greet("Bob", "!"); got != "Hi, Bob!" {
		t.Errorf("greet(Bob, !) = %q", got)
	}
}
