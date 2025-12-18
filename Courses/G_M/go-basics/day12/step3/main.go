package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go worker(ctx)

	time.Sleep(2000 * time.Millisecond)
	cancel()

	time.Sleep(1000 * time.Millisecond)
}

func worker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("worker stopped:", ctx.Err())
			return
		default:
			fmt.Println("working...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}
