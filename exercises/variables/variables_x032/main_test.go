// variables_x032: Затенение именованного результата
// Make the tests pass!
// I AM NOT DONE
//
// Функция parsePort возвращает номер порта и ошибку через именованные результаты.
// Код не компилируется: компилятор жалуется на затенённые результаты.
// Тренирует: именованные возвращаемые значения и затенение при голом return.
// Сложность: hard
package main_test

import (
	"errors"
	"strconv"
	"testing"
)

func parsePort(s string) (n int, err error) {
	if s != "" {
		n, err := strconv.Atoi(s)
		if err != nil {
			return
		}
		if n <= 0 || n > 65535 {
			err = errors.New("port out of range")
		}
		return
	}
	return 80, nil
}

func TestParsePort(t *testing.T) {
	if n, err := parsePort("8080"); err != nil || n != 8080 {
		t.Errorf("parsePort(8080) = %d, %v", n, err)
	}
	if n, err := parsePort(""); err != nil || n != 80 {
		t.Errorf("parsePort(\"\") = %d, %v; want 80", n, err)
	}
	if _, err := parsePort("70000"); err == nil {
		t.Errorf("parsePort(70000) should fail")
	}
}
