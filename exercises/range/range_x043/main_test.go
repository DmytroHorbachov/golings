// range_x043: Группировка структур
// Make the tests pass!
// I AM NOT DONE
//
// byTeam группирует имена игроков по команде; порядок внутри команды — как во входе.
// Тренирует: range по срезу структур и map со срезами.
// Сложность: medium
package main_test

import (
	"reflect"
	"testing"
)

type Player struct{ Name, Team string }

func byTeam(ps []Player) map[string][]string {
	out := map[string][]string{}
	for _, p := range ps {
		team := p.Name
		out[team] = []string{p.Name}
	}
	return out
}

func TestByTeam(t *testing.T) {
	got := byTeam([]Player{{"ann", "red"}, {"bob", "blue"}, {"cid", "red"}})
	if !reflect.DeepEqual(got, map[string][]string{"red": {"ann", "cid"}, "blue": {"bob"}}) {
		t.Errorf("byTeam = %v", got)
	}
}
