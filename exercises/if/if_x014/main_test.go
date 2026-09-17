// if_x014: FizzBuzz
// Make the tests pass!
// I AM NOT DONE
//
// fizzbuzz должна вернуть "FizzBuzz" для кратных 15, "Fizz" для 3, "Buzz" для 5.
// Тренирует: порядок проверок, когда условия перекрываются.
// Сложность: medium
package main_test

import (
	"strconv"
	"testing"
)

func fizzbuzz(n int) string {
	if n%3 == 0 {
		return "Fizz"
	} else if n%15 == 0 {
		return "FizzBuzz"
	} else if n%5 == 0 {
		return "Buzz"
	}
	return strconv.Itoa(n)
}

func TestFizzbuzz(t *testing.T) {
	cases := map[int]string{3: "Fizz", 5: "Buzz", 15: "FizzBuzz", 30: "FizzBuzz", 7: "7"}
	for in, want := range cases {
		if got := fizzbuzz(in); got != want {
			t.Errorf("fizzbuzz(%d) = %s, want %s", in, got, want)
		}
	}
}
