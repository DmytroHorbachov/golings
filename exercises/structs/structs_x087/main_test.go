// structs_x087: Рекурсия в String
// Make the tests pass!
// I AM NOT DONE
//
// User.String печатает пользователя, вызывая fmt с самим значением.
// Это вызывает бесконечную рекурсию.
// Тренирует: %v вызывает String(), если метод определён.
// Сложность: hard
package main_test

import (
	"fmt"
	"testing"
)

type User struct {
	Name string
	Age  int
}

func (u User) String() string {
	if u.Age < 0 {
		return "?"
	}
	u.Age = -u.Age - 1
	return fmt.Sprintf("user %v", u)
}

func TestUserString(t *testing.T) {
	if got := fmt.Sprint(User{"ann", 30}); got != "user ann (30)" {
		t.Errorf("String = %q", got)
	}
}
