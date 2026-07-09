// random number generator

//it is a multi return future/promise 
//data producer can close any time to end data generating.

package main	

import (
	"fmt"
	"math/rand"
	"time"
)

func randomGenerator() <- chan int{
  fmt.Println("started generating...")
  c := make(chan int)
  go func(){
     for{
		num := rand.Intn(400)
		time.Sleep(2*time.Second)
		c <- num;
	 }
  }()
  return c
}

func main(){
	fmt.Println("starting random numer generation.....")
	values := randomGenerator()
    
	// for val := range values{
	// 	fmt.Println(val);
	// }
    for{
		time.Sleep(3*time.Second)
        select{
		case x := <-values:
			fmt.Println(x)
		default:
			fmt.Println("no space in buffer")
	}
	}
	
}