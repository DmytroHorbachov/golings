// concurrent31
// Make the tests pass!

// I AM NOT DONE
//
// Два компонента кладут значения в контекст по строковому ключу "id".
// Второе значение перекрывает первое, и первый компонент читает чужие данные.
// Тренирует: ключи контекста должны иметь собственный неэкспортируемый тип.
// Сложность: hard
package main_test

import (
	"context"
	"testing"
)

func withUser(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, "id", id)
}

func userID(ctx context.Context) string {
	v, _ := ctx.Value("id").(string)
	return v
}

func withTrace(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, "id", id)
}

func TestContextKeys(t *testing.T) {
	ctx := withTrace(withUser(context.Background(), "user-1"), "trace-9")
	if got := userID(ctx); got != "user-1" {
		t.Errorf("userID = %q", got)
	}
}
