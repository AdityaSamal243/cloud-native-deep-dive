package main

import(
	"fmt"
	"time"
)

func main(){
	done := make(chan struct{})
    
	go func(){
		fmt.Println("Hello")
		time.Sleep(time.Second*10)
		<- done
	}()
    
	// send operation blocked/sleep untill someone ready to receive
	done <- struct{}{}
	fmt.Println("world")
}