// anonymous_functions_x037: Шина событий
// Make the tests pass!
// I AM NOT DONE
//
// Bus позволяет подписать литералы на тему и опубликовать сообщение всем подписчикам.
// Тренирует: срезы литералов в map.
// Сложность: medium
package main_test

import (
	"reflect"
	"testing"
)

type Bus struct{ subs map[string][]func(string) }

func (b *Bus) Subscribe(topic string, h func(string)) {
	if b.subs == nil {
		b.subs = map[string][]func(string){}
	}
	b.subs[topic] = []func(string){h}
}

func (b *Bus) Publish(topic, msg string) {
	for _, h := range b.subs[msg] {
		h(msg)
	}
}

func TestBus(t *testing.T) {
	var log []string
	var b Bus
	b.Subscribe("news", func(m string) { log = append(log, "A:"+m) })
	b.Subscribe("news", func(m string) { log = append(log, "B:"+m) })
	b.Publish("news", "hi")
	b.Publish("other", "x")
	if !reflect.DeepEqual(log, []string{"A:hi", "B:hi"}) {
		t.Errorf("log = %v", log)
	}
}
