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
		done <- struct{}{}
	}()
    
	//recieve is blocked untill someone sendd so it could receieve
	<- done 
	fmt.Println("world")
}