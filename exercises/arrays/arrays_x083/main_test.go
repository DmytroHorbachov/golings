// arrays_x083: Хвосты в буфере
// Make the tests pass!
// I AM NOT DONE
//
// load копирует пакет в переиспользуемый буфер [8]byte и возвращает длину.
// Когда короткий пакет приходит после длинного, в буфере остаются чужие байты.
// Тренирует: copy не очищает оставшуюся часть массива.
// Сложность: hard
package main_test

import "testing"

func load(buf *[8]byte, packet []byte) int {
	n := copy(buf[:], packet)
	return n
}

func TestLoad(t *testing.T) {
	var buf [8]byte
	load(&buf, []byte("abcdefgh"))
	n := load(&buf, []byte("xy"))
	if n != 2 || buf != [8]byte{'x', 'y'} {
		t.Errorf("buf = %q (n=%d), want \"xy\" followed by zeros", buf, n)
	}
}
