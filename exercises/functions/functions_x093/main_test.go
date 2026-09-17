// functions_x093: Цикл инициализации
// Make the tests pass!
// I AM NOT DONE
//
// Команда help перечисляет все команды из той же таблицы, в которой записана сама.
// Код не компилируется: «initialization cycle».
// Тренирует: порядок инициализации переменных пакета и функции init.
// Сложность: hard
package main_test

import (
	"sort"
	"strings"
	"testing"
)

var commands = map[string]func() string{
	"help": help,
	"ping": func() string { return "pong" },
}

func help() string {
	names := make([]string, 0, len(commands))
	for name := range commands {
		names = append(names, name)
	}
	sort.Strings(names)
	return strings.Join(names, ",")
}

func TestHelp(t *testing.T) {
	if got := commands["help"](); got != "help,ping" {
		t.Errorf("help() = %q, want %q", got, "help,ping")
	}
	if got := commands["ping"](); got != "pong" {
		t.Errorf("ping() = %q", got)
	}
}
