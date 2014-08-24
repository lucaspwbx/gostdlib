package bytes

import (
	"bytes"
	"testing"
)

func TestContains(t *testing.T) {
	a := []byte("go")
	b := []byte("golang")
	if c := bytes.Contains(b, a); c != true {
		t.Errorf("Should be true")
	}
}

func TestCount(t *testing.T) {
	sep := []byte("go")
	a := []byte("golang rules, golang rules")
	if num := bytes.Count(a, sep); num != 2 {
		t.Errorf("Should be 2")
	}
}

func TestHasPrefix(t *testing.T) {
	prefix := []byte("go")
	s := []byte("golang")
	if c := bytes.HasPrefix(s, prefix); c != true {
		t.Errorf("Should be true")
	}
}

func TestHasSuffix(t *testing.T) {
	suffix := []byte("go")
	s := []byte("lang go")
	if c := bytes.HasSuffix(s, suffix); c != true {
		t.Errorf("should be true")
	}
}

func TestIndex(t *testing.T) {
	sep := []byte("go")
	s := []byte("best programming language: golang")
	if i := bytes.Index(s, sep); i < 0 {
		t.Errorf("Index should be > -1")
	}
}

func TestIndexAny(t *testing.T) {
	chars := "go"
	s := []byte("best programming language: golang")
	if i := bytes.IndexAny(s, chars); i < 0 {
		t.Errorf("Index should be > -1")
	}
}

func TestIndexByte(t *testing.T) {
	b := byte('g')
	s := []byte("golang")
	if i := bytes.IndexByte(s, b); i == -1 {
		t.Errorf("Shoud be > 0")
	}
}

func TestIndexFunc(t *testing.T) {
	g := rune('g')
	f := func(r rune) bool { return r == g }
	s := []byte("golang")
	if i := bytes.IndexFunc(s, f); i == -1 {
		t.Errorf("Should be >= 0")
	}
}

func TestIndexrune(t *testing.T) {
	r := rune('g')
	s := []byte("golang")
	if i := bytes.IndexRune(s, r); i == -1 {
		t.Errorf("Should be >= 0")
	}
}

func TestLastIndex(t *testing.T) {
	sep := []byte("go")
	s := []byte("go go")
	if i := bytes.LastIndex(s, sep); i != 3 {
		t.Errorf("Should be 3, got %d", i)
	}
}
