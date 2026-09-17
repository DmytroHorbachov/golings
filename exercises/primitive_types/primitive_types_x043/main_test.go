// primitive_types_x043: Температура с единицей
// Make the tests pass!
// I AM NOT DONE
//
// toCelsius разбирает строки вида "36.6C" или "98.6F" и возвращает градусы Цельсия.
// Тренирует: работу с последним байтом строки и ParseFloat.
// Сложность: medium
package main_test

import (
	"errors"
	"math"
	"strconv"
	"testing"
)

func toCelsius(s string) (float64, error) {
	if len(s) < 2 {
		return 0, errors.New("too short")
	}
	v, err := strconv.ParseFloat(s[:len(s)-1], 64)
	if err != nil {
		return 0, err
	}
	switch s[len(s)-1] {
	case 'C':
		return v, nil
	case 'F':
		return v*9/5 + 32, nil
	}
	return v, nil
}

func TestToCelsius(t *testing.T) {
	if v, err := toCelsius("36.6C"); err != nil || v != 36.6 {
		t.Errorf("toCelsius(36.6C) = %v, %v", v, err)
	}
	if v, err := toCelsius("212F"); err != nil || math.Abs(v-100) > 1e-9 {
		t.Errorf("toCelsius(212F) = %v, %v", v, err)
	}
	if _, err := toCelsius("300K"); err == nil {
		t.Errorf("toCelsius(300K) should fail")
	}
}
