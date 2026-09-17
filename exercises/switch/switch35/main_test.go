// switch35
// Make the tests pass!

// I AM NOT DONE
//
// describe должна вернуть "nil" и для nil-интерфейса, и для nil-указателя *User.
// Тренирует: nil-указатель в интерфейсе попадает в ветку своего типа, а не в case nil.
// Сложность: hard
package main_test

import "testing"

type User struct{ Name string }

func describe(v interface{}) string {
	switch u := v.(type) {
	case nil:
		return "nil"
	case *User:
		return "user " + u.Name
	}
	return "other"
}

func TestDescribe(t *testing.T) {
	var none *User
	cases := []struct {
		in   interface{}
		want string
	}{{nil, "nil"}, {none, "nil"}, {&User{"ann"}, "user ann"}, {5, "other"}}
	for _, c := range cases {
		if got := describe(c.in); got != c.want {
			t.Errorf("describe(%#v) = %s, want %s", c.in, got, c.want)
		}
	}
}
