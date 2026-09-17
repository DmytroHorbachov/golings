// structs45
// Make the tests pass!

// I AM NOT DONE
//
// Notifier встраивает интерфейс Sender. Структура создана без Sender,
// и вызов метода паникует.
// Тренирует: встроенный интерфейс — поле, которое по умолчанию nil.
// Сложность: hard
package main_test

import "testing"

type Sender interface{ Send(msg string) bool }

type Notifier struct {
	Sender
	Channel string
}

func (n Notifier) Notify(msg string) bool {
	return n.Send("[" + n.Channel + "] " + msg)
}

type okSender struct{}

func (okSender) Send(string) bool { return true }

func TestNotifier(t *testing.T) {
	if (Notifier{Channel: "ops"}).Notify("hi") {
		t.Errorf("Notify without sender should return false")
	}
	if !(Notifier{okSender{}, "ops"}).Notify("hi") {
		t.Errorf("Notify with sender should return true")
	}
}
