// maps_x087: Числа из JSON
// Make the tests pass!
// I AM NOT DONE
//
// portFrom достаёт порт из JSON-конфигурации, разобранной в map[string]interface{}.
// Утверждение типа к int всегда неудачно.
// Тренирует: encoding/json декодирует числа в interface{} как float64.
// Сложность: hard
package main_test

import (
	"encoding/json"
	"testing"
)

func portFrom(data []byte) (int, bool) {
	var cfg map[string]interface{}
	if err := json.Unmarshal(data, &cfg); err != nil {
		return 0, false
	}
	p, ok := cfg["port"].(int)
	return p, ok
}

func TestPortFrom(t *testing.T) {
	if p, ok := portFrom([]byte(`{"port": 8080}`)); !ok || p != 8080 {
		t.Errorf("portFrom = %d, %v", p, ok)
	}
	if _, ok := portFrom([]byte(`{"port": "80"}`)); ok {
		t.Errorf("string port should be rejected")
	}
}
