// structs_x073: Неэкспортируемое поле и JSON
// Make the tests pass!
// I AM NOT DONE
//
// После Unmarshal поле token остаётся пустым, хотя ключ есть в JSON.
// Тренирует: encoding/json игнорирует неэкспортируемые поля даже с тегом.
// Сложность: hard
package main_test

import (
	"encoding/json"
	"testing"
)

type Session struct {
	User  string `json:"user"`
	token string `json:"token"`
}

func tokenOf(data []byte) string {
	var s Session
	_ = json.Unmarshal(data, &s)
	return s.Token
}

func TestTokenOf(t *testing.T) {
	if got := tokenOf([]byte(`{"user":"ann","token":"xyz"}`)); got != "xyz" {
		t.Errorf("tokenOf = %q", got)
	}
}
