// functions_x077: defer в цикле
// Make the tests pass!
// I AM NOT DONE
//
// processAll должна открыть, использовать и закрыть каждый ресурс,
// прежде чем переходить к следующему.
// Тренирует: defer срабатывает при выходе из функции, а не из итерации цикла.
// Сложность: hard
package main_test

import (
	"reflect"
	"testing"
)

type resource struct {
	name string
	log  *[]string
}

func open(name string, log *[]string) *resource {
	*log = append(*log, "open "+name)
	return &resource{name, log}
}

func (r *resource) Close() { *r.log = append(*r.log, "close "+r.name) }

func processAll(names []string) []string {
	var log []string
	for _, n := range names {
		r := open(n, &log)
		defer r.Close()
		log = append(log, "use "+r.name)
	}
	return log
}

func TestProcessAll(t *testing.T) {
	want := []string{"open a", "use a", "close a", "open b", "use b", "close b"}
	if got := processAll([]string{"a", "b"}); !reflect.DeepEqual(got, want) {
		t.Errorf("processAll = %v, want %v", got, want)
	}
}
