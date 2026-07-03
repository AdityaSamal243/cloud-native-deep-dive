//returning receive-only channel (<-chan int) as a result from a function

package main

import (
	"math/rand"
	"fmt"
	"time"
)

func LongTimeRequest() <-chan int32 {  // returns a receive-only channel as a result
       r := make(chan int32)  //unbuffered channel
	   go func(){
		  // simulation of a workload 
		  time.Sleep(time.Second * 5)
		  r <- rand.Int31n(100)  //send a random number to the channel
	   }()
	   return r
}

func sumOfSquares(a, b int32) int32{
	return a*a + b*b
}

func main(){
     a,b := LongTimeRequest(), LongTimeRequest()  
	 fmt.Println("calc started...")
     fmt.Println(sumOfSquares(<-a , <-b))  //receive values from the channels and pass them to the function
}
