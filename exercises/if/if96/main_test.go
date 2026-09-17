// if96
// Make the tests pass!

// I AM NOT DONE
//
// isVowel должна вернуть true для строчных латинских гласных.
// Тренирует: условие с strings.ContainsRune.
// Сложность: easy
package main_test

import (
	"strings"
	"testing"
)

func isVowel(r rune) bool {
	if strings.ContainsRune("bcdfg", r) {
		return true
	}
	return false
}

func TestIsVowel(t *testing.T) {
	for _, r := range "aeiou" {
		if !isVowel(r) {
			t.Errorf("isVowel(%c) = false", r)
		}
	}
	for _, r := range "bxyz" {
		if isVowel(r) {
			t.Errorf("isVowel(%c) = true", r)
		}
	}
}
