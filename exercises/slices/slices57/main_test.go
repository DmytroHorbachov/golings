// slices57
// Make the tests pass!

// I AM NOT DONE
//
// header возвращает первые 2 байта пакета как отдельный буфер, к которому потом
// дописывают данные. Дописывание портит тело пакета.
// Тренирует: ёмкость подсреза простирается до конца родительского массива.
// Сложность: hard
package main_test

import (
	"reflect"
	"testing"
)

func header(packet []byte) []byte {
	return packet[0:2]
}

func TestHeader(t *testing.T) {
	packet := []byte{0xCA, 0xFE, 1, 2, 3}
	h := header(packet)
	h = append(h, 0xFF)
	if !reflect.DeepEqual(packet, []byte{0xCA, 0xFE, 1, 2, 3}) {
		t.Errorf("packet corrupted: %v", packet)
	}
	if !reflect.DeepEqual(h, []byte{0xCA, 0xFE, 0xFF}) {
		t.Errorf("header = %v", h)
	}
}
