// anonymous_functions52
// Make the tests pass!

// I AM NOT DONE
//
// formatter возвращает литерал форматирования чисел для локали:
// "ru" — "1 234,50", "en" — "1,234.50".
// Тренирует: литералы, настроенные параметрами фабрики.
// Сложность: medium
package main_test

import (
	"fmt"
	"strings"
	"testing"
)

func formatter(locale string) func(float64) string {
	group, dec := ",", "."
	if locale == "ru" {
		group, dec = ",", " "
	}
	return func(v float64) string {
		s := fmt.Sprintf("%.2f", v)
		intPart, frac := s[:len(s)-3], s[len(s)-2:]
		var parts []string
		for len(intPart) > 3 {
			parts = append([]string{intPart[len(intPart)-3:]}, parts...)
			intPart = intPart[:len(intPart)-3]
		}
		parts = append([]string{intPart}, parts...)
		return strings.Join(parts, dec) + group + frac
	}
}

func TestFormatter(t *testing.T) {
	if got := formatter("ru")(1234.5); got != "1 234,50" {
		t.Errorf("ru = %q", got)
	}
	if got := formatter("en")(1234567.891); got != "1,234,567.89" {
		t.Errorf("en = %q", got)
	}
}
