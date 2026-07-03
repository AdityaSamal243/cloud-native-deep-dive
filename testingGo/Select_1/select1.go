// implement fan-in — merge 5 goroutine output channels into one using select

package main

import (
	"fmt"
	"sync"
	"time"
)

func sender(ch chan<- int, id int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		fmt.Printf("worker %d started\n", id)
		ch <- id*10 + i	
		time.Sleep(time.Second * 100)
	}
}
func fanIn(ch1, ch2, ch3, ch4, ch5 chan int) <-chan int {
	fmt.Println("fanin started")
	result := make(chan int, 25)
	go func() {
		defer close(result)
		remaining := 5
		for remaining > 0 {
			select {
			case v,ok := <-ch1:
				if !ok {
					remaining--
					continue
				}
				result <- v
			case v,ok := <-ch2:
				if !ok {
					remaining--
					continue
				}
				result <- v
			case v,ok := <-ch3:
				if !ok {
					remaining--
					continue
				}
				result <- v
			case v,ok := <-ch4:
				if !ok {
					remaining--
					continue
				}
				result <- v
			case v,ok := <-ch5:
				if !ok {
					remaining--
					continue
				}
				result <- v
			case <-time.After(time.Second * 200):
				fmt.Println("timeout")
				return
			}
		}
	}()
	return result
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)
	ch4 := make(chan int)
	ch5 := make(chan int)

	var wg sync.WaitGroup
	wg.Add(5)
	
	go sender(ch1, 1, &wg)
	go sender(ch2, 2, &wg)
	go sender(ch3, 3, &wg)
	go sender(ch4, 4, &wg)
	go sender(ch5, 5, &wg)

	go func() {
		wg.Wait()
		close(ch1)
		close(ch2)
		close(ch3)
		close(ch4)
		close(ch5)

	}()

	
	mergedfanIn := fanIn(ch1, ch2, ch3, ch4, ch5)
	for v := range mergedfanIn {
		fmt.Println(v)
	}

}
