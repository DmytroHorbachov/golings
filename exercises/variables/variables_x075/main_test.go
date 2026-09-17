// variables_x075: Строка с табуляцией
// Make the tests pass!
// I AM NOT DONE
//
// Функция row должна собрать строку таблицы: поля разделены символом табуляции.
// Тренирует: escape-последовательности в интерпретируемых строках.
// Сложность: easy
package main_test

import "testing"

func row(name, city string) string {
	sep := `\t`
	return name + sep + city
}

func TestRow(t *testing.T) {
	got := row("Ann", "Oslo")
	if got != "Ann"+string(rune(9))+"Oslo" {
		t.Errorf("row(Ann, Oslo) = %q, want tab-separated", got)
	}
}
