package in_memory_test

import (
	"sync"
	"testing"
)

func runConcurrently[T any](t *testing.T, items []T, worker func(T) error) {
	t.Helper()
	var wg sync.WaitGroup
	errCh := make(chan error, len(items))

	for _, it := range items {
		wg.Go(func() {
			if err := worker(it); err != nil {
				errCh <- err
			}
		})
	}

	wg.Wait()
	close(errCh)
	for err := range errCh {
		t.Fatalf("concurrent work failed: %v", err)
	}
}
