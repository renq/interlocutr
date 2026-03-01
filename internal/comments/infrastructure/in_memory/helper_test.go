package in_memory_test

import (
	"fmt"
	"sync"
	"testing"
)

func makeIDs(n int) []string {
	ids := make([]string, n)
	for i := 0; i < n; i++ {
		ids[i] = fmt.Sprintf("id-%d", i)
	}
	return ids
}

func runConcurrently[T any](t *testing.T, items []T, worker func(T) error) {
	t.Helper()
	var wg sync.WaitGroup
	errCh := make(chan error, len(items))

	for _, it := range items {
		it := it
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := worker(it); err != nil {
				errCh <- err
			}
		}()
	}

	wg.Wait()
	close(errCh)
	for err := range errCh {
		t.Fatalf("concurrent work failed: %v", err)
	}
}
