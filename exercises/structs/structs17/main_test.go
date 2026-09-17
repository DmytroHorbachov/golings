// structs17
// Make the tests pass!

// I AM NOT DONE
//
// parseStrict должна отклонять JSON с неизвестными полями (например, опечатками).
// json.Unmarshal молча их игнорирует.
// Тренирует: json.Decoder.DisallowUnknownFields.
// Сложность: hard
package main_test

import (
	"bytes"
	"encoding/json"
	"testing"
)

type Limits struct {
	MaxUsers int `json:"max_users"`
}

func parseStrict(data []byte) (Limits, error) {
	var l Limits
	err := json.Unmarshal(data, &l)
	_ = bytes.NewReader
	return l, err
}

func TestParseStrict(t *testing.T) {
	if l, err := parseStrict([]byte(`{"max_users":5}`)); err != nil || l.MaxUsers != 5 {
		t.Errorf("parseStrict = %v, %v", l, err)
	}
	if _, err := parseStrict([]byte(`{"max_user":5}`)); err == nil {
		t.Errorf("typo in field name should be rejected")
	}
}
