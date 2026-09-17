// structs85
// Make the tests pass!

// I AM NOT DONE
//
// Поле Name должно сериализоваться в JSON под ключом "name".
// Тренирует: теги полей для encoding/json.
// Сложность: easy
package main_test

import (
	"encoding/json"
	"testing"
)

type Pet struct {
	Name string `json:"title"`
	Age  int    `json:"age"`
}

func TestPetJSON(t *testing.T) {
	b, _ := json.Marshal(Pet{"Rex", 3})
	if string(b) != `{"name":"Rex","age":3}` {
		t.Errorf("json = %s", b)
	}
}
