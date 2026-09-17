// maps_x082: Функция вернула nil-map
// Make the tests pass!
// I AM NOT DONE
//
// loadTags возвращает map тегов; для пустого ввода возвращается nil,
// и вызывающий код паникует при добавлении тега.
// Тренирует: nil-map можно читать, но нельзя в неё писать.
// Сложность: hard
package main_test

import (
	"strings"
	"testing"
)

func loadTags(s string) map[string]bool {
	if s == "" {
		return nil
	}
	tags := map[string]bool{}
	for _, t := range strings.Split(s, ",") {
		tags[t] = true
	}
	return tags
}

func TestLoadTags(t *testing.T) {
	tags := loadTags("")
	tags["new"] = true
	if len(tags) != 1 {
		t.Errorf("tags = %v", tags)
	}
	if len(loadTags("a,b")) != 2 {
		t.Errorf("loadTags(a,b) wrong")
	}
}
