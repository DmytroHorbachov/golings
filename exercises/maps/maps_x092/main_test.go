// maps_x092: Ключ с указателем внутри
// Make the tests pass!
// I AM NOT DONE
//
// Посещения учитываются по ключу {город, *Пользователь}. Одинаковые
// по содержимому пользователи считаются разными.
// Тренирует: структура-ключ сравнивает поля-указатели по адресу.
// Сложность: hard
package main_test

import "testing"

type Person struct{ ID int }

type visitKey struct {
	City   string
	Person *Person
}

func key(city string, p Person) visitKey { return visitKey{city, &p} }

func TestVisits(t *testing.T) {
	visits := map[visitKey]int{}
	visits[key("Kazan", Person{ID: 7})]++
	visits[key("Kazan", Person{ID: 7})]++
	if len(visits) != 1 || visits[key("Kazan", Person{ID: 7})] != 2 {
		t.Errorf("visits = %v", visits)
	}
}
