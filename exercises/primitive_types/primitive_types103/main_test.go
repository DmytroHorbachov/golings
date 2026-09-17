// primitive_types103
// Make the tests pass!

// I AM NOT DONE
//
// parsePercent превращает "45.5%" в 0.455. Строка без знака % — ошибка.
// Тренирует: strings.CutSuffix-подобную логику и strconv.ParseFloat.
// Сложность: medium
package main_test

import (
	"errors"
	"strconv"
	"strings"
	"testing"
)

func parsePercent(s string) (float64, error) {
	v, err := strconv.ParseFloat(s, 64)
	return v, err
}

func TestParsePercent(t *testing.T) {
	_, _ = errors.New, strings.HasSuffix
	if v, err := parsePercent("45.5%"); err != nil || v != 0.455 {
		t.Errorf("parsePercent(45.5%%) = %v, %v", v, err)
	}
	if v, err := parsePercent("100%"); err != nil || v != 1 {
		t.Errorf("parsePercent(100%%) = %v, %v", v, err)
	}
	if _, err := parsePercent("45"); err == nil {
		t.Errorf("parsePercent(45) should fail")
	}
}
