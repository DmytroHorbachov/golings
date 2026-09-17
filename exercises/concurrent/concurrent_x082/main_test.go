// concurrent_x082: Once и разные функции
// Make the tests pass!
// I AM NOT DONE
//
// Один sync.Once используется и для загрузки конфигурации, и для открытия базы.
// Вторая функция никогда не выполняется.
// Тренирует: Once выполняет только первую переданную функцию.
// Сложность: hard
package main_test

import (
	"sync"
	"testing"
)

type App struct {
	once   sync.Once
	config string
	db     string
}

func (a *App) Config() string {
	a.once.Do(func() { a.config = "loaded" })
	return a.config
}

func (a *App) DB() string {
	a.once.Do(func() { a.db = "connected" })
	return a.db
}

func TestApp(t *testing.T) {
	var a App
	if a.Config() != "loaded" || a.DB() != "connected" {
		t.Errorf("config = %q, db = %q", a.config, a.db)
	}
}
