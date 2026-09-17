// concurrent_x034: Значение из контекста
// Make the tests pass!
// I AM NOT DONE
//
// requestID достаёт идентификатор запроса из контекста.
// Тренирует: context.WithValue и ctx.Value.
// Сложность: easy
package main_test

import (
	"context"
	"testing"
)

type ctxKey string

const reqIDKey ctxKey = "request-id"

func requestID(ctx context.Context) string {
	v, _ := ctx.Value("request-id").(string)
	return v
}

func TestRequestID(t *testing.T) {
	ctx := context.WithValue(context.Background(), reqIDKey, "abc-1")
	if got := requestID(ctx); got != "abc-1" {
		t.Errorf("requestID = %q", got)
	}
}
