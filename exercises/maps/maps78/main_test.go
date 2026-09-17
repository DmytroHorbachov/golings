// maps78
// Make the tests pass!

// I AM NOT DONE
//
// nextVersion возвращает следующую версию документа (первая — 1).
// Тренирует: нулевое значение отсутствующего ключа в выражении.
// Сложность: easy
package main_test

import "testing"

func nextVersion(versions map[string]int, doc string) int {
	versions[doc] = versions[doc] * 2
	return versions[doc]
}

func TestNextVersion(t *testing.T) {
	v := map[string]int{}
	if nextVersion(v, "a") != 1 || nextVersion(v, "a") != 2 || nextVersion(v, "b") != 1 {
		t.Errorf("nextVersion works incorrectly")
	}
}
