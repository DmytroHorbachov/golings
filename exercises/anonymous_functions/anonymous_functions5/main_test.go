// anonymous_functions5
// Make the tests pass!

// I AM NOT DONE
//
// countWords считает слова в нескольких горутинах, каждая пишет в общую
// захваченную переменную без синхронизации.
// Тренирует: литералы в горутинах, изменяющие общую переменную, создают гонку.
// Сложность: hard
package main_test

import (
	"strings"
	"testing"
)

func countWords(texts []string) int {
	total := 0
	done := make(chan struct{})
	for _, t := range texts {
		go func(t string) {
			total += len(strings.Fields(t))
			done <- struct{}{}
		}(t)
	}
	for range texts {
		<-done
	}
	return total
}

func TestCountWords(t *testing.T) {
	texts := make([]string, 50)
	for i := range texts {
		texts[i] = "one two three"
	}
	if got := countWords(texts); got != 150 {
		t.Errorf("countWords = %d, want 150", got)
	}
}
