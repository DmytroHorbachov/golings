// if_x095: Побочный эффект в условии
// Make the tests pass!
// I AM NOT DONE
//
// nextTicket выдаёт номер из очереди, если очередь не пуста.
// Вызов pop происходит и в условии, и в теле — номера теряются.
// Тренирует: выражения в условии if вычисляются при каждой проверке.
// Сложность: hard
package main_test

import (
	"reflect"
	"testing"
)

type Queue struct{ items []int }

func (q *Queue) pop() (int, bool) {
	if len(q.items) == 0 {
		return 0, false
	}
	v := q.items[0]
	q.items = q.items[1:]
	return v, true
}

func nextTicket(q *Queue) int {
	if _, ok := q.pop(); ok {
		v, _ := q.pop()
		return v
	}
	return -1
}

func TestNextTicket(t *testing.T) {
	q := &Queue{items: []int{1, 2, 3}}
	got := []int{nextTicket(q), nextTicket(q), nextTicket(q), nextTicket(q)}
	if want := []int{1, 2, 3, -1}; !reflect.DeepEqual(got, want) {
		t.Errorf("tickets = %v, want %v", got, want)
	}
}
