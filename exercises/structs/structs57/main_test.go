// structs57
// Make the tests pass!

// I AM NOT DONE
//
// decodeAll разбирает JSON-записи в одну и ту же переменную.
// Поля, отсутствующие в очередной записи, сохраняют значения из предыдущей.
// Тренирует: json.Unmarshal не обнуляет поля, которых нет во входных данных.
// Сложность: hard
package main_test

import (
	"encoding/json"
	"testing"
)

type Item struct {
	Name  string `json:"name"`
	Color string `json:"color"`
}

func decodeAll(lines []string) []Item {
	var out []Item
	var it Item
	for _, l := range lines {
		if err := json.Unmarshal([]byte(l), &it); err == nil {
			out = append(out, it)
		}
	}
	return out
}

func TestDecodeAll(t *testing.T) {
	got := decodeAll([]string{`{"name":"car","color":"red"}`, `{"name":"box"}`})
	if len(got) != 2 || got[1].Color != "" {
		t.Errorf("decodeAll = %+v", got)
	}
}
