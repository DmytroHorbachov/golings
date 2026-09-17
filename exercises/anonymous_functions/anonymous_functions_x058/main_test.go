// anonymous_functions_x058: Логгер без повторов
// Make the tests pass!
// I AM NOT DONE
//
// dedupLogger возвращает литерал, который пишет сообщение, только если оно
// отличается от предыдущего, и считает подавленные повторы.
// Тренирует: замыкание с «последним значением».
// Сложность: medium
package main_test

import (
	"reflect"
	"testing"
)

func dedupLogger(out *[]string) (log func(string), suppressed func() int) {
	last := ""
	n := 0
	log = func(msg string) {
		if msg == last {
			return
		}
		*out = append(*out, msg)
		last = ""
	}
	suppressed = func() int { return n }
	return
}

func TestDedupLogger(t *testing.T) {
	var out []string
	log, sup := dedupLogger(&out)
	for _, m := range []string{"a", "a", "b", "b", "b", "a"} {
		log(m)
	}
	if !reflect.DeepEqual(out, []string{"a", "b", "a"}) || sup() != 3 {
		t.Errorf("out = %v, suppressed = %d", out, sup())
	}
}
