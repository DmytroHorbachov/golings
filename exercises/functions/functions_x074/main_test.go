// functions_x074: defer nil-функции
// Make the tests pass!
// I AM NOT DONE
//
// process должна выполнить работу и вызвать onDone, если он задан.
// При onDone == nil функция паникует, причём уже после работы.
// Тренирует: defer nil-функции паникует в момент выполнения отложенного вызова.
// Сложность: hard
package main_test

import "testing"

func process(items []int, onDone func(int)) int {
	sum := 0
	defer onDone(len(items))
	for _, v := range items {
		sum += v
	}
	return sum
}

func TestProcess(t *testing.T) {
	done := -1
	if got := process([]int{1, 2}, func(n int) { done = n }); got != 3 || done != 2 {
		t.Errorf("process with callback = %d, done=%d", got, done)
	}
	if got := process([]int{4}, nil); got != 4 {
		t.Errorf("process without callback = %d, want 4", got)
	}
}
