// variables22
// Make the tests pass!

// I AM NOT DONE
//
// Дни недели заданы через iota, а метод String должен возвращать их названия.
// Сейчас нумерация и таблица имён не совпадают.
// Тренирует: iota, именованные типы и методы у них.
// Сложность: easy
package main_test

import "testing"

type Weekday int

const (
	Monday Weekday = iota + 1
	Tuesday
	Wednesday
)

var weekdayNames = [...]string{"Mon", "Tue", "Wed"}

func (d Weekday) String() string {
	return weekdayNames[d]
}

func TestWeekdayString(t *testing.T) {
	cases := map[Weekday]string{Monday: "Mon", Tuesday: "Tue", Wednesday: "Wed"}
	for d, want := range cases {
		if got := d.String(); got != want {
			t.Errorf("Weekday(%d).String() = %s, want %s", int(d), got, want)
		}
	}
}
