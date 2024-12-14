package main

import "testing"

var testString = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aliquam sit amet nunc eleifend, efficitur risus in, viverra eros.\n" +
	"Sed semper turpis nec leo iaculis porttitor. Etiam elit augue, finibus a eros et, ultrices viverra purus.\n" +
	"Vestibulum sit amet nulla interdum, tincidunt lacus eget, tristique eros. Donec et nisi eleifend tellus scelerisque aliquet nec quis arcu.\n" +
	"Nam tristique posuere fringilla. Donec quis lectus id sem tempus fermentum. Aenean consectetur orci a mi imperdiet, nec efficitur lorem sollicitudin😀.\n"

var test = struct {
	str        string
	lines      int
	words      int
	bts        int
	characters int
}{
	str:        testString,
	lines:      4,
	words:      76,
	bts:        522,
	characters: 519,
}

func TestLineCount(t *testing.T) {
	got := lineCount([]byte(test.str))
	want := test.lines
	if got != want {
		t.Fatalf("Wanted %d, got %d", want, got)
	}
}

func TestWordCount(t *testing.T) {
	got := wordCount([]byte(test.str))
	want := test.words
	if got != want {
		t.Fatalf("Wanted %d, got %d", want, got)
	}
}

func TestByteCount(t *testing.T) {
	got := byteCount([]byte(test.str))
	want := test.bts
	if got != want {
		t.Fatalf("Wanted %d, got %d", want, got)
	}
}

func TestCharacterCount(t *testing.T) {
	got := characterCount([]byte(test.str))
	want := test.characters
	if got != want {
		t.Fatalf("Wanted %d, got %d", want, got)
	}
}
