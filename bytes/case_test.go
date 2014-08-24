package bytes

import (
	"bytes"
	"testing"
)

func TestTitle(t *testing.T) {
	word := []byte("nachtmysticum")
	title := bytes.Title(word)
	if c := bytes.Equal(title, []byte("Nachtmysticum")); c != true {
		t.Errorf("Error")
	}
}

func TestToTitle(t *testing.T) {
	word := []byte("golang")
	title := bytes.ToTitle(word)
	if c := bytes.Equal(title, []byte("GOLANG")); c != true {
		t.Errorf("Error")
	}
}

func TestToLower(t *testing.T) {
	word := []byte("GOLANG")
	lower := bytes.ToLower(word)
	if c := bytes.Equal(lower, []byte("golang")); c != true {
		t.Errorf("Error")
	}
}

func TestToUpper(t *testing.T) {
	word := []byte("golang")
	upper := bytes.ToUpper(word)
	if c := bytes.Equal(upper, []byte("GOLANG")); c != true {
		t.Errorf("Error")
	}
}
