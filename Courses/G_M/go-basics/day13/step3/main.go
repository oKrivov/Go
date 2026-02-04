package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	jobs := make(chan int)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)

	defer cancel()

	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(ctx, i, jobs, &wg)
	}

	go func() {
		for i := 1; i <= 10; i++ {
			time.Sleep(300 * time.Millisecond)
			jobs <- i
		}
		close(jobs)
	}()

	wg.Wait()
	fmt.Println("main: all workers stopped")
}

func worker(ctx context.Context, id int, jobs chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("worker %d clossed context: %v\n", id, ctx.Err())
			return
		case job, ok := <-jobs:
			if !ok {
				fmt.Printf("worker %d, jobs closed\n", id)
				return
			}
			fmt.Printf("worker %d proccesing job %v\n", id, job)
		}
	}
}
