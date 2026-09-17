// variables_x043: Контрольная сумма
// Make the tests pass!
// I AM NOT DONE
//
// Функция checksum должна вернуть сумму всех байт строки.
// Для длинных строк сумма получается подозрительно маленькой.
// Тренирует: переполнение byte (uint8) при накоплении.
// Сложность: hard
package main_test

import "testing"

func checksum(s string) int {
	var sum byte
	for i := 0; i < len(s); i++ {
		sum += s[i]
	}
	return int(sum)
}

func TestChecksum(t *testing.T) {
	if got := checksum("abc"); got != 294 {
		t.Errorf("checksum(abc) = %d, want 294", got)
	}
	if got := checksum("zzzz"); got != 488 {
		t.Errorf("checksum(zzzz) = %d, want 488", got)
	}
}
