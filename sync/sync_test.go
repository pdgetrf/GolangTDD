package sync

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := NewCounter()
		counter.Inc()
		counter.Inc()
		counter.Inc()
		assertCounter(t, 3, counter)
	})

	t.Run("concurrent access", func(t *testing.T) {
		wantedCount := 1000
		counter := Counter{}

		var wg sync.WaitGroup
		wg.Add(wantedCount)

		for i := 0; i < wantedCount; i++ {
			go func(w *sync.WaitGroup) {
				counter.Inc()
				wg.Done()
			}(&wg)
		}

		wg.Wait()

		assertCounter(t, wantedCount, &counter)
	})
}

func assertCounter(t *testing.T, expect int, actual *Counter) {
	t.Helper()
	if expect != actual.Value() {
		t.Errorf("got %d, want %d", actual.Value(), expect)

	}
}
