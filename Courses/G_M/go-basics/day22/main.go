// ##################
//
//	Step 1
//
// HTTP + Worker Pool + Backpressure
//
// ##################
package main

import (
	"fmt"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	select {
	case jobs <- w:
		// отправили в очередь
	default:
		// очередь переполнена
		http.Error(w, "server busy", http.StatusTooManyRequests)
	}
}

var jobs = make(chan http.ResponseWriter, 5)

func worker(id int, jobs <-chan http.ResponseWriter) {
	for w := range jobs {
		fmt.Printf("worker %d started\n", id)
		time.Sleep(1 * time.Second)
		fmt.Fprintln(w, "done")
	}
}

func main() {
	for i := 1; i <= 3; i++ {
		go worker(i, jobs)
	}

	http.HandleFunc("/work", handler)

	fmt.Println("server started :8080")
	http.ListenAndServe(":8080", nil)
}
