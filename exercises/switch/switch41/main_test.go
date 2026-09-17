// switch41
// Make the tests pass!

// I AM NOT DONE
//
// describe возвращает "read", "write", "read-write" или "none" для набора флагов.
// Для Read|Write возвращается "none": switch по значению не видит комбинацию.
// Тренирует: битовые флаги нельзя разбирать switch-ем по отдельным константам.
// Сложность: hard
package main_test

import "testing"

const (
	Read = 1 << iota
	Write
)

func describe(flags int) string {
	switch flags {
	case Read:
		return "read"
	case Write:
		return "write"
	}
	return "none"
}

func TestDescribe(t *testing.T) {
	cases := map[int]string{0: "none", Read: "read", Write: "write", Read | Write: "read-write"}
	for in, want := range cases {
		if got := describe(in); got != want {
			t.Errorf("describe(%b) = %s, want %s", in, got, want)
		}
	}
}
