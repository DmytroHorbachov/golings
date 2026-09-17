// structs33
// Make the tests pass!

// I AM NOT DONE
//
// decode разбирает JSON в структуру, но возвращает пустую структуру и ошибку.
// Тренирует: json.Unmarshal требует указатель на значение.
// Сложность: hard
package main_test

import (
	"encoding/json"
	"testing"
)

type Point struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func decode(data []byte) (Point, error) {
	var p Point
	err := json.Unmarshal(data, p)
	return p, err
}

func TestDecode(t *testing.T) {
	p, err := decode([]byte(`{"x":1,"y":2}`))
	if err != nil || p != (Point{1, 2}) {
		t.Errorf("decode = %v, %v", p, err)
	}
}
