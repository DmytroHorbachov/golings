// maps35
// Make the tests pass!

// I AM NOT DONE
//
// emailOf возвращает email пользователя из map структур.
// Тренирует: доступ к полю значения map.
// Сложность: easy
package main_test

import "testing"

type User struct {
	Name  string
	Email string
}

func emailOf(users map[int]User, id int) string {
	return users[id].Name
}

func TestEmailOf(t *testing.T) {
	users := map[int]User{7: {"ann", "ann@x.io"}}
	if emailOf(users, 7) != "ann@x.io" || emailOf(users, 8) != "" {
		t.Errorf("emailOf works incorrectly")
	}
}
