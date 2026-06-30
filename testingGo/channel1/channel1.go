package main

import (
	"fmt"
	"sync"
)

func isPrime(x int) bool {
	if x < 2 {
		return false
	}
	if x == 2 {
		return true
	}
	if x%2 == 0 {
		return false
	}
	for j := 3; j*j <= x; j += 2 {
		if x%j == 0 {
			return false
		}
	}
	return true
}

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		fmt.Printf("worker %d started working..on job %d\n", id, job)
		if isPrime(job) {
			results <- job
		}
	}

}

func main() {
	fmt.Println("fan-in/fan-out 1 producer sending to 5 goroutines..")
	jobs := make(chan int, 1000)
	results := make(chan int, 1000)
	//1 producer adds 2-1000000 to channel
	go func() {
		for i := 2; i <= 100000; i++ {
			jobs <- i
		}
		close(jobs)
	}()
	var wg sync.WaitGroup

	for w := 1; w <= 1000; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}
	go func() {
		wg.Wait()
		close(results)
	}()
	cnt := 0
	for range results {
		cnt++
	}
	fmt.Println(cnt)

}
