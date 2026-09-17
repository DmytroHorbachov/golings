// slices_x094: Индекс при нулевой длине
// Make the tests pass!
// I AM NOT DONE
//
// collectIDs заполняет срез, созданный с ёмкостью, по индексу и паникует.
// Тренирует: ёмкость не даёт права обращаться к элементам за длиной.
// Сложность: hard
package main_test

import (
	"reflect"
	"testing"
)

func collectIDs(users []map[string]int) []int {
	ids := make([]int, 0, len(users))
	for i, u := range users {
		ids[i] = u["id"]
	}
	return ids
}

func TestCollectIDs(t *testing.T) {
	got := collectIDs([]map[string]int{{"id": 3}, {"id": 7}})
	if !reflect.DeepEqual(got, []int{3, 7}) {
		t.Errorf("collectIDs = %v", got)
	}
}
