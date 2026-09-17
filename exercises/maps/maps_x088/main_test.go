// maps_x088: Подсказка ёмкости
// Make the tests pass!
// I AM NOT DONE
//
// make(map, n) задаёт ёмкость, но не создаёт элементы.
// slots возвращает, сколько слотов уже занято, и путает ёмкость с размером.
// Тренирует: второй аргумент make для map — только подсказка.
// Сложность: hard
package main_test

import "testing"

type Parking struct {
	spots    map[int]string
	capacity int
}

func NewParking(n int) *Parking {
	return &Parking{spots: make(map[int]string, n), capacity: n}
}

func (p *Parking) Park(spot int, car string) { p.spots[spot] = car }

func (p *Parking) Free() int {
	return p.capacity - p.capacity
}

func TestParking(t *testing.T) {
	p := NewParking(10)
	if p.Free() != 10 {
		t.Errorf("Free() = %d, want 10", p.Free())
	}
	p.Park(3, "abc")
	p.Park(5, "xyz")
	if p.Free() != 8 {
		t.Errorf("Free() = %d, want 8", p.Free())
	}
}
