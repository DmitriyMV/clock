package clock_test

import (
	"sync"
	"testing"
	"time"

	"github.com/LopatkinEvgeniy/clock"
)

func TestFakeClockNow(t *testing.T) {
	t.Run("use NewFakeClock constructor", func(t *testing.T) {
		c := clock.NewFakeClock()
		expected := time.Time{}

		for i := 0; i < 100; i++ {
			c.Add(time.Hour)
			expected = expected.Add(time.Hour)

			now := c.Now()
			if now != expected {
				t.Fatalf("unexpected now result, expected: %s, actual: %s", expected, now)
			}
		}
	})

	t.Run("use NewFakeClockAt constructor", func(t *testing.T) {
		initialTime := time.Now()
		c := clock.NewFakeClockAt(initialTime)
		expected := initialTime

		for i := 0; i < 100; i++ {
			c.Add(time.Hour)
			expected = expected.Add(time.Hour)

			now := c.Now()
			if now != expected {
				t.Fatalf("unexpected now result, expected: %s, actual: %s", expected, now)
			}
		}
	})
}

func TestFakeClockNowStress(t *testing.T) {
	c := clock.NewFakeClock()

	wg := sync.WaitGroup{}
	for i := 0; i < 100000; i++ {
		wg.Add(1)
		go func() {
			c.Add(time.Minute)
			wg.Done()
		}()
	}
	wg.Wait()

	expected := (time.Time{}).Add(100000 * time.Minute)
	now := c.Now()
	if now != expected {
		t.Fatalf("unexpected now result, expected: %s, actual: %s", expected, now)
	}
}