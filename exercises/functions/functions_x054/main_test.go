// functions_x054: Конвейер функций
// Make the tests pass!
// I AM NOT DONE
//
// pipeline(fs...) должна вернуть функцию, применяющую fs по порядку: сначала fs[0].
// Тренирует: срезы функций и накопление результата.
// Сложность: medium
package main_test

import "testing"

func pipeline(fs ...func(int) int) func(int) int {
	return func(x int) int {
		result := 0
		for _, f := range fs {
			result = f(x)
		}
		return result
	}
}

func TestPipeline(t *testing.T) {
	inc := func(x int) int { return x + 1 }
	sq := func(x int) int { return x * x }
	if got := pipeline(inc, sq)(3); got != 16 {
		t.Errorf("pipeline(inc, sq)(3) = %d, want 16", got)
	}
	if got := pipeline()(7); got != 7 {
		t.Errorf("pipeline()(7) = %d, want 7", got)
	}
}
