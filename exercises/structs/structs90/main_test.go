// structs90
// Make the tests pass!

// I AM NOT DONE
//
// Event сериализуется в JSON с ключами "id", "title", "tags" и обратно без потерь.
// Тренирует: теги и json.Marshal/Unmarshal.
// Сложность: medium
package main_test

import (
	"encoding/json"
	"reflect"
	"testing"
)

type Event struct {
	ID    int      `json:"id"`
	Title string   `json:"name"`
	Tags  []string `json:"-"`
}

func TestEventJSON(t *testing.T) {
	in := Event{7, "launch", []string{"go"}}
	b, err := json.Marshal(in)
	if err != nil || string(b) != `{"id":7,"title":"launch","tags":["go"]}` {
		t.Fatalf("json = %s, %v", b, err)
	}
	var out Event
	if err := json.Unmarshal(b, &out); err != nil || !reflect.DeepEqual(in, out) {
		t.Errorf("round trip = %+v, %v", out, err)
	}
}
