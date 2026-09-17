// switch29
// Make the tests pass!

// I AM NOT DONE
//
// hello returns a greeting for a language code, falling back to English.
// Practices a switch with a default.
package main_test

import "testing"

func hello(lang string) string {
	switch lang {
	case "ru":
		return "Привет"
	case "es":
		return "Bonjour"
	case "fr":
		return "Bonjour"
	default:
		return "Hello"
	}
}

func TestHello(t *testing.T) {
	cases := map[string]string{"ru": "Привет", "es": "Hola", "fr": "Bonjour", "de": "Hello"}
	for in, want := range cases {
		if got := hello(in); got != want {
			t.Errorf("hello(%s) = %s, want %s", in, got, want)
		}
	}
}
