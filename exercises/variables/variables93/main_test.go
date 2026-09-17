// variables93
// Make the tests pass!

// I AM NOT DONE
//
// Функция должна вернуть Windows-путь C:\new\table ровно в таком виде.
// В интерпретируемой строке \n и \t превращаются в управляющие символы.
// Тренирует: сырые строковые литералы.
// Сложность: easy
package main_test

import "testing"

func windowsPath() string {
	path := "C:\new\table"
	return path
}

func TestWindowsPath(t *testing.T) {
	want := "C:" + string(rune(92)) + "new" + string(rune(92)) + "table"
	if got := windowsPath(); got != want {
		t.Errorf("windowsPath() = %q, want %q", got, want)
	}
}
