// primitive_types39
// Make the tests pass!

// I AM NOT DONE
//
// isRussian checks that a string holds lowercase russian letters only.
// Words with the letter "ё" are rejected right now.
// 'ё' sits outside the contiguous 'а'..'я' range in Unicode.
package main_test

import "testing"

func isRussian(s string) bool {
	for _, r := range s {
		if r < 'а' || r > 'я' {
			return false
		}
	}
	return s != ""
}

func TestIsRussian(t *testing.T) {
	cases := map[string]bool{"ёлка": true, "мир": true, "ещё": true, "hello": false, "Мир": false, "": false}
	for in, want := range cases {
		if got := isRussian(in); got != want {
			t.Errorf("isRussian(%q) = %v, want %v", in, got, want)
		}
	}
}
