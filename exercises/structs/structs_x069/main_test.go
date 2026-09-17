// structs_x069: Скрытый метод встроенного типа
// Make the tests pass!
// I AM NOT DONE
//
// Admin встраивает User и переопределяет Describe, но должен включать
// описание пользователя. Сейчас метод вызывает сам себя.
// Тренирует: доступ к скрытому методу встроенного типа через его имя.
// Сложность: hard
package main_test

import "testing"

type User struct{ Name string }

func (u User) Describe() string { return "user " + u.Name }

type Admin struct {
	User
	Level int
}

func (a Admin) Describe() string {
	if a.Level > 100 {
		return "admin"
	}
	a.Level += 100
	return a.Describe() + " (admin)"
}

func TestAdminDescribe(t *testing.T) {
	if got := (Admin{User{"ann"}, 1}).Describe(); got != "user ann (admin)" {
		t.Errorf("Describe = %q", got)
	}
}
