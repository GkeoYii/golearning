package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	ctx, stop := context.WithTimeout(context.Background(), 3*time.Second)
	defer stop()

	ch := make(chan int, 10)
	go func() {
		defer wg.Done()
		producer(ctx, 3, ch)
	}()
	go func() {
		defer wg.Done()
		producer(ctx, 5, ch)
	}()
	go consumer(ch)
	wg.Wait()
}

func producer(ctx context.Context, factor int, out chan<- int) {

	for i := 1; ; i++ {
		select {
		case <-ctx.Done():
			fmt.Println("producer end")
			return
		default:
			fmt.Println("product : ", i*factor)
			out <- i * factor
		}
		time.Sleep(1 * time.Second)
	}
}

func consumer(in <-chan int) {
	for v := range in {
		fmt.Println("consume : ", v)
	}
}
