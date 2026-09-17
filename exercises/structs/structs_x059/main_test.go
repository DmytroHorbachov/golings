// structs_x059: Форматирование через Stringer
// Make the tests pass!
// I AM NOT DONE
//
// Temperature.String выводит значение с одним знаком и единицей, "-" для отсутствия.
// Тренирует: String() и fmt с разными значениями.
// Сложность: medium
package main_test

import (
	"fmt"
	"testing"
)

type Temperature struct {
	Value float64
	Valid bool
}

func (t Temperature) String() string {
	return fmt.Sprintf("%v°C", t.Value)
}

func TestTemperatureString(t *testing.T) {
	got := fmt.Sprintf("%v %v", Temperature{21.456, true}, Temperature{})
	if got != "21.5°C -" {
		t.Errorf("got %q", got)
	}
}
