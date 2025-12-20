package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// step 1

func generator(n int, ctx context.Context) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for i := 1; i <= n; i++ {
			select {
			case <-ctx.Done():
				return
			case out <- i:
			}
		}
	}()

	return out
}

// func square(in <-chan int) <-chan int {
// 	out := make(chan int)

// 	go func() {
// 		defer close(out)

// 		for n := range in {
// 			out <- n * n
// 			time.Sleep(100 * time.Millisecond)
// 		}
// 	}()

//		return out
//	}

// func squareFanOut(ctx context.Context, in <-chan int, workers int) <-chan int {
// 	out := make(chan int)
// 	var wg sync.WaitGroup

// 	wg.Add(workers)

// 	for i := 0; i < workers; i++ {
// 		go func() {
// 			defer wg.Done()

// 			for n := range in {
// 				select {
// 				case <-ctx.Done():
// 					return
// 				case out <- n * n:
// 				}
// 			}
// 		}()
// 	}

// 	go func() {
// 		wg.Wait()
// 		close(out)
// 	}()

// 	return out
// }

func double(ctx context.Context, in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for n := range in {
			select {
			case <-ctx.Done():
				return
			case out <- n + n:
			}
		}
	}()

	return out
}

func square(ctx context.Context, in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for n := range in {
			select {
			case <-ctx.Done():
				return
			case out <- n * n:
			}
		}
	}()

	return out
}

func merge(ctx context.Context, chans ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup

	for _, ch := range chans {
		wg.Add(1)
		go func(c <-chan int) {
			defer wg.Done()
			for v := range ch {
				select {
				case <-ctx.Done():
					return
				case out <- v:
				}
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	in := generator(10, ctx)

	// fan-out
	// sq := squareFanOut(ctx, gen, 3)
	c1 := square(ctx, in)
	c2 := square(ctx, in)
	c3 := square(ctx, in)

	// fan-in
	merged := merge(ctx, c1, c2, c3)

	// next stage
	dbl := double(ctx, merged)

	for v := range dbl {
		fmt.Println(v)
	}
}
