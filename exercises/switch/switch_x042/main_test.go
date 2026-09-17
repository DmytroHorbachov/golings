// switch_x042: Простой маршрутизатор
// Make the tests pass!
// I AM NOT DONE
//
// route возвращает имя обработчика по методу и пути.
// Тренирует: switch по составному ключу-строке.
// Сложность: medium
package main_test

import "testing"

func route(method, path string) string {
	switch method + " " + path {
	case "GET /users":
		return "listUsers"
	case "GET /users/new":
		return "createUser"
	default:
		return "listUsers"
	}
}

func TestRoute(t *testing.T) {
	cases := []struct{ m, p, want string }{
		{"GET", "/users", "listUsers"}, {"POST", "/users", "createUser"},
		{"DELETE", "/users", "notFound"}, {"GET", "/posts", "notFound"},
	}
	for _, c := range cases {
		if got := route(c.m, c.p); got != c.want {
			t.Errorf("route(%s %s) = %s, want %s", c.m, c.p, got, c.want)
		}
	}
}
