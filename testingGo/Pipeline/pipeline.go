package main

import (
	"fmt"
)

func main(){
	naturals := make(chan int)
	squarer := make(chan int)

	//counter
	go func(){
		for x := 0;x<=3 ; x++{
             naturals <- x
		}
		close(naturals) // tell the squarer we are done sending.
	}()

	//squarer
	go func(){
		for x:= range naturals{  //automatically stops when natural closes
			squarer <- x*x
		}
		close(squarer)
	}()

	//printer
	for x :=range squarer{
		fmt.Println(x)
	}
}