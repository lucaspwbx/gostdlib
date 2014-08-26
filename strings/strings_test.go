package strings

import (
	"strings"
	"testing"
	"unicode"
)

func TestFields(t *testing.T) {
	if fields := strings.Fields("  foo bar baz  "); len(fields) != 3 {
		t.Errorf("Error")
	}
}

func TestFieldsFunc(t *testing.T) {
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	if fields := strings.FieldsFunc(" foo1;bar2;baz3...", f); len(fields) != 3 {
		t.Errorf("Error")
	}
}

func TestContains(t *testing.T) {
	if r := strings.Contains("seafood", "foo"); r != true {
		t.Errorf("Error")
	}
}

func TestContainsAny(t *testing.T) {
	tests := []struct {
		word   string
		word2  string
		result bool
	}{
		{"test", "i", false},
		{"test", "t", true},
	}
	for _, v := range tests {
		if r := strings.Contains(v.word, v.word2); r != v.result {
			t.Errorf("Error")
		}
	}
}

func TestCount(t *testing.T) {
	if r := strings.Count("cheese", "e"); r != 3 {
		t.Errorf("Error")
	}
}

func TestEqualFold(t *testing.T) {
	if r := strings.EqualFold("Go", "go"); r != true {
		t.Errorf("Error")
	}
}

func TestIndex(t *testing.T) {
	if r := strings.Index("chicken", "ken"); r != 4 {
		t.Errorf("Error")
	}
}

func TestIndexAny(t *testing.T) {
	if r := strings.IndexAny("chicken", "aeiou"); r != 2 {
		t.Errorf("Error")
	}
}

func TestJoin(t *testing.T) {
	s := []string{"foo", "bar", "baz"}
	if r := strings.Join(s, ", "); r != "foo, bar, baz" {
		t.Errorf("Error")
	}
}

func TestRepeat(t *testing.T) {
	if r := strings.Repeat("na", 2); r != "nana" {
		t.Errorf("Error")
	}
}

func TestReplace(t *testing.T) {
	if r := strings.Replace("oink", "k", "ky", 1); r != "oinky" {
		t.Errorf("Error")
	}
	if r := strings.Replace("oink oink", "oink", "moo", -1); r != "moo moo" {
		t.Errorf("Error")
	}
}

func TestSplit(t *testing.T) {
	if r := strings.Split("a,b,c", ","); len(r) != 3 {
		t.Errorf("Error")
	}
}

func TestTitle(t *testing.T) {
	if r := strings.Title("test"); r != "Test" {
		t.Errorf("Error")
	}
}

func TestToTitle(t *testing.T) {
	if r := strings.ToTitle("loud noises"); r != "LOUD NOISES" {
		t.Errorf("Error")
	}
}

func TestTrimSpace(t *testing.T) {
	if r := strings.TrimSpace("\t\n a lone gopher"); r != "a lone gopher" {
		t.Errorf("Error")
	}
}

func TestToUpper(t *testing.T) {
	if r := strings.ToUpper("Gopher"); r != "GOPHER" {
		t.Errorf("Error")
	}
}

func TestToLower(t *testing.T) {
	if r := strings.ToLower("GOPHER"); r != "gopher" {
		t.Errorf("Error")
	}
}
