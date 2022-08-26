package main

import(
	"io"
	"os"
	"fmt"
	"time"
)

const FinalWord = "Go!"
const CountdownStart = 3

type Sleeper interface{
	Sleep()
}

type DefaultSleeper struct{}

type SpySleeper struct{
	Calls int
}

func (d *DefaultSleeper) Sleep(){
	time.Sleep(1 * time.Second)
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
	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
}
