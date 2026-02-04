package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type Job struct {
	ctx context.Context
	w   http.ResponseWriter
}

var jobs = make(chan Job, 2) // очередь запросов

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		fmt.Printf("worker %d: started job\n", id)

		select {
		case <-time.After(3 * time.Second):
			fmt.Fprintln(job.w, "done")
			fmt.Printf("worker %d: finished job\n", id)

		case <-job.ctx.Done():
			fmt.Printf("worker %d: request canceled\n", id)
		}
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	job := Job{
		ctx: r.Context(),
		w:   w,
	}

	select {
	case jobs <- job:
		fmt.Println("handler: job sent to queue")
	default:
		http.Error(w, "server busy", http.StatusTooManyRequests)
	}
}

func main() {
	var wg sync.WaitGroup

	// 1️⃣ стартуем воркеры
	for i := 1; i <= 2; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	// 2️⃣ регистрируем handler
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", hello)

	// 3️⃣ создаём сервер
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	// 4️⃣ запускаем сервер
	go func() {
		fmt.Println("server started :8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Println("server error:", err)
		}
	}()

	// 5️⃣ ловим сигналы ОС
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	<-sigCh

	fmt.Println("\nshutdown signal received")

	// 6️⃣ корректно останавливаем сервер
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	server.Shutdown(ctx)
	close(jobs)

	// 7️⃣ ждём воркеры
	wg.Wait()
	fmt.Println("server stopped gracefully")
}
