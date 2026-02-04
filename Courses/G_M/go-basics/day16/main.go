package main

import (
	"context"
	"fmt"
	"sync"
)

func generator(ctx context.Context, n int) <-chan int {
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

func square(
	ctx context.Context,
	in <-chan int,
	errCh chan<- error,
	cancel context.CancelFunc,
) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for n := range in {
			if n > 10 {
				select {
				case errCh <- fmt.Errorf("square error %d > 10", n):
				default:
				}
				cancel()
				return
			}

			select {
			case <-ctx.Done():
				return
			case out <- n * n:
			}
		}
	}()

	return out
}

func double(
	ctx context.Context,
	in <-chan int,
	errCh chan<- error,
	cancel context.CancelFunc,
) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for n := range in {

			if n > 400 {
				select {
				case errCh <- fmt.Errorf("double error %d > 40", n):
				default:
				}
				cancel()
				return
			}

			select {
			case <-ctx.Done():
				return
			case out <- n * 2:
			}
		}
	}()

	return out
}

func triple(
	ctx context.Context,
	in <-chan int,
	errCh chan<- error,
	cancel context.CancelFunc,
) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		for n := range in {
			if n > 1500 {
				select {
				case errCh <- fmt.Errorf("triple error %d > 50", n):
				default:
				}
				cancel()
				return
			}
			select {
			case <-ctx.Done():
				return
			case out <- n * 3:
			}
		}
	}()

	return out
}

func merge(ctx context.Context, chans ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup

	wg.Add(len(chans))

	for _, ch := range chans {
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
	ctx, cancel := context.WithCancel(context.Background())
	errCh := make(chan error, 1)

	defer cancel()

	gen := generator(ctx, 10)

	sq1 := square(ctx, gen, errCh, cancel)
	sq2 := square(ctx, gen, errCh, cancel)
	sq3 := square(ctx, gen, errCh, cancel)

	mergedSquares := merge(ctx, sq1, sq2, sq3)

	db := double(ctx, mergedSquares, errCh, cancel)
	tr := triple(ctx, db, errCh, cancel)

	for v := range tr {
		fmt.Println(v)
	}

	go func() {
		if err := <-errCh; err != nil {
			fmt.Println("pipeline error:", err)
			cancel()
		}
	}()

	for v := range tr {
		fmt.Println(v)
	}

	fmt.Println("main done")
}
