//CLI takes N as input → spawns N workers → each worker processes
//a "job" (simulate with time.Sleep + random int) → fan-in to results → print ordered output → timeout after 5s → graceful shutdown on SIGINT

package main

import (
	"fmt"
	"sync"
	"os"
	"os/signal"
	"time"
	"math/rand"
)
type Result struct{
	jobId int
	value int 
}
//approach 1
func main() {
	n := 20
	input := make(chan int,20)
	output := make(chan Result,20)
	done := make(chan struct{})
	var wg sync.WaitGroup

    // signal + timeout -- closes done

	sigs := make( chan os.Signal,1)
	signal.Notify(sigs,os.Interrupt)
	timeout := time.After(5*time.Second)

	go func(){
		select{
		case <- sigs:
			fmt.Println("shut down")
			close(done)
		case <-timeout:
			fmt.Println("timeout")
			close(done)
		}
	}()

	for i := 1; i <= n; i++ {
        wg.Add(1)
		go Worker(i,input,output,&wg,done)
	}
	for i :=1; i<=20; i++{
		input <- i
	}
	close(input)


	
	
}

func Worker(id int,input chan int, output chan <- Result , wg *sync.WaitGroup,done <- chan struct{}){
	defer wg.Done()
	for{
		select{
		case <- done:
			return
		case id,ok := <-input:
			if !ok{
				return
			}
			time.Sleep(2*time.Second)
			output <- Result{jobId: id, value: rand.Intn(100)}
		}
	}
  
}
