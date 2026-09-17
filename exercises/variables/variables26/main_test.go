// variables26
// Make the tests pass!

// I AM NOT DONE
//
// Функция revoke должна снимать флаг, независимо от того, был ли он установлен.
// Повторный вызов сейчас снова включает флаг.
// Тренирует: оператор сброса битов &^ и отличие от XOR.
// Сложность: hard
package main_test

import "testing"

const (
	FlagRead uint8 = 1 << iota
	FlagWrite
	FlagAdmin
)

func revoke(perms, flag uint8) uint8 {
	return perms ^ flag
}

func TestRevoke(t *testing.T) {
	p := FlagRead | FlagAdmin
	p = revoke(p, FlagAdmin)
	if p != FlagRead {
		t.Errorf("after first revoke perms = %03b, want %03b", p, FlagRead)
	}
	p = revoke(p, FlagAdmin)
	if p != FlagRead {
		t.Errorf("after second revoke perms = %03b, want %03b", p, FlagRead)
	}
}
