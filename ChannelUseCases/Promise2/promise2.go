// passing send-only channel ( chan <- int) as a parameter to a function

package main

import(
	"fmt"
	"time"
	"math/rand"
)


func LongTimeRequest( r chan<- int32){
     time.Sleep(time.Second * 5) // simulation of workload
	 r <- rand.Int31n(100)   // send-only channel so we can only send data to the channel
}

func SumOfSquares(a, b int32) int32{
	return a*a + b*b
}

func main(){
	ra := make(chan int32)
	rb := make(chan int32)

	go LongTimeRequest(ra)   // concurrently both are called
	go LongTimeRequest(rb)  // send-only channel is passed as a parameter to the function
    
	fmt.Println("calc started...")

	// after 5 seconds, the values will be received.
	fmt.Println(SumOfSquares(<-ra , <-rb))  // receive values from the channels and pass them to the function
}