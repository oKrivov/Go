package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, done <-chan struct{}) {
	for {
		select {
		case job, ok := <-jobs:
			if !ok {
				fmt.Printf("worker %d: jobs closed\n", id)
				return
			}
			fmt.Printf("worker %d processing job %d\n", id, job)
			time.Sleep(500 * time.Millisecond)

		case <-done:
			fmt.Printf("worker %d stopped\n", id)
			return
		}
	}
}

func main() {
	jobs := make(chan int)
	done := make(chan struct{})

	// запускаем 3 worker
	for i := 1; i <= 3; i++ {
		go worker(i, jobs, done)
	}

	// producer
	go func() {
		for i := 1; i <= 10; i++ {
			jobs <- i
			time.Sleep(200 * time.Millisecond)
		}
		close(jobs)
	}()

	time.Sleep(2 * time.Second)
	close(done)

	time.Sleep(1 * time.Second)
}
