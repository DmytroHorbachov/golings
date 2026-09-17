// if_x081: Необязательный флаг-указатель
// Make the tests pass!
// I AM NOT DONE
//
// verbose принимает *bool: nil означает «не задано» и трактуется как false.
// Сейчас при nil функция паникует.
// Тренирует: разыменование nil-указателя в условии.
// Сложность: hard
package main_test

import "testing"

func verbose(flag *bool) string {
	if *flag {
		return "loud"
	}
	return "quiet"
}

func TestVerbose(t *testing.T) {
	yes, no := true, false
	if verbose(&yes) != "loud" || verbose(&no) != "quiet" {
		t.Errorf("verbose with explicit flag is wrong")
	}
	if got := verbose(nil); got != "quiet" {
		t.Errorf("verbose(nil) = %q, want quiet", got)
	}
}
