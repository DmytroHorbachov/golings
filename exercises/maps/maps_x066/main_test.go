// maps_x066: Многие ко многим
// Make the tests pass!
// I AM NOT DONE
//
// byCourse превращает map «студент -> курсы» в «курс -> студенты» (отсортированные).
// Тренирует: обращение связи многие-ко-многим.
// Сложность: medium
package main_test

import (
	"reflect"
	"sort"
	"testing"
)

func byCourse(enroll map[string][]string) map[string][]string {
	out := map[string][]string{}
	for student, courses := range enroll {
		for _, c := range courses {
			out[student] = append(out[student], c)
		}
	}
	return out
}

func TestByCourse(t *testing.T) {
	got := byCourse(map[string][]string{"ann": {"go", "db"}, "bob": {"go"}})
	if !reflect.DeepEqual(got, map[string][]string{"go": {"ann", "bob"}, "db": {"ann"}}) {
		t.Errorf("byCourse = %v", got)
	}
}
