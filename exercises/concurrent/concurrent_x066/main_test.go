// concurrent_x066: Отключение ветки select
// Make the tests pass!
// I AM NOT DONE
//
// mergeTwo читает из двух каналов, пока оба не закроются. Закрытый канал
// должен быть исключён из select присваиванием nil.
// Тренирует: nil-канал блокирует ветку select навсегда.
// Сложность: medium
package main_test

import (
	"testing"
	"time"
)

func mergeTwo(a, b <-chan int) int {
	sum := 0
	for a != nil || b != nil {
		select {
		case v, ok := <-a:
			if !ok {
				continue
			}
			sum += v
		case v, ok := <-b:
			if !ok {
				continue
			}
			sum += v
		}
	}
	return sum
}

func TestMergeTwo(t *testing.T) {
	a, b := make(chan int, 2), make(chan int, 1)
	a <- 1
	a <- 2
	b <- 10
	close(a)
	close(b)
	res := make(chan int, 1)
	go func() { res <- mergeTwo(a, b) }()
	select {
	case got := <-res:
		if got != 13 {
			t.Errorf("mergeTwo = %d", got)
		}
	case <-time.After(time.Second):
		t.Fatal("mergeTwo never finished")
	}
}
