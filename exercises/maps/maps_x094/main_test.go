// maps_x094: false вместо удаления
// Make the tests pass!
// I AM NOT DONE
//
// online отмечает пользователей в сети в map[string]bool. При выходе
// пользователь помечается false, и счётчик онлайна становится неверным.
// Тренирует: len(map) считает все ключи, включая значения false.
// Сложность: hard
package main_test

import "testing"

type Online map[string]bool

func (o Online) Login(u string) { o[u] = true }

func (o Online) Logout(u string) {
	o[u] = false
}

func (o Online) Count() int { return len(o) }

func TestOnline(t *testing.T) {
	o := Online{}
	o.Login("ann")
	o.Login("bob")
	o.Logout("ann")
	if o.Count() != 1 {
		t.Errorf("Count = %d, want 1", o.Count())
	}
}
