// functions_x092: Сравнение функций
// Make the tests pass!
// I AM NOT DONE
//
// hasPlugin должна проверить, зарегистрирован ли плагин с таким именем.
// Код не компилируется: структуры сравниваются через ==.
// Тренирует: функции несравнимы, а значит несравнимы и структуры с полями-функциями.
// Сложность: hard
package main_test

import "testing"

type Plugin struct {
	Name string
	Run  func() string
}

func hasPlugin(ps []Plugin, target Plugin) bool {
	for _, p := range ps {
		if p == target {
			return true
		}
	}
	return false
}

func TestHasPlugin(t *testing.T) {
	gzip := Plugin{"gzip", func() string { return "compressed" }}
	auth := Plugin{"auth", func() string { return "ok" }}
	if !hasPlugin([]Plugin{gzip, auth}, Plugin{Name: "auth"}) {
		t.Errorf("auth plugin should be found")
	}
	if hasPlugin([]Plugin{gzip}, auth) {
		t.Errorf("auth plugin should not be found")
	}
}
