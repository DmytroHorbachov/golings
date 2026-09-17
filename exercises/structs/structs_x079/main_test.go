// structs_x079: Копия с общей map
// Make the tests pass!
// I AM NOT DONE
//
// Clone конфигурации должен быть независим, но изменения в копии видны в оригинале.
// Тренирует: при копировании структуры map копируется по ссылке.
// Сложность: hard
package main_test

import "testing"

type Config struct {
	Name string
	Opts map[string]string
}

func (c Config) Clone() Config {
	return c
}

func TestClone(t *testing.T) {
	c := Config{"base", map[string]string{"debug": "off"}}
	d := c.Clone()
	d.Opts["debug"] = "on"
	if c.Opts["debug"] != "off" {
		t.Errorf("original changed: %v", c.Opts)
	}
}
