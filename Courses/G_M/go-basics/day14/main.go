package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/*
// step 1
func main() {
	ch := make(chan int, 2)

	fmt.Println("sen 1")
	ch <- 1
	// fmt.Println(<-ch)

	fmt.Println("sen 2")
	ch <- 2
	// fmt.Println(<-ch)

	fmt.Println("sen 3")
	ch <- 3 //где зависнет?
	// fmt.Println(<-ch)

	fmt.Println("done")

		// 	Вопросы:
		// 	1. Какие строки выведутся?
		// 	2. Почему?

		// 	Ответы:
		// 	1.
		// 	fatal error: all goroutines are asleep - deadlock!
		// goroutine 1 [chan send]:
		// main.main()
		//         /Users/Oleg/Desktop/s_21/Go/Courses/G_M/go-basics/day14/main.go:18 +0xf8
		// exit status 2
		// 	2.
		// 	Выводится потому что канал заблакирован,
		// 	т.е. мы записали в канал 2 занчания,но ничего не считали,
		// 	После принятия 2-го значения канал заблакировался,
		// 	но вы дальше передаем туда еще одно значение
}



// step 2
func main() {
	jobs := make(chan int, 5)
	ctx, cancel := context.WithTimeout(context.Background(), 2000*time.Millisecond)

	defer cancel()

	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(ctx, i, jobs, &wg)
	}

	go func() {
		for i := 1; i <= 10; i++ {
			time.Sleep(300 * time.Millisecond)
			jobs <- i
		}
		close(jobs)
	}()

	wg.Wait()
	fmt.Println("main: all workers stoped")
}

func worker(ctx context.Context, id int, jobs chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("worker %d canceled: %v\n", id, ctx.Err())
			return
		case job, ok := <-jobs:
			if !ok {
				fmt.Printf("worker %d, jobs is clossed\n", id)
				return
			}
			fmt.Printf("worker %d proccesing job %v\n", id, job)
		}
	}
}

*/

// step3

func main() {
	jobs := make(chan int, 5)
	ctx, cancel := context.WithTimeout(context.Background(), 2000*time.Millisecond)

	defer cancel()

	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(ctx, i, jobs, &wg)
	}

	go func() {
		for i := 1; i <= 10; i++ {
			time.Sleep(300 * time.Millisecond)
			select {
			case jobs <- i:
				fmt.Println("send", i)
			case <-ctx.Done():
				fmt.Println("produser stopped")
				return
			}
		}
		close(jobs)
	}()

	wg.Wait()
	fmt.Println("main: all workers stoped")
}

func worker(ctx context.Context, id int, jobs chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("worker %d canceled: %v\n", id, ctx.Err())
			return
		case job, ok := <-jobs:
			if !ok {
				fmt.Printf("worker %d, jobs is clossed\n", id)
				return
			}
			fmt.Printf("worker %d proccesing job %v\n", id, job)
		}
	}
}

// step 4

// Ответь своими словами:
// Чем buffered канал отличается от unbuffered?
// буферезированный канал, может принимать столько значений сколько в нем указано прежде чем заблокируется,
// в отличии от не буферещированного который блокируется сразу после того как приял значение.

// Почему буфер не заменяет context?
// буфер не дает сигнал остановки и закрыть канал

// Кто должен закрывать канал и почему?
// канал закрывает тот, кто передает в канал.
