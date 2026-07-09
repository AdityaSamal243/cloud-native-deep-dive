//data aggregation..recieves from many send into one

package main

import (
	"fmt"
	"sync"
)

func dataAggregator(inputs...<-chan int) <- chan int{
      output := make(chan int,10)
	  var wg sync.WaitGroup
	  fmt.Println("started....")
	  for _,in := range inputs{
		wg.Add(1)
          go func(in <- chan int){
			defer wg.Done()
            for val := range in{
				output <- val
			}
		  }(in)
	  }
	  go func(){
		wg.Wait()
		close(output)
	  }()
	  return output
}


func main(){
	
	fmt.Println("started processing multiple stream into one....")
	input1 := make(chan int,10)
	input2 := make(chan int, 10)
	input3 := make(chan int, 10)
	input4 := make(chan int, 10)
	input5 := make(chan int, 10)

	for i:=1;i<=10;i++{
		input1 <- i
	}
	for i:=11;i<=20;i++{
		input2 <- i
	}
	for i:=21;i<=30;i++{
		input3 <- i
	}
	for i:=31;i<=40;i++{
		input4 <- i
	}
	for i:=41;i<=50;i++{
		input5 <- i
	}
	close(input1)
	close(input2)
	close(input3)
	close(input4)
	close(input5)

	result := dataAggregator(input1,input2,input3,input4,input5)

	for res := range result{
		fmt.Println(res)
	}
	
}