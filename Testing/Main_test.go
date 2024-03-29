package main

import (
	"bytes"
	"testing"
)

func TestAdd_ok(t *testing.T) {
	got := Add(4, 5)
	if got != 9 {
		t.Errorf("Add(4,5) = %d; want 9", got)
	}
}

func TestAdd_nok(t *testing.T) {
	got := Add(3, 5)
	if got != 9 {
		t.Errorf("Add(3,5) = %d; want 9", got)
	}
}

func BenchmarkRandInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(879, 898)
	}
}

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	spySleeper := &SpySleeper{}

	Countdown(buffer, spySleeper)

	got := buffer.String()
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

	if spySleeper.Calls != 4 {
		t.Errorf("not enough calls to sleeper, want 4 got %d", spySleeper.Calls)
	}
}
