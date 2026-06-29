// concpat1 — worker pool
// A fixed set of worker goroutines pulls jobs from one channel and pushes
// results to another. Implement the worker so the pool produces every result.

// I AM NOT DONE
package main_test

import (
	"sort"
	"sync"
	"testing"
)

// square is a worker: it reads jobs and sends each job's square to results.
func square(jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	// FIXME: range the jobs channel and send each square to results.
	_ = jobs
}

// runPool runs `workers` squarers over inputs and returns the sorted results.
func runPool(workers int, inputs []int) []int {
	jobs := make(chan int)
	results := make(chan int, len(inputs))
	var wg sync.WaitGroup

	for i := 0; i < workers; i++ {
		wg.Add(1)
		go square(jobs, results, &wg)
	}

	go func() {
		for _, n := range inputs {
			jobs <- n
		}
		close(jobs)
	}()

	wg.Wait()
	close(results)

	out := make([]int, 0, len(inputs))
	for r := range results {
		out = append(out, r)
	}
	sort.Ints(out)
	return out
}

func TestWorkerPool(t *testing.T) {
	got := runPool(3, []int{1, 2, 3, 4, 5})
	want := []int{1, 4, 9, 16, 25}
	if len(got) != len(want) {
		t.Fatalf("want %v, got %v", want, got)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("at %d: want %d, got %d", i, want[i], got[i])
		}
	}
}
