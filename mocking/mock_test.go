package mocking

import (
	"bytes"
	"testing"
)

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	spySleeper := &SpySleeper{}

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
