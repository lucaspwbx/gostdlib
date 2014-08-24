package bytes

import (
	"bytes"
	"testing"
)

func TestCompare(t *testing.T) {
	a := []byte("rails")
	b := []byte("ruby")
	if c := bytes.Compare(a, b); c != -1 {
		t.Errorf("Result should be -1")
	}

	if c := bytes.Compare(b, a); c != 1 {
		t.Errorf("Result should be 1")
	}

	a = []byte("golang")
	b = []byte("golang")
	if c := bytes.Compare(a, b); c != 0 {
		t.Errorf("Result should be 0")
	}
}

func TestEqual(t *testing.T) {
	a := []byte("golang")
	b := []byte("golang")

	if c := bytes.Equal(a, b); c != true {
		t.Errorf("Should be true")
	}

	b = []byte("ruby")

	if c := bytes.Equal(a, b); c != false {
		t.Errorf("Should be false")
	}
}

func TestEqualFold(t *testing.T) {
	a := []byte("golang")
	b := []byte("goLanG")

	if c := bytes.EqualFold(a, b); c != true {
		t.Errorf("Should be true")
	}
}
