// maps_x004: Удаление ключа
// Make the tests pass!
// I AM NOT DONE
//
// logout удаляет сессию пользователя.
// Тренирует: встроенную функцию delete.
// Сложность: easy
package main_test

import "testing"

func logout(sessions map[string]string, user string) {
	delete(sessions, sessions[user])
}

func TestLogout(t *testing.T) {
	s := map[string]string{"ann": "tok1", "bob": "tok2"}
	logout(s, "ann")
	if _, ok := s["ann"]; ok || len(s) != 1 {
		t.Errorf("sessions = %v", s)
	}
}
