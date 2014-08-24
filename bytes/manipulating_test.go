package bytes

import (
	"bytes"
	"testing"
)

func TestReplace(t *testing.T) {
}

func TestRunes(t *testing.T) {
	test := []byte("go")
	runes := bytes.Runes(test)
	expected := []rune{103, 111}
	for k, v := range runes {
		if v != expected[k] {
			t.Errorf("Error")
		}
	}
}

func TestRepeat(t *testing.T) {
	n := 2
	test := []byte("go")
	final := bytes.Repeat(test, n)
	if c := bytes.Equal(final, []byte("gogo")); c != true {
		t.Errorf("Error")
	}
}
