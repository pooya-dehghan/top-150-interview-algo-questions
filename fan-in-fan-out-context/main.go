package main

import (
	"context"
	"fmt"
	"time"
)

func sendJobs(jobs []int, jobsCh chan<- int) {
	for _, v := range jobs {
		jobsCh <- v
	}

	close(jobsCh)
}

func readers(jobsCh <-chan int, resultCh chan<- int, ctx context.Context) {
	for {
		select {
		case v, ok := <-jobsCh:
			if !ok {
				return
			}

			time.Sleep(500 * time.Millisecond)

			resultCh <- v * v
		case <-ctx.Done():
			return
		}

	}
}

func main() {

	jobs := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	jobsCh := make(chan int)

	resultCh := make(chan int)

	ctx, cancel := context.WithCancel(context.Background())

	go sendJobs(jobs, jobsCh)

	workerCount := 3

	for range workerCount {
		go readers(jobsCh, resultCh, ctx)
	}

	fmt.Println("results:")
	for res := range resultCh {
		fmt.Println(res)
	}

	func() {
		close(resultCh)
		cancel()
	}()
}
