package main

import (
	"context"
	"fmt"
	"time"
)

// first step generator
func generator(list []int, ctx context.Context) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)

		for _, v := range list {
			select {
			case <-ctx.Done():
				return
			case out <- v:
			}
		}
	}()

	return out
}

// second step multiplier
func multiplier(mlCh <-chan int, ctx context.Context) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for {
			select {
			case v, ok := <-mlCh:
				if !ok {
					return
				}
				time.Sleep(200 * time.Millisecond)
				select {
				case out <- v * 2:
				case <-ctx.Done():
					return
				}
			case <-ctx.Done():
				return
			}
		}
	}()

	return out
}

// third step printer
func printer(printCh <-chan int) {
	for v := range printCh {
		fmt.Printf(" %v : ", v)
	}
}

func main() {
	list := []int{1, 2, 3, 4, 5, 6, 7, 8}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mlCh := generator(list, ctx)
	prCh := multiplier(mlCh, ctx)
	printer(prCh)

}
