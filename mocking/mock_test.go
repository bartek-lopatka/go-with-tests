package mocking

import (
	"bytes"
	"testing"
)

type CountdownOperationsSpy struct {
	Calls []string
}

func (s *CountdownOperationsSpy) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *CountdownOperationsSpy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
}

const write = "write"
const sleep = "sleep"

func TestCountdown(t *testing.T) {
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

	if spySleeper.Calls != 4 {
		t.Errorf("Not enough calls to sleeper, wanted 4, got %d", spySleeper.Calls)
	}
}
