// functions_x097: Переиспользуемый буфер
// Make the tests pass!
// I AM NOT DONE
//
// scan передаёт в callback срез, который будет перезаписан после возврата.
// fields сохраняет эти срезы, и в результате оказываются испорченные данные.
// Тренирует: контракт «данные действительны только внутри вызова».
// Сложность: hard
package main_test

import "testing"

func scan(data string, emit func([]byte)) {
	buf := make([]byte, 0, 16)
	for i := 0; i < len(data); i++ {
		if data[i] == ',' {
			emit(buf)
			buf = buf[:0]
			continue
		}
		buf = append(buf, data[i])
	}
	emit(buf)
}

func fields(data string) [][]byte {
	var out [][]byte
	scan(data, func(b []byte) {
		out = append(out, b)
	})
	return out
}

func TestFields(t *testing.T) {
	got := fields("ab,cd,e")
	want := []string{"ab", "cd", "e"}
	if len(got) != len(want) {
		t.Fatalf("fields returned %d items, want %d", len(got), len(want))
	}
	for i := range want {
		if string(got[i]) != want[i] {
			t.Errorf("field %d = %q, want %q", i, got[i], want[i])
		}
	}
}
