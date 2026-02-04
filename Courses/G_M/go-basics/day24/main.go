package main

import (
	"context"
	"fmt"
	"time"
)

// type Job struct {
// 	id int
// }

// var jobs = make(chan Job, 5)

// func worker(id int) {
// 	for job := range jobs {
// 		fmt.Printf("worker %d started job %d\n", id, job.id)
// 		time.Sleep(3 * time.Second)
// 		fmt.Printf("worker %d finished job %d\n", id, job.id)
// 	}
// }

// func hello(w http.ResponseWriter, r *http.Request) {
// 	job := Job{id: time.Now().Nanosecond()}

// 	select {
// 	case jobs <- job:
// 		w.WriteHeader(http.StatusAccepted)
// 		fmt.Fprintln(w, "job accepted")
// 	default:
// 		http.Error(w, "server busy", http.StatusTooManyRequests)
// 	}
// }

// func main() {
// 	for i := 1; i <= 2; i++ {
// 		go worker(i)
// 	}
// 	http.HandleFunc("/hello", hello)
// 	fmt.Println("server started :8080")
// 	http.ListenAndServe(":8080", nil)
// }

// #############################################################################

// STEP 1
// Simple HTTP-server

// func hello(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "Hello, world")
// }

// func main() {
// 	http.HandleFunc("/hello", hello)

// 	fmt.Println("server started :8080")
// 	http.ListenAndServe(":8080", nil)
// }

// STEP 2
// r.Context()
// func hello(w http.ResponseWriter, r *http.Request) {
// 	ctx := r.Context()

// 	select {
// 	case <-ctx.Done():
// 		fmt.Println("request canceld")
// 		return
// 	default:
// 		fmt.Fprintln(w, "Hello, WORLD")
// 	}
// }

//  STEP 3
// Отмена запроса

// func hello(w http.ResponseWriter, r *http.Request) {
// 	ctx := r.Context()

// 	select {
// 	case <-time.After(3 * time.Second):
// 		fmt.Fprintln(w, "DONE")
// 		fmt.Println("done")

// 	case <-ctx.Done():
// 		fmt.Println("request canceld")

// 	}
// }

// #############################################################################

// context.WithCancel

// func worker(ctx context.Context) {
// 	<-ctx.Done()
// 	fmt.Println("worker stoped", ctx.Err())
// }

// func main() {
// 	ctx, cancel := context.WithCancel(context.Background())

// 	go worker(ctx)

// 	time.Sleep(2 * time.Second)
// 	cancel() // ! отмена

// 	time.Sleep(time.Second)
// }

// context.WithTimeout / WithDeadline

func worker(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("timeout", ctx.Err())
	case <-time.After(5 * time.Second):
		fmt.Println("done")
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	go worker(ctx)

	time.Sleep(5 * time.Second)
}
