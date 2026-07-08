//acquiring ownership through receiving values from a channel

package main

import (
	"time"
	"log"
	"math/rand"
)

func ServeCustomer(bar24x7 chan int, c int, seat int){
	log.Print("++ Customer#",c,"drinks at seat#",seat)
	time.Sleep(time.Second*time.Duration(2+rand.Intn(400)))
	log.Print("-- Customer",c,"free seat#",seat)
	bar24x7 <- seat;  // free the seat and leave bar

}

func main(){
	bar24x7 := make(chan int,10)
	//place seats in a bar 
	for i:=0 ; i < cap(bar24x7) ; i++{
		bar24x7 <- i;  //None of the send will block 
	}
	//after the above loop if we try to send it will block/sleep waiting as buffer is full

	for customerid := 0; ; customerid++{
		time.Sleep(2*time.Second)
		seat := <-bar24x7
		go ServeCustomer(bar24x7,customerid,seat)
	}

	for{time.Sleep(time.Second)}

}