package main

import(
	"bytes"
	"testing"
)

func testGreet(t *testing.T){
	buffer := bytes.Buffer{}
	Greet(&buffer, "")

	got := buffer.String()
	want := "Hello, JP"

	if got != want{
		t.Errorf("got %q want %q",got, want)
	}//if
}//testGreet
