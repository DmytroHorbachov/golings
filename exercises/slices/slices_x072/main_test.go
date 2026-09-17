// slices_x072: nil и пустой срез в JSON
// Make the tests pass!
// I AM NOT DONE
//
// API должно возвращать "[]" для пустого списка, а не "null".
// Тренирует: encoding/json кодирует nil-срез как null.
// Сложность: hard
package main_test

import (
	"encoding/json"
	"testing"
)

func activeUsers(all []string, active map[string]bool) []string {
	var out []string
	for _, u := range all {
		if active[u] {
			out = append(out, u)
		}
	}
	return out
}

func TestActiveUsersJSON(t *testing.T) {
	b, _ := json.Marshal(activeUsers([]string{"ann"}, nil))
	if string(b) != "[]" {
		t.Errorf("json = %s, want []", b)
	}
	b, _ = json.Marshal(activeUsers([]string{"ann", "bob"}, map[string]bool{"bob": true}))
	if string(b) != `["bob"]` {
		t.Errorf("json = %s", b)
	}
}
