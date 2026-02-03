package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// ограничение на параллельную работу
var workers = make(chan struct{}, 2)

func heavyWork(ctx context.Context) error {
	select {
	case <-time.After(3 * time.Second):
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// пробуем занять worker
	select {
	case workers <- struct{}{}:
		// слот занят
		defer func() { <-workers }()
	default:
		http.Error(w, "server busy", http.StatusTooManyRequests)
		return
	}

	// выполняем работу
	err := heavyWork(ctx)
	if err != nil {
		http.Error(w, "request canceled", http.StatusRequestTimeout)
		return
	}

	fmt.Fprintln(w, "done")
}

func main() {
	http.HandleFunc("/hello", hello)

	fmt.Println("server started :8080")
	http.ListenAndServe(":8080", nil)
}
