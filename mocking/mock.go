package mocking

import (
	"fmt"
	"io"
	"os"
	"time"
)

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

type Sleeper interface {
	Sleep()
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(duration time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

const lastWord = "Go!"
const countdownStart = 3

func main() {
	sleeper := &DefaultSleeper{}
	Countdown(sleeper, os.Stdout)
}

func Countdown(sleeper Sleeper, out io.Writer) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}
	sleeper.Sleep()
	fmt.Fprint(out, lastWord)
}
