// functions_x062: Поиск по источникам
// Make the tests pass!
// I AM NOT DONE
//
// lookup ищет ключ последовательно в нескольких источниках-функциях
// и возвращает первое найденное значение.
// Тренирует: функции с результатом (value, ok).
// Сложность: medium
package main_test

import "testing"

type source func(key string) (string, bool)

func fromMap(m map[string]string) source {
	return func(k string) (string, bool) {
		v, ok := m[k]
		return v, ok
	}
}

func lookup(key string, sources ...source) (string, bool) {
	for _, s := range sources {
		v, _ := s(key)
		return v, true
	}
	return "", false
}

func TestLookup(t *testing.T) {
	env := fromMap(map[string]string{"PORT": "9000"})
	defaults := fromMap(map[string]string{"PORT": "80", "HOST": "localhost"})
	if v, ok := lookup("PORT", env, defaults); !ok || v != "9000" {
		t.Errorf("PORT = %q, %v", v, ok)
	}
	if v, ok := lookup("HOST", env, defaults); !ok || v != "localhost" {
		t.Errorf("HOST = %q, %v", v, ok)
	}
	if _, ok := lookup("USER", env, defaults); ok {
		t.Errorf("USER should not be found")
	}
}
