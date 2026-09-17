// structs68
// Make the tests pass!

// I AM NOT DONE
//
// History хранит последние limit команд; Add добавляет, Last(n) возвращает n последних.
// Тренирует: структуру с ограниченным срезом.
// Сложность: medium
package main_test

import (
	"reflect"
	"testing"
)

type History struct {
	limit int
	cmds  []string
}

func (h *History) Add(cmd string) {
	h.cmds = append(h.cmds, cmd)
}

func (h *History) Last(n int) []string {
	return h.cmds[len(h.cmds)-n:]
}

func TestHistory(t *testing.T) {
	h := &History{limit: 3}
	for _, c := range []string{"ls", "cd", "pwd", "git"} {
		h.Add(c)
	}
	if !reflect.DeepEqual(h.Last(2), []string{"pwd", "git"}) || len(h.Last(10)) != 3 {
		t.Errorf("history = %v", h.cmds)
	}
}
