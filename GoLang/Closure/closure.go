package main

import "fmt"

func counter() func() int{
	count := 0

	return func() int{
		count++
		return count
	}
}
// if a variable of outer scope is used 
// in inner function then it is called closure. 
// in this example count is a variable of outer scope and 
// it is used in inner function which is returned 
// by counter function. so counter function returns a closure 
// which can access and modify the count variable even 
// after 
// counter function has finished executing.

func main(){

	c1 := counter() // we will store the returned function in c1
	fmt.Println(c1()) // 1
	fmt.Println(c1()) // 2
	fmt.Println(c1()) // 3
}