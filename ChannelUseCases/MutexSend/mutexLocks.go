//lock on send 

package main

import (
	"fmt"
)

func main(){
    
	counter := 0
	mutex := make(chan struct{},1)
	increase := func(){
		mutex <- struct{}{} //lock
        counter++
		defer func(){ <- mutex }() // unlock
	}
	increase100 := func(done chan <- struct{}){
       for i:=0;i<10000;i++{
		increase()
	   }
	   done <-struct{}{}

	}

	done:=make(chan struct{})
	go increase100(done)
	go increase100(done)
	<-done
	<-done 
	fmt.Println(counter)

}