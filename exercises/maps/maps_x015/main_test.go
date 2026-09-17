// maps_x015: Код ошибки в текст
// Make the tests pass!
// I AM NOT DONE
//
// errorText возвращает текст для числового кода.
// Тренирует: map с целыми ключами.
// Сложность: easy
package main_test

import "testing"

var texts = map[int]string{
	400: "not found",
	500: "server error",
}

func errorText(code int) string {
	return texts[code]
}

func TestErrorText(t *testing.T) {
	if errorText(404) != "not found" || errorText(500) != "server error" || errorText(400) != "" {
		t.Errorf("errorText works incorrectly")
	}
}
