package main

import (
	"fmt"
	"sync"
)

func worker(id int,jobs <-chan int, results chan<-int, wg *sync.WaitGroup){
	defer wg.Done()
	for job := range jobs{
             fmt.Printf("worker %d started working..on job %d\n",id, job)
			 x := job;
			 count :=0;
			 for j :=1;j<=x;j++{
				if x%j==0{
					count++
				}
			 }
			 if count==2{
                results <- x;
			 }
	}

}

func main(){
	fmt.Println("fan-in/fan-out 1 producer sending to 5 goroutines..")
	jobs := make(chan int,5)
	results := make(chan int,5)
	//1 producer adds 2-1000000 to channel
	go func(){
		for i :=2;i<=100000;i++{
		    jobs <- i
	    }
		close(jobs)
	}()
    var wg sync.WaitGroup;

	for w:=1;w<=5;w++{
		wg.Add(1)
		go worker(w,jobs,results,&wg)
	}
	go func() {
		wg.Wait()
		close(results)
	}()
    cnt :=0;
	for range results{
        cnt++;
	}
	fmt.Println(cnt)

}