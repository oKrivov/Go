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

var sem = make(chan struct{}, 2)

func hello(w http.ResponseWriter, r *http.Request) {
	select {
	case sem <- struct{}{}:
		defer func() {
			<-sem
		}()
	default:
		http.Error(w, "server busy", http.StatusTooManyRequests)
		return
	}

	fmt.Println("server started")

	select {
	case <-time.After(3 * time.Second):
		fmt.Fprintln(w, "done")
		fmt.Println("server finished")
	case <-r.Context().Done():
		fmt.Println("equest canceled")
		return
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", hello)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	// запуск сервера
	go func() {
		fmt.Println("server started :8080")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Println("server error", err)
		}
	}()

	// ждём сигнал
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	<-sigCh

	fmt.Println("\nshutdown signal received")

	// контекст для graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// корректная остановка
	if err := srv.Shutdown(ctx); err != nil {
		fmt.Println("shutdown error:", err)
	}

	fmt.Println("server stopped gracefully")
}
