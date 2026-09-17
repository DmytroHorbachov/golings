// switch53
// Make the tests pass!

// I AM NOT DONE
//
// handle обрабатывает команду. Команда exit должна завершать работу,
// даже если она совпадает с предыдущей командой.
// Тренирует: case могут быть переменными, и тогда при совпадении выигрывает первый.
// Сложность: hard
package main_test

import "testing"

func handle(cmd, prev string) string {
	switch cmd {
	case prev:
		return "repeat"
	case "exit":
		return "bye"
	}
	return "run " + cmd
}

func TestHandle(t *testing.T) {
	cases := []struct{ cmd, prev, want string }{
		{"ls", "ls", "repeat"}, {"ls", "cd", "run ls"}, {"exit", "ls", "bye"}, {"exit", "exit", "bye"},
	}
	for _, c := range cases {
		if got := handle(c.cmd, c.prev); got != c.want {
			t.Errorf("handle(%s, %s) = %s, want %s", c.cmd, c.prev, got, c.want)
		}
	}
}
