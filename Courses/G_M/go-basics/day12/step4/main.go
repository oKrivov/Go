package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	jobs := make(chan int)

	// 🔹 fan-out: запускаем 3 worker
	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(ctx, i, jobs, &wg)
	}
	// 🔹 producer
	go func() {
		for i := 1; i <= 10; i++ {
			jobs <- i
			time.Sleep(300 * time.Millisecond)
		}

		close(jobs)
	}()

	// 🔹 даём поработать
	time.Sleep(2 * time.Second)
	fmt.Println("main: cancel context")
	cancel()

	// 🔹 ждём корректного завершения
	wg.Wait()
	fmt.Println("main: all workers stopped")

}

func worker(ctx context.Context, id int, jobs chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("worker %d stopped: %v\n", id, ctx.Err())
			return
		case job, ok := <-jobs:
			if !ok {
				fmt.Printf("worker %d: jobs cancel closed\n", id)
				return
			}
			fmt.Printf("worker %d proccesing job %d\n", id, job)
			time.Sleep(500 * time.Millisecond)
		}
	}
}
