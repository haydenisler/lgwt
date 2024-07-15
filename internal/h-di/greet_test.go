package hdi

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Hayden")

	got := buffer.String()
	want := "Hello, Hayden"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
