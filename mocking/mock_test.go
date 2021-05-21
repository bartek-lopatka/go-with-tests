package mocking

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

type CountdownOperationsSpy struct {
	Calls []string
}

func (s *CountdownOperationsSpy) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *CountdownOperationsSpy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

const write = "write"
const sleep = "sleep"

func TestCountdown(t *testing.T) {

	t.Run("prints 3 2 1 Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		spySleeper := &CountdownOperationsSpy{}

		Countdown(spySleeper, buffer)
		got := buffer.String()

		want := `3
2
1
Go!`
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
	t.Run("sleep before every print", func(t *testing.T) {
		spySleepPrinter := &CountdownOperationsSpy{}
		Countdown(spySleepPrinter, spySleepPrinter)

		want := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
			t.Errorf("wanted calls %v, but got %v", want, spySleepPrinter.Calls)
		}
	})
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 8 * time.Second

	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}

	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept for %v, but slept for %v", sleepTime, spyTime.durationSlept)
	}
}
