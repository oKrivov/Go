// 1️⃣ Версия 1 — БЕЗ context ❌
// ❌ Проблема: клиент ушёл, сервер всё равно работает

// package main

// import (
// 	"fmt"
// 	"net/http"
// 	"time"
// )

// func hello(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("handler started")

// 	time.Sleep(5 * time.Second) // имитация долгой работы

// 	fmt.Fprintln(w, "hello")
// 	fmt.Println("handler finished")
// }

// func main() {
// 	http.HandleFunc("/hello", hello)

// 	fmt.Println("day23 server started :8080")
// 	http.ListenAndServe(":8080", nil)
// }

// 2️⃣ Версия 2 — С context ✅
// 🎯 Цель: остановить работу, если клиент ушёл
// package main

// import (
// 	"fmt"
// 	"net/http"
// 	"time"
// )

// func hello(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("handler started")

// 	ctx := r.Context() // ← ВАЖНО

// 	select {
// 	case <-time.After(5 * time.Second):
// 		fmt.Fprintln(w, "hello")
// 		fmt.Println("handler finished normally")

// 	case <-ctx.Done():
// 		fmt.Println("handler canceled:", ctx.Err())
// 		return
// 	}
// }

// func main() {
// 	http.HandleFunc("/hello", hello)

// 	fmt.Println("server started :8080")
// 	http.ListenAndServe(":8080", nil)
// }

// 3️⃣ Версия 3 — HTTP + worker pool + context 🚀
// package main

// import (
// 	"context"
// 	"fmt"
// 	"net/http"
// 	"sync"
// 	"time"
// )

// type Job struct {
// 	ctx  context.Context
// 	w    http.ResponseWriter
// 	done chan struct{}
// }

// var jobs = make(chan Job, 2)

// func worker(id int, wg *sync.WaitGroup) {
// 	defer wg.Done()

// 	for job := range jobs {
// 		fmt.Printf("worker %d started\n", id)

// 		select {
// 		case <-time.After(3 * time.Second):
// 			fmt.Fprintln(job.w, "done")
// 			fmt.Printf("worker %d finished\n", id)
// 			close(job.done)

// 		case <-job.ctx.Done():
// 			fmt.Printf("worker %d canceled request\n", id)
// 			close(job.done)
// 		}
// 	}
// }

// func hello(w http.ResponseWriter, r *http.Request) {
// 	job := Job{
// 		ctx:  r.Context(),
// 		w:    w,
// 		done: make(chan struct{}),
// 	}

// 	select {
// 	case jobs <- job:
// 		fmt.Println("job sent to worker")
// 	// case <-r.Context().Done():
// 	// 	return
// 	default:
// 		http.Error(w, "server busy", http.StatusTooManyRequests)
// 		return
// 	}

// 	select {
// 	case <-job.done:
// 	case <-r.Context().Done():
// 	}
// }

// func main() {
// 	var wg sync.WaitGroup

// 	// стартуем воркеры
// 	for i := 1; i <= 2; i++ {
// 		wg.Add(1)
// 		go worker(i, &wg)
// 	}

// 	http.HandleFunc("/hello", hello)

// 	// go func() {
// 	fmt.Println("server started :8080")
// 	http.ListenAndServe(":8080", nil)
// 	// }()

// 	// просто держим main живым
// 	// select {}
// }

package main

import (
	"fmt"
	"net/http"
	"time"
)

// func hello(w http.ResponseWriter, r *http.Request) {
// 	ctx := r.Context()
// 	fmt.Fprintln(w, "Hello, Oleg!")
// 	select {
// 	case <-time.After(2 * time.Second):
// 		fmt.Fprintln(w, "done")
// 	case <-ctx.Done():
// 		fmt.Println("client gone")
// 	}
// }

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("request started")

	select {
	case <-time.After(3 * time.Second):
		fmt.Fprintln(w, "done")
		fmt.Println("response sent")

	case <-r.Context().Done():
		fmt.Println("request canceled")
	}
}

func main() {
	http.HandleFunc("/hello", hello)

	fmt.Println("server started :8080")

	http.ListenAndServe(":8080", nil)
}
