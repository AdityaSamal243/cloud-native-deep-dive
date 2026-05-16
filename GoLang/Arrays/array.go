package main

import "fmt"

// array numbered sequence of specific length  (same type)
func main(){
	// int ==0 , string ==empty , bool == false
	// by default all elements are 0 when declared
   var nums [4]int // array of 4 integers
   nums[0] = 1
   nums[1] = 2

//    to get element at index 0
   fmt.Println(nums[0]) 
   fmt.Println(nums[1])
   fmt.Println(nums)

   fmt.Println(len(nums))

   var vals [4]bool;
   vals[2] = true
//    by default all elements are false
   fmt.Println(vals)

   var name [3]string;
//    by default all elements are empty string
   name[0] = "aditan"
   fmt.Println(name);

//    how to add elements when we declare an array
    arr:= [3]int{1,2,3}
	fmt.Println(arr)

	// ---------------------------------------------------
	// 2d-arrys
	// ------------------------------------------------
	mularr := [2][2]int{{23,24},{25,26}}
	fmt.Println(mularr)

	// in golang instead of arrays mostly slices are used
	// use array when you know the size...memory optimized , constant time access 
	// but if size not given use slices === dynamic 
}