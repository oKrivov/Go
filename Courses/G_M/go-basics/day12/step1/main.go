// Шаг 1
package main

import (
	"fmt"
	"time"
)

func main() {
	jobs := make(chan string)
	done := make(chan struct{})

	// producer
	go func() {
		defer close(jobs)

		for {
			select {
			case <-done:
				return
			default:
				jobs <- "working..."
				time.Sleep(300 * time.Millisecond)
			}
		}
	}()

	// consumer
	go func() {
		for {
			select {
			case job, ok := <-jobs:
				if !ok {
					return
				}
				fmt.Println(job)
			case <-done:
				fmt.Println("stopped")
				return

			}
		}
	}()
	time.Sleep(1 * time.Second)
	close(done)

	time.Sleep(1 * time.Second)
}

// func main() {
// 	jobs := make(chan int)
// 	done := make(chan struct{})

// 	go func() {
// 		for i := 1; i <= 5; i++ {
// 			jobs <- i
// 			time.Sleep(300 * time.Millisecond)
// 		}
// 		close(done)
// 	}()

// 	go func() {
// 		for {
// 			select {
// 			case job := <-jobs:
// 				fmt.Println("job:", job)
// 			case <-done:
// 				fmt.Println("shutdown")
// 				return
// 			}
// 		}
// 	}()

// 	time.Sleep(2 * time.Second)
// }
