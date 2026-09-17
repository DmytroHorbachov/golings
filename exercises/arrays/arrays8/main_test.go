// arrays8
// Make the tests pass!

// I AM NOT DONE
//
// Таблица сообщений заполняется индексированным литералом.
// Код не компилируется: один и тот же индекс указан дважды.
// Тренирует: индексы в литерале массива должны быть уникальными.
// Сложность: hard
package main_test

import "testing"

const (
	OK       = 0
	NotFound = 1
	Denied   = 2
)

var text = [3]string{
	OK:     "ok",
	OK:     "not found",
	Denied: "denied",
}

func TestText(t *testing.T) {
	if text[OK] != "ok" || text[NotFound] != "not found" || text[Denied] != "denied" {
		t.Errorf("text = %v", text)
	}
}
