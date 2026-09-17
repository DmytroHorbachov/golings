// concurrent_x055: Барьер готовности
// Make the tests pass!
// I AM NOT DONE
//
// waitReady запускает сервисы и ждёт, пока каждый сообщит о готовности.
// Тренирует: сбор сигналов от N горутин.
// Сложность: medium
package main_test

import (
	"sort"
	"testing"
	"time"
)

func waitReady(names []string) []string {
	ready := make(chan string)
	for _, n := range names {
		go func(n string) {
			time.Sleep(time.Millisecond)
			ready <- n
		}(n)
	}
	var got []string
	got = append(got, <-ready)
	sort.Strings(got)
	return got
}

func TestWaitReady(t *testing.T) {
	got := waitReady([]string{"db", "api", "cache"})
	if len(got) != 3 || got[0] != "api" {
		t.Errorf("waitReady = %v", got)
	}
}
