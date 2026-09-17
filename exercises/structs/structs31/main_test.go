// structs31
// Make the tests pass!

// I AM NOT DONE
//
// Money сериализуется строкой "12.34 RUB" через MarshalJSON. Но при
// сериализации значения (не указателя) метод не вызывается.
// Тренирует: методы указателя не входят в набор методов значения.
// Сложность: hard
package main_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Money struct {
	Cents    int
	Currency string
}

func (m *Money) MarshalJSON() ([]byte, error) {
	return json.Marshal(fmt.Sprintf("%d.%02d %s", m.Cents/100, m.Cents%100, m.Currency))
}

type Invoice struct {
	Total Money `json:"total"`
}

func TestInvoiceJSON(t *testing.T) {
	b, _ := json.Marshal(Invoice{Money{1234, "RUB"}})
	if string(b) != `{"total":"12.34 RUB"}` {
		t.Errorf("json = %s", b)
	}
}
