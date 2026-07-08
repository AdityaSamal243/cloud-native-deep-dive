// we will use select statement to implement first reponse first wins
// use 5 background goroutines/worker
// the wroker who send response 1st wins.

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func source(test chan <- int32){
	ra,rb := rand.Int31(), rand.Intn(3)+1
	time.Sleep(time.Second * time.Duration(rb))
	select{
	case test <- ra:
	default:
	}
}

func main(){
    
	test := make(chan int32,1) // buffered channel of size 1
	for i :=0 ; i<5 ;i++{
		go source(test)
	}

	rnd := <- test // it is blocked here as soon as elemnt comes in buffer. it receives and return
	fmt.Println("the first response is=",rnd)
}