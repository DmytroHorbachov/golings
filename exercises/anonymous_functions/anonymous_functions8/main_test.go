// anonymous_functions8
// Make the tests pass!

// I AM NOT DONE
//
// defaultConfig должна быть значением Config, вычисленным литералом при инициализации.
// Код не компилируется: литерал не вызван.
// Тренирует: func() T { ... } и func() T { ... }() — разные вещи.
// Сложность: hard
package main_test

import "testing"

type Config struct {
	Workers int
	Name    string
}

var defaultConfig Config = func() Config {
	c := Config{Name: "svc"}
	c.Workers = 4
	return c
}

func TestDefaultConfig(t *testing.T) {
	if defaultConfig.Workers != 4 || defaultConfig.Name != "svc" {
		t.Errorf("defaultConfig = %+v", defaultConfig)
	}
}
