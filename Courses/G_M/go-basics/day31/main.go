package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type Job struct {
	ctx    context.Context
	result chan string
}

var jobs = make(chan Job, 2)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		fmt.Printf("worker %d started\n", id)

		select {
		case <-time.After(5 * time.Second):
			job.result <- "done by worker"
		case <-job.ctx.Done():
			fmt.Printf("worker %d canceled\n", id)
		}
	}
}

func timeoutMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
		defer cancel()

		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

func hello(w http.ResponseWriter, r *http.Request) {
	job := Job{
		ctx:    r.Context(),
		result: make(chan string),
	}

	select {
	case jobs <- job:
		fmt.Println("job sent")

	case <-r.Context().Done():
		http.Error(w, "request canceled", http.StatusRequestTimeout)
		return
	}

	select {
	case res := <-job.result:
		fmt.Fprintln(w, res)
	case <-r.Context().Done():
		http.Error(w, "timeout", http.StatusRequestTimeout)
	}
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 2; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/hello", hello)

	// go func() {
	fmt.Println("server started :8080")
	http.ListenAndServe(":8080", timeoutMiddleware(mux))
	// }()

	// select {}
}
