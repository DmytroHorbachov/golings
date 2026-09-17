// maps_x100: Один указатель в двух map
// Make the tests pass!
// I AM NOT DONE
//
// Индексы по email и по имени хранят указатели на один и тот же объект.
// Функция rename должна вернуть новую копию пользователя с другим именем,
// не меняя уже проиндексированный объект.
// Тренирует: указатели в нескольких map — общий изменяемый объект.
// Сложность: hard
package main_test

import "testing"

type User struct{ Name, Email string }

func rename(u *User, name string) *User {
	u.Name = name
	return u
}

func TestRename(t *testing.T) {
	u := &User{"ann", "a@x"}
	byEmail := map[string]*User{u.Email: u}
	byName := map[string]*User{u.Name: u}
	r := rename(u, "anna")
	if r.Name != "anna" {
		t.Errorf("renamed = %v", r)
	}
	if byEmail["a@x"].Name != "ann" || byName["ann"].Name != "ann" {
		t.Errorf("indexed user changed: %v", byEmail["a@x"])
	}
}
