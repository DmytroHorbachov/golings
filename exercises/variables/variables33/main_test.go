// variables33
// Make the tests pass!

// I AM NOT DONE
//
// Функция configure должна установить глобальный уровень логирования.
// После вызова глобальная переменная не меняется.
// Тренирует: локальная переменная скрывает одноимённую переменную пакета.
// Сложность: hard
package main_test

import "testing"

var logLevel = "info"

func configure(debug bool) {
	if debug {
		logLevel := "debug"
		_ = logLevel
	}
}

func TestConfigure(t *testing.T) {
	logLevel = "info"
	configure(false)
	if logLevel != "info" {
		t.Errorf("configure(false): logLevel = %q, want info", logLevel)
	}
	configure(true)
	if logLevel != "debug" {
		t.Errorf("configure(true): logLevel = %q, want debug", logLevel)
	}
}
