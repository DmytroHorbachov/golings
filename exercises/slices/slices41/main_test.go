// slices41
// Make the tests pass!

// I AM NOT DONE
//
// addAndRename добавляет пользователя и затем меняет имя первого через
// ранее взятый указатель. Изменение теряется.
// Тренирует: append может перенести элементы в новый массив, старые указатели устаревают.
// Сложность: hard
package main_test

import "testing"

type User struct{ Name string }

func addAndRename(users []User) []User {
	first := &users[0]
	users = append(users, User{"new"})
	first.Name = "admin"
	return users
}

func TestAddAndRename(t *testing.T) {
	users := []User{{"ann"}}
	users = addAndRename(users)
	if users[0].Name != "admin" || users[1].Name != "new" {
		t.Errorf("users = %v", users)
	}
}
