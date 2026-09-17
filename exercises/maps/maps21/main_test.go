// maps21
// Make the tests pass!

// I AM NOT DONE
//
// snapshot сохраняет состояние настроек, чтобы потом можно было откатиться.
// После изменения настроек «снимок» меняется тоже.
// Тренирует: присваивание map копирует ссылку, а не содержимое.
// Сложность: hard
package main_test

import "testing"

func snapshot(settings map[string]string) map[string]string {
	saved := settings
	return saved
}

func TestSnapshot(t *testing.T) {
	settings := map[string]string{"mode": "safe"}
	saved := snapshot(settings)
	settings["mode"] = "turbo"
	if saved["mode"] != "safe" {
		t.Errorf("snapshot changed: %v", saved)
	}
}
