// primitive_types_x002: Подстрока
// Make the tests pass!
// I AM NOT DONE
//
// areaCode должна вернуть три цифры кода из номера вида "(495)1234567".
// Тренирует: срез строки s[i:j] (j не включается).
// Сложность: easy
package main_test

import "testing"

func areaCode(phone string) string {
	return phone[0:3]
}

func TestAreaCode(t *testing.T) {
	if got := areaCode("(495)1234567"); got != "495" {
		t.Errorf("areaCode = %q, want 495", got)
	}
}
