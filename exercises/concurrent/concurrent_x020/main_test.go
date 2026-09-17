// concurrent_x020: Результат-структура
// Make the tests pass!
// I AM NOT DONE
//
// fetch возвращает через канал структуру с данными и ошибкой.
// Тренирует: каналы структур.
// Сложность: easy
package main_test

import "testing"

type Result struct {
	Body string
	Err  error
}

func fetch(url string) Result {
	ch := make(chan Result, 1)
	go func() {
		ch <- Result{Err: nil, Body: ""}
	}()
	return <-ch
}

func TestFetch(t *testing.T) {
	if r := fetch("/a"); r.Err != nil || r.Body != "page /a" {
		t.Errorf("fetch = %+v", r)
	}
}
