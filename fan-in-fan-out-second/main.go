package main

import (
	"fmt"
	"sync"
	"time"
)

func sendJobs(jobs []int, jobsCh chan<- int) {
	for _, v := range jobs {
		jobsCh <- v
	}

	close(jobsCh)
}

func readers(jobsCh <-chan int, resultCh chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for v := range jobsCh {
		time.Sleep(500 * time.Millisecond)

		resultCh <- v * v

	}
}

func main() {
	wg := sync.WaitGroup{}

	jobs := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	jobsCh := make(chan int)

	resultCh := make(chan int)

	go sendJobs(jobs, jobsCh)

	go func() {
		wg.Wait()
		close(resultCh)
	}()

	workerCount := 3

	wg.Add(workerCount)

	for range workerCount {
		go readers(jobsCh, resultCh, &wg)
	}

	fmt.Println("results:")
	for res := range resultCh {
		fmt.Println(res)
	}
}
