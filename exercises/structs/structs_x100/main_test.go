// structs_x100: Несериализуемое поле
// Make the tests pass!
// I AM NOT DONE
//
// Job содержит канал для отмены; json.Marshal падает с ошибкой из-за него.
// Тренирует: тег json:"-" исключает поле из сериализации.
// Сложность: hard
package main_test

import (
	"encoding/json"
	"testing"
)

type Job struct {
	ID     int           `json:"id"`
	Cancel chan struct{} `json:"cancel"`
}

func TestJobJSON(t *testing.T) {
	b, err := json.Marshal(Job{ID: 7, Cancel: make(chan struct{})})
	if err != nil || string(b) != `{"id":7}` {
		t.Errorf("json = %s, %v", b, err)
	}
}
