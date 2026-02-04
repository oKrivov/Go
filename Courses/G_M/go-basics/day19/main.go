package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func worker(ctx context.Context, id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("worker %d stopped\n", id)
			return
		case job, ok := <-jobs:
			if !ok {
				return
			}
			fmt.Printf("worker %d proccesing job %v\n", id, job)
			time.Sleep(300 * time.Millisecond)
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	jobs := make(chan int, 2)

	var wg sync.WaitGroup

	defer cancel()

	// start 2 workers
	for i := 1; i <= 2; i++ {
		wg.Add(1)
		go worker(ctx, i, jobs, &wg)
	}
	// start producer

	go func() {
		defer close(jobs)

		for i := 1; i > 0; i++ {
			select {
			case <-ctx.Done():
				return
			case jobs <- i:
			}
		}
	}()

	// wait workers
	wg.Wait()
	fmt.Println("main: all workers stopped")
}
