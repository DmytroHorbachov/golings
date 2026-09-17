// generics93
// Make the tests pass!

// I AM NOT DONE
//
// Метод Queue.Push объявлен с получателем Queue без [T].
// Код не компилируется.
// Тренирует: получатель метода обобщённого типа должен перечислять параметры.
// Сложность: hard
package main_test

import "testing"

type Queue[T any] struct{ items []T }

func (q *Queue) Push(v T) { q.items = append(q.items, v) }

func (q *Queue[T]) Len() int { return len(q.items) }

func TestQueuePush(t *testing.T) {
	var q Queue[string]
	q.Push("a")
	if q.Len() != 1 {
		t.Errorf("Len = %d", q.Len())
	}
}
