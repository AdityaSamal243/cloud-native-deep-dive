package main

import (
	"log"
	"time"
)

func main() {
	ready, done := make(chan struct{}), make(chan struct{})
	go worker(0, ready, done)
	go worker(1, ready, done)
	go worker(2, ready, done)

	time.Sleep(time.Second * 10)
	//1-to-N notififcations -- done is unbuffered so main sends token 1 by 1 to 3 workers so 1 sends value to N

	// ready <- struct{}{}
	// ready <- struct{}{}
	// ready <- struct{}{}

	//or 

	close(ready)

	//N to 1 notifications---- Notfications from different sources into 1,.
	<-done
	<-done
	<-done
}

func worker(id int, ready <-chan struct{}, done chan<- struct{}) {
	<-ready // blocked here and wait for notification
	log.Print("worker", id, "starts..")
	time.Sleep(time.Second * time.Duration(id+1))
	log.Print("worker", id, "job done..")
	// notify main goroutine work is finished
	done <- struct{}{}

}
