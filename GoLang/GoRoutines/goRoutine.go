package main

import "fmt"
//goroutines ?
//they are lightweight threads (multithreading) to run process concurrently

func task(id int){
     fmt.Println("doing task",id)
}

func main(){
    // for i:=0;i<=10;i++{
	// 	task(i) //it occurs in order
	// }

	// i want the task function to run in  parallel
	// just add go
	 for i:=0;i<=10;i++{
		go task(i) // all process runs concurrently
	}
	
}

//cpu intensive work can be done using goroutine