package main

import(
	"testing"
	"bytes"
)

func TestCountdown(t *testing.T){
	buffer := &bytes.Buffer{}
	s := &SpySleeper{}	

	Countdown(buffer,s)

	got := buffer.String()
	want:= `3
2
1
Go!`

	if got != want{
		t.Errorf("got %q want %q", got, want)
	}
	if s.Calls != 3{
		t.Errorf("not enough calls to sleeper, want 3 got %d", s.Calls)
	}
}
