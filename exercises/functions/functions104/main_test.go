// functions104
// Make the tests pass!

// I AM NOT DONE
//
// execute должна найти команду по имени и выполнить её,
// а для неизвестной команды вернуть "unknown command".
// Тренирует: map[string]func и проверку наличия ключа.
// Сложность: medium
package main_test

import "testing"

var commands = map[string]func() string{
	"ping":    func() string { return "pong" },
	"version": func() string { return "1.0" },
}

func execute(name string) string {
	cmd := commands[name]
	return cmd()
}

func TestExecute(t *testing.T) {
	cases := map[string]string{"ping": "pong", "version": "1.0", "rm": "unknown command"}
	for in, want := range cases {
		if got := execute(in); got != want {
			t.Errorf("execute(%q) = %q, want %q", in, got, want)
		}
	}
}
