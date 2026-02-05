// // ###########################
// // Pet-project (версия 1)
// // ###########################
// package main

// import (
// 	"context"
// 	"fmt"
// 	"net/http"
// 	"sync"
// 	"time"
// )

// type Job struct {
// 	ctx context.Context
// 	id  int
// }

// var jobs = make(chan Job, 2) //  ОЧЕРЕДЬ (маленькая специально)

// func worker(id int, wg *sync.WaitGroup) {
// 	defer wg.Done()

// 	for job := range jobs {
// 		fmt.Printf("worker %d: start job: %d\n", id, job.id)

// 		select {
// 		case <-time.After(3 * time.Second):
// 			fmt.Printf("worker %d: done job: %d\n", id, job.id)
// 		case <-job.ctx.Done():
// 			fmt.Printf("worker %d:  job: %d canceled\n", id, job.id)

// 		}
// 	}
// }

// func taskHandler(w http.ResponseWriter, r *http.Request) {
// 	job := Job{
// 		ctx: r.Context(),
// 		id:  time.Now().Nanosecond(),
// 	}

// 	select {
// 	case jobs <- job:
// 		fmt.Fprintln(w, "task accepted")
// 	default:
// 		http.Error(w, "server busy", http.StatusTooManyRequests)
// 	}
// }

// func main() {
// 	var wg sync.WaitGroup

// 	// стартуем воркеры
// 	for i := 1; i <= 2; i++ {
// 		wg.Add(1)
// 		go worker(i, &wg)
// 	}

// 	http.HandleFunc("/task", taskHandler)

// 	fmt.Println("server started :8080")
// 	http.ListenAndServe(":8080", nil)

// 	// (сюда код не дойдёт без shutdown)
// 	close(jobs)
// 	wg.Wait()
// }

// ###########################
// handler ЖДЁТ worker (sync)
// ###########################

// package main

// import (
// 	"context"
// 	"fmt"
// 	"net/http"
// 	"time"
// )

// type Job struct {
// 	ctx    context.Context
// 	result chan string
// }

// var jobs = make(chan Job, 2)

// func worker(id int) {
// 	for job := range jobs {
// 		fmt.Printf("worker %d: start\n", id)
// 		select {
// 		case <-time.After(3 * time.Second):
// 			job.result <- "done by worker"
// 		case <-job.ctx.Done():
// 			fmt.Printf("worker %d: canceled", id)
// 			job.result <- "canceled"
// 		}
// 	}
// }

// func taskHandler(w http.ResponseWriter, r *http.Request) {
// 	job := Job{
// 		ctx:    r.Context(),
// 		result: make(chan string, 1),
// 	}

// 	// пытаемся положить задачу
// 	select {
// 	case jobs <- job:
// 		fmt.Println("task accepted")
// 	default:
// 		http.Error(w, "server busy", http.StatusTooManyRequests)
// 	}

// 	// КЛЮЧЕВО: handler ЖДЁТ worker
// 	select {
// 	case res := <-job.result:
// 		fmt.Fprintln(w, res)
// 	case <-r.Context().Done():
// 		fmt.Println("handler: client canceled")
// 	}
// }

// func main() {
// 	go worker(1)
// 	go worker(2)

// 	http.HandleFunc("/task", taskHandler)
// 	fmt.Println("server started :8080")
// 	http.ListenAndServe(":8080", nil)
// }

// ###########################
// graceful shutdown
// ###########################

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
	id int
}

var (
	jobs = make(chan Job, 2)
	wg   sync.WaitGroup
)

func worker(id int, ctx context.Context) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("worker %d: sutdown\n", id)
			return
		case job, ok := <-jobs:
			if !ok {
				fmt.Printf("worker %d: queue closed\n", id)
				return
			}
			fmt.Printf("worker %d: start job %d\n", id, job.id)
			time.Sleep(3 * time.Second)
			fmt.Printf("worker %d: finished job %d\n", id, job.id)

		}
	}
}

func taskHandler(w http.ResponseWriter, r *http.Request) {
	job := Job{id: time.Now().Nanosecond()}

	select {
	case jobs <- job:
		fmt.Fprintln(w, "job acceped")
	default:
		http.Error(w, "server busy", http.StatusTooManyRequests)
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	for i := 1; i <= 2; i++ {
		wg.Add(1)
		go worker(i, ctx)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/task", taskHandler)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	go func() {
		fmt.Println("server started :8080")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	<-sig
	fmt.Println("\nshitdown signal received")

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		fmt.Println("server shutdown eerror", err)
	}

	cancel()
	close(jobs)
	wg.Wait()

	fmt.Println("graceful shutdown complete")
}
