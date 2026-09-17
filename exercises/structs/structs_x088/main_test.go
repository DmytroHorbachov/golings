// structs_x088: Конфликт полей при встраивании в JSON
// Make the tests pass!
// I AM NOT DONE
//
// Record встраивает Meta и Audit, у обоих есть поле ID с тегом "id".
// encoding/json молча пропускает конфликтующие поля.
// Тренирует: правила продвижения полей в encoding/json.
// Сложность: hard
package main_test

import (
	"encoding/json"
	"testing"
)

type Meta struct {
	ID int `json:"id"`
}

type Audit struct {
	ID int `json:"id"`
}

type Record struct {
	Meta
	Audit
}

func TestRecordJSON(t *testing.T) {
	b, _ := json.Marshal(Record{Meta{1}, Audit{2}})
	if string(b) != `{"meta_id":1,"audit_id":2}` {
		t.Errorf("json = %s", b)
	}
}
