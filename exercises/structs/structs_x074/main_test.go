// structs_x074: omitempty и false
// Make the tests pass!
// I AM NOT DONE
//
// Настройка Notify по умолчанию true; явное false должно попадать в JSON,
// а незаданное значение — нет. С omitempty false теряется.
// Тренирует: omitempty пропускает нулевые значения, включая false.
// Сложность: hard
package main_test

import (
	"encoding/json"
	"testing"
)

type Settings struct {
	Notify bool `json:"notify,omitempty"`
}

func TestSettingsJSON(t *testing.T) {
	var unset Settings
	b, _ := json.Marshal(unset)
	if string(b) != `{}` {
		t.Errorf("unset = %s", b)
	}
	off := false
	var s Settings
	_ = json.Unmarshal([]byte(`{"notify":false}`), &s)
	b, _ = json.Marshal(s)
	if string(b) != `{"notify":false}` {
		t.Errorf("explicit false = %s", b)
	}
	_ = off
}
