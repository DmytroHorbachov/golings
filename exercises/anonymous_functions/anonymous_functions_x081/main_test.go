// anonymous_functions_x081: Горутина и изменённая переменная
// Make the tests pass!
// I AM NOT DONE
//
// report запускает горутину, которая должна отправить исходное имя.
// Горутина читает переменную позже, когда она уже изменена.
// Тренирует: литерал в горутине читает переменную в момент выполнения.
// Сложность: hard
package main_test

import "testing"

func report() string {
	name := "draft"
	ready := make(chan struct{})
	out := make(chan string)
	go func() {
		<-ready
		out <- name
	}()
	name = "final"
	close(ready)
	return <-out
}

func TestReport(t *testing.T) {
	if got := report(); got != "draft" {
		t.Errorf("report = %q, want draft", got)
	}
}
