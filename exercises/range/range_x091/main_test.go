// range_x091: return в колбэке
// Make the tests pass!
// I AM NOT DONE
//
// forEach вызывает функцию для каждого элемента. findFirstBig пытается
// остановить перебор через return, но return выходит только из колбэка.
// Тренирует: return внутри функционального литерала не прерывает внешний цикл.
// Сложность: hard
package main_test

import "testing"

func forEach(s []int, f func(int)) {
	for _, v := range s {
		f(v)
	}
}

func findFirstBig(s []int) (int, int) {
	found, calls := -1, 0
	forEach(s, func(v int) {
		calls++
		if v > 10 {
			found = v
			return
		}
	})
	return found, calls
}

func TestFindFirstBig(t *testing.T) {
	found, calls := findFirstBig([]int{1, 15, 20, 3})
	if found != 15 || calls != 2 {
		t.Errorf("findFirstBig = %d after %d calls; want 15 after 2", found, calls)
	}
}
