// variables_x054: Именованный результат и recover
// Make the tests pass!
// I AM NOT DONE
//
// Функция safeDiv должна вернуть ошибку вместо паники при делении на ноль.
// Паника перехватывается, но вызывающий код получает nil.
// Тренирует: отложенная функция может изменить только именованные результаты.
// Сложность: hard
package main_test

import (
	"fmt"
	"testing"
)

func safeDiv(a, b int) (int, error) {
	var err error
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("recovered: %v", r)
		}
	}()
	return a / b, err
}

func TestSafeDiv(t *testing.T) {
	if q, err := safeDiv(9, 3); err != nil || q != 3 {
		t.Errorf("safeDiv(9, 3) = %d, %v", q, err)
	}
	if _, err := safeDiv(1, 0); err == nil {
		t.Errorf("safeDiv(1, 0) should return an error")
	}
}
