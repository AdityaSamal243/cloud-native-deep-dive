//acquiring ownership through receiving values from a channel

package main

import (
	"time"
	"log"
	"math/rand"
)

func ServeCustomer(bar24x7 chan int, consumer chan int){
    for c := range consumer{
		log.Print("customer#",c,"enters the bar")
		seat := <-bar24x7   //need a seat to drink
		log.Print("++ Customer#",c,"drinks at seat#",seat)
		time.Sleep(time.Second*time.Duration(2+rand.Intn(400)))
		log.Print("-- Customer",c,"free seat#",seat)
		bar24x7 <- seat;  // free the seat and leave bar
	}	

}

func main(){
	bar24x7 := make(chan int,10)
	//place seats in a bar 
	for i:=0 ; i < cap(bar24x7) ; i++{
		bar24x7 <- i;  //None of the send will block 
	}
	//after the above loop if we try to send it will block/sleep waiting as buffer is full

	consumer := make(chan int,10)
	for i:=0 ; i < cap(bar24x7) ; i++{
		go ServeCustomer(bar24x7,consumer)
	}

	for customerid := 0; ; customerid++{
		time.Sleep(1*time.Second)
		consumer <- customerid
	}

	for{time.Sleep(time.Second)}

}