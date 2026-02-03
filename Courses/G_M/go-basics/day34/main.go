// Timeouts, deadlines, graceful shutdown
package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func hello(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
	defer cancel()

	fmt.Println("incoming request:", r.Method, r.URL.Path)

	select {
	case <-time.After(4 * time.Second):
		fmt.Fprintln(w, "done")
	case <-ctx.Done():
		fmt.Println("handler canceled:", ctx.Err())
		http.Error(w, "timeout", http.StatusRequestTimeout)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", hello)

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	go func() {
		fmt.Println("server started :8080")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	<-sig

	fmt.Println("\nshutdown signal received")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	srv.Shutdown(ctx)
	fmt.Println("server stopped gracefully")
}
