// peak limit is used to limit number of concurrent request
// without blocking any request.

package main

import (
	"log"
	"math/rand"
	"time"
)


func main(){
	bar24x7 := make(chan int,10);
	for i := 0; i < cap(bar24x7) ; i++{
		//fill the seats 
		bar24x7 <- i;
	}
    //bar has 10 seats //no of worker =11
	consumer := make(chan int,10)
	for i := 0 ; i< 11 ; i++{
		go serve(bar24x7, consumer)
	}

	for custid := 0; ;custid++{
		time.Sleep(2*time.Second)
		consumer <- custid
	}
}

func serve(bar24x7 chan int, consumer chan int){
	for c := range consumer{
        log.Print("customer#", c, "enters the bar")
		select{
		case seat := <-bar24x7:
			log.Print("++ Customer#", c, "drinks at seat#", seat)
			time.Sleep(time.Second * time.Duration(2+rand.Intn(400)))
			log.Print("-- Customer", c, "free seat#", seat)
			bar24x7 <- seat // free the seat and leave bar
		default:
			log.Print("all seats taken")
		}
	} 
}