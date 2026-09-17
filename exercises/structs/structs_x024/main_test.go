// structs_x024: Экспортируемые поля для JSON
// Make the tests pass!
// I AM NOT DONE
//
// Поле email не попадает в JSON, потому что оно неэкспортируемое.
// Тренирует: encoding/json видит только экспортируемые поля.
// Сложность: easy
package main_test

import (
	"encoding/json"
	"testing"
)

type Contact struct {
	Name  string `json:"name"`
	email string `json:"email"`
}

func TestContactJSON(t *testing.T) {
	var c Contact
	_ = json.Unmarshal([]byte(`{"name":"ann","email":"a@x"}`), &c)
	b, _ := json.Marshal(c)
	if string(b) != `{"name":"ann","email":"a@x"}` {
		t.Errorf("json = %s", b)
	}
}
