// if_x098: Все флаги или любой
// Make the tests pass!
// I AM NOT DONE
//
// hasAll должна вернуть true, только если установлены ВСЕ биты из mask.
// Сейчас хватает одного совпавшего бита.
// Тренирует: разницу между flags&mask != 0 и flags&mask == mask.
// Сложность: hard
package main_test

import "testing"

const (
	Read uint8 = 1 << iota
	Write
	Exec
)

func hasAll(flags, mask uint8) bool {
	if flags&mask != 0 {
		return true
	}
	return false
}

func TestHasAll(t *testing.T) {
	if !hasAll(Read|Write|Exec, Read|Write) {
		t.Errorf("rwx has rw")
	}
	if hasAll(Read, Read|Write) {
		t.Errorf("r does not have rw")
	}
	if hasAll(Exec, Read|Write) {
		t.Errorf("x does not have rw")
	}
}
