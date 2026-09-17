// algorithms_x068: Implement Trie (префиксное дерево)
// Make the tests pass!
// I AM NOT DONE
//
// Паттерн: префиксное дерево. Реализуйте Insert, Search (точное слово)
// и StartsWith (есть ли слово с таким префиксом). Слова из строчных латинских букв.
// Сложность: medium. Ожидаемая асимптотика: O(len) на операцию, O(общая длина) по памяти
package main_test

import "testing"

type Trie struct {
	children map[byte]*Trie
	isWord   bool
}

func NewTrie() *Trie { return &Trie{children: map[byte]*Trie{}} }

func (t *Trie) Insert(w string) {
}

func (t *Trie) node(prefix string) *Trie {
	cur := t
	for i := 0; i < len(prefix); i++ {
		next, ok := cur.children[prefix[i]]
		if !ok {
			return nil
		}
		cur = next
	}
	return cur
}

func (t *Trie) Search(w string) bool {
	return false
}

func (t *Trie) StartsWith(p string) bool {
	return false
}

func TestTrie(t *testing.T) {
	tr := NewTrie()
	tr.Insert("apple")
	if !tr.Search("apple") || tr.Search("app") || !tr.StartsWith("app") {
		t.Errorf("basic trie operations failed")
	}
	tr.Insert("app")
	if !tr.Search("app") || tr.StartsWith("banana") {
		t.Errorf("after inserting app")
	}
	tr.Insert("")
	if !tr.Search("") || !tr.StartsWith("") {
		t.Errorf("empty word handling failed")
	}
}
