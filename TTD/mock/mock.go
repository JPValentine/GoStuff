package main

import(
	"io"
	"os"
	"fmt"
	"time"
)

const FinalWord = "Go!"
const CountdownStart = 3
const write = "write"
const sleep = "sleep"

type Sleeper interface{
	Sleep()
}

type SpySleeper struct{
	Calls int
}

type ConfigSleeper struct{
	duration time.Duration
	sleep func(time.Duration)
}

type SpyTime struct{
	durationSlept time.Duration
}

type SpyCountdownOperations struct{
	Calls []string
}

func (c *ConfigSleeper) Sleep(){
	c.sleep(c.duration)
}

func (s *SpyTime) Sleep(duration time.Duration){
	s.durationSlept = duration
}

func (s *SpyCountdownOperations) Sleep(){
	s.Calls = append(s.Calls, sleep)
}

func ( s *SpyCountdownOperations) Write(p []byte) (n int, err error){
	s.Calls = append(s.Calls, write)
	return
}

func (s *SpySleeper) Sleep(){
	s.Calls++
}

func Countdown(out io.Writer, s Sleeper){
	for i := CountdownStart; i>0; i--{
		fmt.Fprintln(out, i)
		s.Sleep()
	}
	fmt.Fprint(out, FinalWord)
}

func main(){
	sleeper := &ConfigSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}
