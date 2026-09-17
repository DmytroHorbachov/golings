// variables25
// Make the tests pass!

// I AM NOT DONE
//
// Права доступа хранятся как набор битовых флагов.
// Функция can должна проверять, установлен ли нужный флаг.
// Тренирует: 1 << iota и побитовые операции & и |.
// Сложность: medium
package main_test

import "testing"

type Perm uint8

const (
	Read Perm = iota
	Write
	Exec
)

func can(p, flag Perm) bool {
	return p|flag != 0
}

func TestCan(t *testing.T) {
	p := Read | Exec
	if !can(p, Read) || !can(p, Exec) {
		t.Errorf("Read|Exec should allow Read and Exec")
	}
	if can(p, Write) {
		t.Errorf("Read|Exec should not allow Write")
	}
}
