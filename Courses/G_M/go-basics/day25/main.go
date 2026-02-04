package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type Job struct {
	ctx      context.Context
	resultCh chan string
}

var jobs = make(chan Job, 1)

func worker(id int) {
	for job := range jobs {
		fmt.Printf("worker %d started\n", id)

		select {
		case <-time.After(3 * time.Second):
			job.resultCh <- "done"
			fmt.Printf("worker %d finished\n", id)
		case <-job.ctx.Done():
			fmt.Printf("worker %d canceled\n", id)
		}
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	resultCh := make(chan string, 1)

	job := Job{
		ctx:      r.Context(),
		resultCh: resultCh,
	}

	// пытаемся положить задачу в очередь
	select {
	case jobs <- job:
		fmt.Println("job accepted")
	default:
		http.Error(w, "server busy", http.StatusTooManyRequests)
		return
	}

	// держим HTTP-запрос
	select {
	case result := <-resultCh:
		fmt.Fprintln(w, result)
	case <-r.Context().Done():
		fmt.Println("client ccanceled request")
	}
}

func main() {
	// стартуем 2 worker
	for i := 1; i <= 1; i++ {
		go worker(i)
	}

	http.HandleFunc("/hello", hello)
	fmt.Println("server started :8080")
	http.ListenAndServe(":8080", nil)
}
