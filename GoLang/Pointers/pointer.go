package main

import "fmt"

// pointer = address , this address stores the location of variable
// passed by value = copy of variable is passed to function
func changeNum(num int){
	num = 5
	fmt.Println("in chnageNum:",num)
}
// passed by reference = address of variable is passed to function
func pointerChangeNum(num *int){
	*num = 5 //it derefrences address
	fmt.Println("in pointerChangeNum:",*num)
}


func main(){

	num :=1
	changeNum(num)
	fmt.Println("in main:",num) // num is not changed because we are passing a copy of num to changeNum function

    num1 :=2
	fmt.Println("memory address of num1:",&num1) //gives address of num1
	pointerChangeNum(&num1) // we are passing the address of num1 to pointerChangeNum function
	fmt.Println("in main:",num1) // num1 is changed because we are passing address of num1 to pointerChangeNum function
}