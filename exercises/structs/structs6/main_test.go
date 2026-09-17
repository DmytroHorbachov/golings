// structs6
// Make the tests pass!

// I AM NOT DONE
//
// Поле Nickname не должно попадать в JSON, если оно пустое.
// Тренирует: опцию тега omitempty.
// Сложность: easy
package main_test

import (
	"encoding/json"
	"testing"
)

type Profile struct {
	Name     string `json:"name"`
	Nickname string `json:"nickname"`
}

func TestProfileJSON(t *testing.T) {
	b, _ := json.Marshal(Profile{Name: "ann"})
	if string(b) != `{"name":"ann"}` {
		t.Errorf("json = %s", b)
	}
}
