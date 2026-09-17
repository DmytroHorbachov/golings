// switch39
// Make the tests pass!

// I AM NOT DONE
//
// sumValid суммирует числа из строк, пропуская комментарии "#...".
// break в switch не пропускает итерацию, и комментарии ломают сумму.
// Тренирует: break выходит из switch, а continue переходит к следующей итерации цикла.
// Сложность: hard
package main_test

import (
	"strconv"
	"strings"
	"testing"
)

func sumValid(lines []string) (int, int) {
	sum, bad := 0, 0
	for _, l := range lines {
		switch {
		case strings.HasPrefix(l, "#"):
			break
		case l == "":
			continue
		}
		n, err := strconv.Atoi(l)
		if err != nil {
			bad++
			continue
		}
		sum += n
	}
	return sum, bad
}

func TestSumValid(t *testing.T) {
	sum, bad := sumValid([]string{"1", "# comment", "", "2", "x"})
	if sum != 3 || bad != 1 {
		t.Errorf("sumValid = %d, %d; want 3, 1", sum, bad)
	}
}
