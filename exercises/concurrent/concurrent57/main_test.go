// concurrent57
// Make the tests pass!

// I AM NOT DONE
//
// Deposit вызывает Balance, и оба метода захватывают один мьютекс.
// Мьютекс Go не реентрантный — вызов зависает навсегда.
// Тренирует: выделение внутренних функций без блокировки.
// Сложность: hard
package main_test

import (
	"sync"
	"testing"
	"time"
)

type Wallet struct {
	mu  sync.Mutex
	bal int
}

func (w *Wallet) balance() int { return w.bal }

func (w *Wallet) Balance() int {
	w.mu.Lock()
	defer w.mu.Unlock()
	return w.balance()
}

func (w *Wallet) Deposit(n int) int {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.bal += n
	return w.Balance()
}

func TestDeposit(t *testing.T) {
	w := &Wallet{}
	res := make(chan int, 1)
	go func() { res <- w.Deposit(5) }()
	select {
	case got := <-res:
		if got != 5 {
			t.Errorf("Deposit = %d", got)
		}
	case <-time.After(time.Second):
		t.Fatal("Deposit deadlocked")
	}
}
