// maps_x095: Общая map пакета
// Make the tests pass!
// I AM NOT DONE
//
// defaultHeaders возвращает заголовки по умолчанию; вызывающий дополняет их.
// Изменения одного вызова видны во всех последующих.
// Тренирует: возврат map уровня пакета отдаёт общую изменяемую структуру.
// Сложность: hard
package main_test

import "testing"

var defaults = map[string]string{"Accept": "application/json"}

func defaultHeaders() map[string]string {
	return defaults
}

func TestDefaultHeaders(t *testing.T) {
	h1 := defaultHeaders()
	h1["Authorization"] = "secret"
	h2 := defaultHeaders()
	if _, leaked := h2["Authorization"]; leaked || len(h2) != 1 {
		t.Errorf("headers leaked between calls: %v", h2)
	}
}
