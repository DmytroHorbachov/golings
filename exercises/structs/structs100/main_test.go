// structs100
// Make the tests pass!

// I AM NOT DONE
//
// В структуру User добавили поле Email между Name и Age.
// Позиционные литералы перестали компилироваться (или получили не те значения).
// Тренирует: литералы с именами полей устойчивы к изменениям структуры.
// Сложность: hard
package main_test

import "testing"

type User struct {
	Name  string
	Email string
	Age   int
}

func defaultUser() User {
	return User{"guest", 18}
}

func TestDefaultUser(t *testing.T) {
	u := defaultUser()
	if u.Name != "guest" || u.Age != 18 || u.Email != "" {
		t.Errorf("defaultUser = %+v", u)
	}
}
