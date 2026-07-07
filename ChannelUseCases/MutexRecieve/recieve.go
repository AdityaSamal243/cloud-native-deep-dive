// lock on recieve

package main

import "fmt"

func main(){
	
	counter:=0
	mutex := make(chan struct{})
	mutex <- struct{}{}
	
	increase :=func(){
		<- mutex //lock 
		counter++
		mutex <- struct{}{}
	}

	increase100 := func(done chan <- struct{}){
		for i:=0; i<10000;i++{
			increase()
		}
		done <- struct{}{}
	}

	done := make(chan struct{})
	go increase100(done)
	go increase100(done)

	<-done
	<-done

	fmt.Println(counter)
}

//dry run

// a and b 2 jan hai same time pe i=0 liya 
//suppose a pehle exectute kiye <- mutex
//now mutex empty hogaya
//jab b execute kiya <-mutex nothing to reciece toh blocked hogaya 
// jaise hi a counter ++ kiya and mutex <- struct{}{} bheja
//again active and b revieve krke same process continues