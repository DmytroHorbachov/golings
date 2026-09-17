// primitive_types_x059: Порядок байтов
// Make the tests pass!
// I AM NOT DONE
//
// encode записывает uint32 в 4 байта в порядке big-endian, decode читает обратно.
// Тренирует: пакет encoding/binary.
// Сложность: medium
package main_test

import (
	"encoding/binary"
	"testing"
)

func encode(v uint32) []byte {
	b := make([]byte, 4)
	binary.LittleEndian.PutUint32(b, v)
	return b
}

func decode(b []byte) uint32 {
	return binary.LittleEndian.Uint32(b)
}

func TestEncode(t *testing.T) {
	b := encode(0x01020304)
	if b[0] != 1 || b[3] != 4 {
		t.Errorf("encode = %v, want [1 2 3 4]", b)
	}
	if got := decode([]byte{0, 0, 1, 0}); got != 256 {
		t.Errorf("decode([0 0 1 0]) = %d, want 256", got)
	}
}
