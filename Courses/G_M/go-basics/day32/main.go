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

func main() {
	// 1. корневой контекст приложения
	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)

	defer stop()

	// 2. HTTP server
	server := &http.Server{
		Addr:    ":8080",
		Handler: http.HandlerFunc(hello),
	}

	go func() {
		fmt.Println("server started :8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Println("server error:", err)
		}
	}()

	// 4. ждем сигнал завершения
	<-ctx.Done()
	fmt.Println("\nshutdown signal")

	// 5. корректно останавливаем работу
	sutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	server.Shutdown(sutdownCtx)
	fmt.Println("server stoped")
}

func hello(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	fmt.Println("request started")

	select {
	case <-time.After(5 * time.Second):
		fmt.Fprintln(w, "hello")
		fmt.Println("equest finished")
	case <-ctx.Done():
		fmt.Println("equest canceld")
	}
}
