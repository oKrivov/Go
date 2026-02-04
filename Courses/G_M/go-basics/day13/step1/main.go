package main

import (
	"fmt"
	"sync"
)

func main() {
	jobs := make(chan int)

	// TODO: запусти 3 worker
	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, jobs, &wg)
	}

	// TODO: отправь 10 jobs
	go func() {
		for i := 1; i <= 10; i++ {
			jobs <- i
		}
		// TODO: закрой jobs
		close(jobs)
	}()

	// TODO: дождись worker

	wg.Wait()
	fmt.Println("main: all workers stopped")
}

func worker(id int, jobs chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("worker %d proccesing job %d\n", id, job)
	}
}
