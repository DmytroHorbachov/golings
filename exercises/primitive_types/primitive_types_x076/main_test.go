// primitive_types_x076: Логический сдвиг вправо
// Make the tests pass!
// I AM NOT DONE
//
// unsignedShift должна сдвигать биты int32 вправо, заполняя освободившиеся
// старшие биты нулями (как >>> в Java).
// Тренирует: для знаковых типов >> — арифметический сдвиг.
// Сложность: hard
package main_test

import "testing"

func unsignedShift(x int32, n uint) int32 {
	return x >> n
}

func TestUnsignedShift(t *testing.T) {
	if got := unsignedShift(16, 2); got != 4 {
		t.Errorf("unsignedShift(16, 2) = %d", got)
	}
	if got := unsignedShift(-1, 28); got != 15 {
		t.Errorf("unsignedShift(-1, 28) = %d, want 15", got)
	}
	if got := unsignedShift(-8, 1); got != 2147483644 {
		t.Errorf("unsignedShift(-8, 1) = %d, want 2147483644", got)
	}
}
