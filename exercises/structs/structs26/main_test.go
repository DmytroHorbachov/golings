// structs26
// Make the tests pass!

// I AM NOT DONE
//
// Встроенная структура Address даёт поле City напрямую у Company.
// Тренирует: встраивание (embedding) и продвижение полей.
// Сложность: easy
package main_test

import "testing"

type Address struct{ City string }

type Company struct {
	Name string
	Addr Address
}

func TestCompanyCity(t *testing.T) {
	c := Company{Name: "Acme"}
	c.Address.City = "Kazan"
	if c.City != "Kazan" {
		t.Errorf("City = %q", c.City)
	}
}
