package main

import "fmt"

// i want to recieve unlimited parameters.
func sum(nums ...int) int {
	total :=0
	for _,num := range nums{
		total+=num
	}
	return total
}
func main(){
	fmt.Println(1,2,3,4,5,"hello") // variadic function can take variable number of arguments
    fmt.Println(sum(1,2,3,4,5)) // variadic function can take variable number of arguments

	nums := []int{1,2,3,4,5}
	fmt.Println(sum(nums...)) // to pass slice to variadic function we need to use ... after slice name
}