// structs_x036: Текучий строитель
// Make the tests pass!
// I AM NOT DONE
//
// QueryBuilder собирает SQL: Table, Where (может вызываться несколько раз), Build.
// Тренирует: методы-указатели, возвращающие получатель.
// Сложность: medium
package main_test

import (
	"strings"
	"testing"
)

type QueryBuilder struct {
	table string
	conds []string
}

func (q *QueryBuilder) Table(t string) *QueryBuilder {
	q.table = t
	return q
}

func (q *QueryBuilder) Where(c string) *QueryBuilder {
	q.conds = []string{c}
	return q
}

func (q *QueryBuilder) Build() string {
	s := "SELECT * FROM " + q.table
	if len(q.conds) > 0 {
		s += " WHERE " + strings.Join(q.conds, " OR ")
	}
	return s
}

func TestQueryBuilder(t *testing.T) {
	got := (&QueryBuilder{}).Table("users").Where("age > 18").Where("active").Build()
	if got != "SELECT * FROM users WHERE age > 18 AND active" {
		t.Errorf("Build = %q", got)
	}
}
