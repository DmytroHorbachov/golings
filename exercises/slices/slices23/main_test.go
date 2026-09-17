// slices23
// Make the tests pass!

// I AM NOT DONE
//
// render возвращает содержимое буфера, который потом используется повторно.
// Возвращённые данные перезаписываются следующим вызовом.
// Тренирует: bytes.Buffer.Bytes() действителен только до следующего изменения буфера.
// Сложность: hard
package main_test

import (
	"bytes"
	"testing"
)

type Renderer struct{ buf bytes.Buffer }

func (r *Renderer) render(name string) []byte {
	r.buf.Reset()
	r.buf.WriteString("hello, ")
	r.buf.WriteString(name)
	return r.buf.Bytes()
}

func TestRender(t *testing.T) {
	var r Renderer
	a := r.render("ann")
	b := r.render("bob")
	if string(a) != "hello, ann" || string(b) != "hello, bob" {
		t.Errorf("a = %q, b = %q", a, b)
	}
}
