// if_x058: Статус батареи
// Make the tests pass!
// I AM NOT DONE
//
// batteryStatus: при зарядке — "charging" (или "full" при 100%),
// без зарядки: до 15% — "low", иначе "ok".
// Тренирует: вложенные if по нескольким признакам.
// Сложность: medium
package main_test

import "testing"

func batteryStatus(level int, charging bool) string {
	if level < 15 {
		return "low"
	}
	if charging {
		return "charging"
	}
	return "ok"
}

func TestBatteryStatus(t *testing.T) {
	cases := []struct {
		level    int
		charging bool
		want     string
	}{{100, true, "full"}, {10, true, "charging"}, {50, true, "charging"}, {10, false, "low"}, {15, false, "ok"}}
	for _, c := range cases {
		if got := batteryStatus(c.level, c.charging); got != c.want {
			t.Errorf("batteryStatus(%d, %v) = %s, want %s", c.level, c.charging, got, c.want)
		}
	}
}
