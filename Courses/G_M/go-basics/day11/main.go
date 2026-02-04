package main

import "fmt"

func main() {
	jobs := make(chan int)
	results := make(chan int)

	// fan-out: 3 worker
	for i := 1; i <= 3; i++ {
		go func() {
			for job := range jobs {
				results <- job * 2
			}
		}()
	}

	// отправляем задания
	go func() {
		for i := 1; i <= 10; i++ {
			jobs <- i
		}
		close(jobs)
	}()

	// fan-in: собираем результаты
	sum := 0
	for i := 0; i < 10; i++ {
		sum += <-results
	}

	fmt.Println("sum:", sum)
}

/*
Есть jobs := make(chan int)

Есть results := make(chan int)

Запусти 3 worker goroutine

Каждый worker:

читает из jobs

отправляет job * 2 в results

Отправь в jobs числа 1..10

Собери и выведи сумму результатов*/

// func main() {
// 	jobs := make(chan int)
// 	results := make(chan int)

// 	for i := 1; i <= 3; i++ {
// 		go worker(i, jobs, results)
// 	}

// 	go func() {
// 		for i := 1; i <= 10; i++ {
// 			jobs <- i
// 		}
// 		close(jobs)
// 	}()

// }

// func worker(id int, jobs <-chan int, results chan<- int) {
// 	for job := range jobs {
// 		fmt.Printf("worker %d job %d\n", id, job)
// 		results <- job * 2
// 	}
// }

// шаг 1

// go func() {
// 	id := 1
// 	fmt.Printf("worker %v started\n", id)
// 	fmt.Printf("worker %v finished\n", id)
// }()
// go func() {
// 	id := 2
// 	fmt.Printf("worker %v started\n", id)
// 	fmt.Printf("worker %v finished\n", id)
// }()
// go func() {
// 	id := 3
// 	fmt.Printf("worker %v started\n", id)
// 	fmt.Printf("worker %v finished\n", id)
// }()

// time.Sleep(time.Second)

// Шаг 3

// counter := 0

// var mu sync.Mutex
// var wg sync.WaitGroup

// for i := 0; i < 100; i++ {
// 	wg.Add(1)

// 	go func() {
// 		defer wg.Done()
// 		mu.Lock()
// 		counter++
// 		mu.Unlock()
// 	}()
// }

// wg.Wait()
// fmt.Println(counter)

// шаг 5
// ch := make(chan int)

// go func() {
// 	for i := 0; i < 100; i++ {
// 		ch <- 1
// 	}
// 	close(ch)
// }()

// sum := 0
// for v := range ch {
// 	sum += v
// }

// fmt.Println(sum)
