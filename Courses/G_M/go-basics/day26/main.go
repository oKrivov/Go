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

var jobs = make(chan Job, 2)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		fmt.Printf("worker %d started\n", id)

		select {
		case <-time.After(3 * time.Second):
			fmt.Fprintln(job.w, "done")
			fmt.Printf("worker %d finished\n", id)

		case <-job.ctx.Done():
			fmt.Printf("worker %d canceled\n", id)
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
		fmt.Println("job accepted")
	default:
		http.Error(w, "server busy", http.StatusTooManyRequests)
	}
}

func main() {
	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)

	defer stop()
	var wg sync.WaitGroup

	// стартуем workers
	for i := 1; i <= 2; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("hello", hello)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	go func() {
		fmt.Println("server started :8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Println("server error:", err)
		}
	}()

	// ждём сигнал остановки
	<-ctx.Done()
	fmt.Println("\nshutdown signal received")

	// останавливаем HTTP сервер (не принимает новые запросы)
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	server.Shutdown(shutdownCtx)

	// закрываем очередь — workers доработают
	close(jobs)

	// ждём workers
	wg.Wait()
	fmt.Println("all workers stopped")
}
