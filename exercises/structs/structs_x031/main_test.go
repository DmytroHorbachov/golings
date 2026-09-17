// structs_x031: Метод на указателе вызывается у значения
// Make the tests pass!
// I AM NOT DONE
//
// Inc вызывается у адресуемой переменной-значения; Go сам берёт адрес.
// Тренирует: автоматическое взятие адреса при вызове pointer-метода.
// Сложность: easy
package main_test

import "testing"

type Hits struct{ N int }

func (h *Hits) Inc() {
	h.N += 2
}

func TestInc(t *testing.T) {
	var h Hits
	h.Inc()
	h.Inc()
	if h.N != 2 {
		t.Errorf("N = %d", h.N)
	}
}
