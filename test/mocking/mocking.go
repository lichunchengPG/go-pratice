package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord	= "Go!"
const write = "write"
const sleep = "sleep"

type Sleeper interface {
	Sleep()
}
type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep()  {
	s.Calls++
}

type DefaultSleeper struct {

}

func (d *DefaultSleeper) Sleep()  {
	time.Sleep(1*time.Second)
}

type CountdownOperationsSpy struct {
	Calls []string
}

func (c *CountdownOperationsSpy) Sleep()  {
	c.Calls = append(c.Calls, sleep)
}

func (c *CountdownOperationsSpy) Write(p []byte) (n int, err error) {
	c.Calls = append(c.Calls, write)
	return
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep  func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep()  {
	c.sleep(c.duration)
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration)  {
	s.durationSlept = duration
}

func Countdown(out io.Writer, sleeper Sleeper)  {
	for i := 3; i>0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}
	sleeper.Sleep()
	fmt.Fprint(out, finalWord)
}

func main()  {
	sleep := &ConfigurableSleeper{1*time.Second, time.Sleep}
	Countdown(os.Stdout, sleep)
}
t