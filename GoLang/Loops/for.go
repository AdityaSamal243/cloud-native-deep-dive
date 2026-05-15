package main

import "fmt"

// for is the only construct in go for looping
func main(){
    // while loop
	i :=1
	for i <=3 {
		fmt.Println(i);
		i++
	}

	// infinite loop
	// for{
	// 	fmt.Println("infinite loop")
	// }

	// classic for loop
	for j :=0; j<3 ;j++{
		if j==1{
			break
		}
		fmt.Println(j)
	}


	// range
	for k := range 11{
		fmt.Println(k)	
	}
}