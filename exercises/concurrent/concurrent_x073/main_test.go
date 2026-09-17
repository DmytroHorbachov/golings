// concurrent_x073: Порядок захвата мьютексов
// Make the tests pass!
// I AM NOT DONE
//
// transfer блокирует два счёта. Встречные переводы захватывают мьютексы
// в разном порядке и взаимно блокируются.
// Тренирует: единый порядок захвата блокировок.
// Сложность: hard
package main_test

import (
	"sync"
	"testing"
	"time"
)

type Account struct {
	ID      int
	mu      sync.Mutex
	Balance int
}

func transfer(from, to *Account, amount int) {
	first, second := from, to
	first.mu.Lock()
	defer first.mu.Unlock()
	time.Sleep(time.Millisecond)
	second.mu.Lock()
	defer second.mu.Unlock()
	from.Balance -= amount
	to.Balance += amount
}

func TestTransfer(t *testing.T) {
	a := &Account{ID: 1, Balance: 100}
	b := &Account{ID: 2, Balance: 100}
	var wg sync.WaitGroup
	for i := 0; i < 20; i++ {
		wg.Add(2)
		go func() { defer wg.Done(); transfer(a, b, 1) }()
		go func() { defer wg.Done(); transfer(b, a, 1) }()
	}
	done := make(chan struct{})
	go func() { wg.Wait(); close(done) }()
	select {
	case <-done:
		if a.Balance != 100 || b.Balance != 100 {
			t.Errorf("balances = %d, %d", a.Balance, b.Balance)
		}
	case <-time.After(2 * time.Second):
		t.Fatal("deadlock between transfers")
	}
}
