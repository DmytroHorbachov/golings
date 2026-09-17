// concurrent92
// Make the tests pass!

// I AM NOT DONE
//
// Server принимает задачи через канал; Shutdown закрывает вход и ждёт,
// пока будут обработаны все принятые задачи.
// Тренирует: закрытие канала и ожидание воркера.
// Сложность: medium
package main_test

import (
	"testing"
	"time"
)

type Server struct {
	tasks chan int
	done  chan struct{}
	sum   int
}

func NewServer() *Server {
	s := &Server{tasks: make(chan int, 10), done: make(chan struct{})}
	go func() {
		defer close(s.done)
		for t := range s.tasks {
			time.Sleep(time.Millisecond)
			s.sum += t
		}
	}()
	return s
}

func (s *Server) Shutdown() {
	close(s.done)
}

func TestShutdown(t *testing.T) {
	s := NewServer()
	for i := 1; i <= 5; i++ {
		s.tasks <- i
	}
	s.Shutdown()
	if s.sum != 15 {
		t.Errorf("sum = %d, want 15", s.sum)
	}
}
