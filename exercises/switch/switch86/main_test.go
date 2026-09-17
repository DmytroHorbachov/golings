// switch86
// Make the tests pass!

// I AM NOT DONE
//
// encode переводит число в десятичную или шестнадцатеричную строку.
// Код не компилируется: переменная, объявленная в case, не видна после switch.
// Тренирует: каждая ветка case — отдельная область видимости.
// Сложность: hard
package main_test

import (
	"errors"
	"strconv"
	"testing"
)

func encode(format string, v int) (string, error) {
	switch format {
	case "dec":
		s := strconv.Itoa(v)
	case "hex":
		s := strconv.FormatInt(int64(v), 16)
	default:
		return "", errors.New("unknown format")
	}
	return s, nil
}

func TestEncode(t *testing.T) {
	if got, err := encode("dec", 255); err != nil || got != "255" {
		t.Errorf("encode(dec) = %s, %v", got, err)
	}
	if got, err := encode("hex", 255); err != nil || got != "ff" {
		t.Errorf("encode(hex) = %s, %v", got, err)
	}
	if _, err := encode("bin", 1); err == nil {
		t.Errorf("encode(bin) should fail")
	}
}
