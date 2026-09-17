// switch87
// Make the tests pass!

// I AM NOT DONE
//
// describe переводит балл 1–5 в слово, а для остальных возвращает ошибку.
// Тренирует: switch с возвратом ошибки в default.
// Сложность: medium
package main_test

import (
	"errors"
	"testing"
)

func describe(score int) (string, error) {
	switch score {
	case 1, 2:
		return "bad", nil
	case 3:
		return "ok", nil
	case 4:
		return "great", nil
	}
	return "", nil
}

func TestDescribe(t *testing.T) {
	_ = errors.New
	cases := map[int]string{1: "bad", 2: "bad", 3: "ok", 4: "great", 5: "great"}
	for in, want := range cases {
		if got, err := describe(in); err != nil || got != want {
			t.Errorf("describe(%d) = %s, %v", in, got, err)
		}
	}
	if _, err := describe(0); err == nil {
		t.Errorf("describe(0) should fail")
	}
}
