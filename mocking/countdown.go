package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

// 1. Constants

const (
	finalWord      = "Go!"
	countdownStart = 3
	write          = "write"
	sleep          = "sleep"
)

// 2. Type definitions (interfaces first, then structs)

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct{}

type SpyCountdownOperations struct {
	Calls []string
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

type SpyTime struct {
	durationSlept time.Duration
}

// 3. Methods

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
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

func (s *SpyTime) SetDurationSlept(duration time.Duration) {
	s.durationSlept = duration
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

// 4, Functions

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
		// time.Sleep(1 * time.Second)
	}

	fmt.Fprintf(out, finalWord)
}

// 5. main function

func main() {
	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
}
