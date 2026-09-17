// switch_x086: Накопление через fallthrough
// Make the tests pass!
// I AM NOT DONE
//
// features возвращает набор функций тарифа: "pro" включает всё из "plus",
// а "plus" — всё из "basic". Для "pro" сейчас не хватает функций.
// Тренирует: fallthrough нужно писать явно в каждой ветке.
// Сложность: hard
package main_test

import (
	"reflect"
	"testing"
)

func features(plan string) []string {
	var f []string
	switch plan {
	case "pro":
		f = append(f, "api")
	case "plus":
		f = append(f, "export")
		fallthrough
	case "basic":
		f = append(f, "editor")
	}
	return f
}

func TestFeatures(t *testing.T) {
	if got := features("pro"); !reflect.DeepEqual(got, []string{"api", "export", "editor"}) {
		t.Errorf("features(pro) = %v", got)
	}
	if got := features("plus"); !reflect.DeepEqual(got, []string{"export", "editor"}) {
		t.Errorf("features(plus) = %v", got)
	}
}
