// variables_x092: Блок var
// Make the tests pass!
// I AM NOT DONE
//
// Переменные пакета описывают адрес базы данных.
// Функция dsn должна собрать строку вида "db.local:5432/shop".
// Тренирует: групповое объявление переменных через var ( ... ).
// Сложность: medium
package main_test

import (
	"strconv"
	"testing"
)

var (
	dbHost = "db.local"
	dbPort = 3306
	dbName = "shop"
)

func dsn() string {
	return dbHost + ":" + dbName + "/" + strconv.Itoa(dbPort)
}

func TestDSN(t *testing.T) {
	if got := dsn(); got != "db.local:5432/shop" {
		t.Errorf("dsn() = %q, want %q", got, "db.local:5432/shop")
	}
}
