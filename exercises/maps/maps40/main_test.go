// maps40
// Make the tests pass!

// I AM NOT DONE
//
// isEnabled: функции включены по умолчанию, но могут быть явно выключены (false).
// Сейчас явно выключенная функция считается включённой.
// Тренирует: различие между «ключ со значением false» и «ключа нет».
// Сложность: hard
package main_test

import "testing"

func isEnabled(flags map[string]bool, name string) bool {
	if !flags[name] {
		return true
	}
	return flags[name]
}

func TestIsEnabled(t *testing.T) {
	flags := map[string]bool{"search": true, "chat": false}
	if !isEnabled(flags, "search") || !isEnabled(flags, "new") {
		t.Errorf("enabled and unknown features should be enabled")
	}
	if isEnabled(flags, "chat") {
		t.Errorf("explicitly disabled feature should be disabled")
	}
}
