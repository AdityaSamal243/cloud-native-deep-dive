//suppose there are various sources, response duration vary for each, send request to each source via separate gorutine, 
// use the 1st response and ignore rest.


package main	

import (
	"fmt"
	"time"
	"math/rand"
)

func source(c chan <- int32){  // passing send only channel as parameter
     ra,rb := rand.Int31n(100), rand.Intn(3)+1
	 time.Sleep(time.Second * time.Duration(rb))  //simulate the response duration
	 c <- ra  // send the response to the channel
}

func main(){
	//suppose 5 sources
    startTime := time.Now()
	c := make(chan int32, 5)
	for i:=0;i<cap(c);i++{
		go source(c)
	}
	// recieve the fist res 
	rnd := <-c
	fmt.Println(time.Since(startTime))
	fmt.Println(rnd)
}