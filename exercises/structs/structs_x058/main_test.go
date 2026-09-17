// structs_x058: Методы именованного типа
// Make the tests pass!
// I AM NOT DONE
//
// Celsius и Fahrenheit — именованные типы с методами перевода друг в друга.
// Тренирует: методы на типах, не являющихся структурами.
// Сложность: medium
package main_test

import "testing"

type Celsius float64
type Fahrenheit float64

func (c Celsius) ToF() Fahrenheit {
	return Fahrenheit(c*5/9 + 32)
}

func (f Fahrenheit) ToC() Celsius {
	return Celsius(f-32) * 9 / 5
}

func TestTemperature(t *testing.T) {
	if Celsius(100).ToF() != 212 || Fahrenheit(212).ToC() != 100 || Celsius(-40).ToF() != -40 {
		t.Errorf("conversions are wrong")
	}
}
