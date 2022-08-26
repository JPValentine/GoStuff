package main

import(
	"testing"
	"bytes"
	"reflect"
	"time"
)

func TestCountdown(t *testing.T){
	t.Run("print check", func(t *testing.T){
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
	})
	t.Run("sleep before every print", func(t *testing.T){
		spySleepPrinter := &SpyCountdownOperations{}
		Countdown(spySleepPrinter, spySleepPrinter)

		want := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}
		if !reflect.DeepEqual(want, spySleepPrinter.Calls){
			t.Errorf("wanted calls: %v got calls: %v", want, spySleepPrinter.Calls)
		}
	})
}

func TestConfigSleeper(t *testing.T) {
	sleepTime := 5 * time.Second

	spyTime := &SpyTime{}
	sleeper := ConfigSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime{
		t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
	}
}
