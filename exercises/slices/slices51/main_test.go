// slices51
// Make the tests pass!

// I AM NOT DONE
//
// parentDir возвращает все части пути, кроме последней, склеенные через "/".
// Тренирует: strings.Split и срез результата.
// Сложность: easy
package main_test

import (
	"strings"
	"testing"
)

func parentDir(path string) string {
	parts := strings.Split(path, "/")
	return strings.Join(parts[1:], "/")
}

func TestParentDir(t *testing.T) {
	if got := parentDir("usr/local/bin"); got != "usr/local" {
		t.Errorf("parentDir = %q", got)
	}
}
