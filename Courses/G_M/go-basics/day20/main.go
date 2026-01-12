package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func worker(ctx context.Context, id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("worker %d stopped\n", id)
			return
		case job, ok := <-jobs:
			if !ok {
				fmt.Printf("worker %d: jobs closed\n", id)
				return
			}
			fmt.Printf("worker %d proccesing job %v\n", id, job)
			time.Sleep(300 * time.Microsecond)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	jobs := make(chan int, 5)
	var wg sync.WaitGroup

	// start workers
	for i := 1; i <= 2; i++ {
		wg.Add(1)
		go worker(ctx, i, jobs, &wg)
	}

	// producer
	go func() {
		defer close(jobs)

		i := 1
		for {
			select {
			case <-ctx.Done():
				fmt.Println("producer stopped")
				return
			case jobs <- i:
				fmt.Println("send job", i)
				i++
				time.Sleep(300 * time.Millisecond)
			}
		}
	}()
	// канал для сигналов ОС
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	// ждём сигнал
	<-sigCh
	fmt.Println("\nmain: shutdown signal received")

	// запускаем graceful shutdown
	cancel()

	// wait workers
	wg.Wait()
	fmt.Println("main clossed all workers")
}

/*
1️⃣ Кто инициирует shutdown?
main
2️⃣ Почему jobs закрывает producer, а не main?
кто пишет в канал тот и закрывает его
3️⃣ Зачем signal.Notify, если есть context?
signal.Notify получает внешний сигнал остановки программы.
4️⃣ Что будет, если убрать wg.Wait()?
воркеры умрут после завершения main
5️⃣ Что будет, если не закрыть jobs?
deadlock
*/
