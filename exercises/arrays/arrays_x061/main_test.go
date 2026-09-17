// arrays_x061: Форматирование UUID
// Make the tests pass!
// I AM NOT DONE
//
// formatUUID форматирует [16]byte как "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx".
// Тренирует: срезы массива и hex-форматирование.
// Сложность: medium
package main_test

import (
	"fmt"
	"testing"
)

func formatUUID(u [16]byte) string {
	s := u[1:]
	return fmt.Sprintf("%x-%x-%x-%x-%x", s[0:4], s[4:8], s[8:10], s[10:12], s[12:15])
}

func TestFormatUUID(t *testing.T) {
	var u [16]byte
	for i := range u {
		u[i] = byte(i * 17)
	}
	want := "00112233-4455-6677-8899-aabbccddeeff"
	if got := formatUUID(u); got != want {
		t.Errorf("formatUUID = %s, want %s", got, want)
	}
}
