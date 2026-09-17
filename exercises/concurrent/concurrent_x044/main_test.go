// concurrent_x044: Производитель и потребитель
// Make the tests pass!
// I AM NOT DONE
//
// consume обрабатывает сообщения, пока не получит сигнал done
// и очередь не опустеет.
// Тренирует: select между каналом данных и каналом остановки.
// Сложность: medium
package main_test

import (
	"testing"
	"time"
)

func consume(msgs <-chan string, done <-chan struct{}) []string {
	var out []string
	for {
		select {
		case m := <-msgs:
			out = append(out, m)
		case <-done:
			return out
		}
	}
}

func TestConsume(t *testing.T) {
	msgs := make(chan string, 30)
	done := make(chan struct{})
	close(done)
	for i := 0; i < 20; i++ {
		msgs <- "m"
	}
	result := make(chan []string, 1)
	go func() { result <- consume(msgs, done) }()
	select {
	case got := <-result:
		if len(got) != 20 {
			t.Errorf("consumed %v", got)
		}
	case <-time.After(time.Second):
		t.Fatal("consume is stuck")
	}
}
