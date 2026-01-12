package clock_test

import (
	"testing"
	"time"

	"github.com/renq/interlocutr/internal/infrastructure/clock"

	"github.com/stretchr/testify/assert"
)

func TestClock(t *testing.T) {
	t.Parallel()

	t.Run("returns default time by default", func(t *testing.T) {
		t.Parallel()
		clock := clock.NewClock()
		now := clock.Now()

		assert.WithinDuration(t, now, clock.Now(), 5*time.Second)
	})

	t.Run("clock time can be frozen", func(t *testing.T) {
		t.Parallel()
		frozenTime := time.Date(2026, 12, 6, 21, 37, 12, 123, time.UTC)
		clock := clock.NewClock()
		clock.FreezeTime(frozenTime)

		now := clock.Now()

		assert.Equal(t, frozenTime, now)
	})
}
