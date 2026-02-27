package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	finalWord      = "Go!"
	countdownStart = 3
)

type Sleeper interface {
	Sleep()
}

// type SpySleeper struct {
// 	Calls int
// }

// func (s *SpySleeper) Sleep() {
// 	s.Calls++
// }

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

type SpyCountdownOperations struct {
	Calls []string
}

// Sleep implements Sleeper interface - records each sleep call
func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

// Write implements io.Writer interface - records each write operation
func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

const (
	write = "write"
	sleep = "sleep"
)

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
		// time.Sleep(1 * time.Second)
	}

	fmt.Fprintf(out, finalWord)
}

func main() {
	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
}
