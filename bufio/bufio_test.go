package bufio

import (
	"bufio"
	"os"
	"strings"
	"testing"
)

func TestNewReaderSize(t *testing.T) {
	s := []byte("hello")
	str := strings.NewReader(string(s))
	buffer := bufio.NewReaderSize(str, 10)
	if buffer == nil {
		t.Errorf("Error")
	}
}

func TestPeek(t *testing.T) {
	s := []byte("helloworld")
	str := strings.NewReader(string(s))
	rb := bufio.NewReader(str)
	peeked, _ := rb.Peek(5)
	if string(peeked) != "hello" {
		t.Errorf("Error")
	}
}

func TestReadString(t *testing.T) {
	s := []byte("helloworld")
	str := strings.NewReader(string(s))
	rb := bufio.NewReader(str)
	line, _ := rb.ReadString(' ')
	if line != "helloworld" {
		t.Errorf("Error")
	}
}

func TestReadRune(t *testing.T) {
	s := []byte("helloworld")
	str := strings.NewReader(string(s))
	rb := bufio.NewReader(str)
	r, _, _ := rb.ReadRune()
	if r != 'h' {
		t.Errorf("Error")
	}
}

func TestReadByte(t *testing.T) {
	s := []byte("helloworld")
	str := strings.NewReader(string(s))
	rb := bufio.NewReader(str)
	b, _ := rb.ReadByte()
	if string(b) != "h" {
		t.Errorf("Error")
	}
}

func TestLineRead(t *testing.T) {
	s := []byte("helloworld")
	str := strings.NewReader(string(s))
	rb := bufio.NewReader(str)
	line, _, _ := rb.ReadLine()
	if string(line) != "helloworld" {
		t.Errorf("Error")
	}
}

func TestWriteByte(t *testing.T) {
	wb := bufio.NewWriter(os.Stdout)
	if err := wb.WriteByte('G'); err != nil {
		t.Errorf("Error")
	}
	//defer wb.Flush()
}

func TestWriteRune(t *testing.T) {
	wb := bufio.NewWriter(os.Stdout)
	if written, _ := wb.WriteRune(rune('t')); written != 1 {
		t.Errorf("Error")
	}
}

func TestWriteString(t *testing.T) {
	wb := bufio.NewWriter(os.Stdout)
	if written, _ := wb.WriteString("lucas"); written != 5 {
		t.Errorf("Error")
	}
}
