package mocking

import (
	"bytes"
	"io"
	"reflect"
	"testing"
	"time"
)

type SpySleeperWriter struct {
	Calls []string
}

var _ Sleeper = &SpySleeperWriter{}
var _ io.Writer = &SpySleeperWriter{}

const (
	Sleep = "sleep"
	Write = "write"
)

func (s *SpySleeperWriter) Sleep() {
	s.Calls = append(s.Calls, Sleep)
}

func (s *SpySleeperWriter) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, Write)
	return
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func TestCountdown(t *testing.T) {
	t.Run("output", func(t *testing.T) {
		// Arrange
		buffer := &bytes.Buffer{}
		spySleeper := &SpySleeperWriter{}

		// Act
		Countdown(buffer, spySleeper)

		// Assert
		have := buffer.String()
		want := "3\n2\n1\nGo!\n"

		if have != want {
			t.Errorf("have %q want %q", have, want)
		}
	})

	t.Run("order of operations", func(t *testing.T) {
		// Arrange
		spySleeper := &SpySleeperWriter{}

		// Act
		Countdown(spySleeper, spySleeper)

		// Assert
		have := spySleeper.Calls
		want := []string{
			Write,
			Sleep,
			Write,
			Sleep,
			Write,
			Sleep,
			Write,
		}

		if !reflect.DeepEqual(have, want) {
			t.Errorf("order of operations is wrong. have %v want %v", have, want)
		}
	})
}

func TestConfigurableSleeper(t *testing.T) {
	t.Run("configure, time slept", func(t *testing.T) {
		// Arrange
		sleepDuration := 10 * time.Second

		spyTime := &SpyTime{}
		sleeper := ConfigurableSleeper{duration: sleepDuration, sleep: spyTime.Sleep}

		// Act
		sleeper.Sleep()

		// Assert
		have := spyTime.durationSlept
		want := sleepDuration

		if have != want {
			t.Errorf("have %q want %q", have, want)
		}
	})
}
