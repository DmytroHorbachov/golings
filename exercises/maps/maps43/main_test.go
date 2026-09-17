// maps43
// Make the tests pass!

// I AM NOT DONE
//
// lookupCountry ищет страну по коду без учёта регистра; ключи хранятся в верхнем регистре.
// Тренирует: нормализацию ключа перед поиском.
// Сложность: easy
package main_test

import (
	"strings"
	"testing"
)

var countries = map[string]string{"RU": "Russia", "DE": "Germany"}

func lookupCountry(code string) string {
	return countries[strings.ToLower(code)]
}

func TestLookupCountry(t *testing.T) {
	if lookupCountry("ru") != "Russia" || lookupCountry("De") != "Germany" {
		t.Errorf("lookupCountry works incorrectly")
	}
}
