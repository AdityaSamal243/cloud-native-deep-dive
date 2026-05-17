package main

// slices are dynamic arrays no need to give length before hand
// most used construct in go
// useful methods are there in slices

import (
	"fmt"
	"slices"
)

func main(){
    //   declaration 
	var nums []int // slice of integers
    fmt.Println(nums)  // empty and nil slice 
    fmt.Println(len(nums)) // length is 0 
	
	// 3rd parameter is capacity (optional)
	var arr = make([]int,2) // make is used to create a slice with length 2
    fmt.Println(arr)
	fmt.Println(len(arr)) // length is 2
	fmt.Println(cap(arr)) // capacity is 2
	// capacity --> maximum elemnts that can fit
	
	// add elemnts
	arr = append(arr,3) 
	// append is used to add elements to a slice
	fmt.Println(arr)
	fmt.Println(len(arr))
	fmt.Println(cap(arr)) 

	// can add multiple elements at once
	arr = append(arr,4,5,6)
	fmt.Println(arr)
	fmt.Println(len(arr))
	fmt.Println(cap(arr))

    // to initialize an empty slice 

	var empslice = make([]int,0)
	fmt.Println(empslice)

	// to create a slice
	slice := []int{}
	fmt.Println(slice)
    fmt.Println(len(slice))
	slice = append(slice,1,2,3)
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))


	// to add element via index

	arr1 := make([]int,2)
	arr1[0]=2
	arr1[1]=3
	fmt.Println(arr1)
	fmt.Println(len(arr1))

    //  copy function
    // if destination slice is smaller than source slice then only part of the source slice will be copied to destination slice
	// if dest slice is 0 nothing will be copied
	var nums1 = make([]int,0,5)
	nums1 = append(nums1, 2)
	var nums2 = make([]int,len(nums1))

	fmt.Println(nums1, nums2)

	copy(nums2, nums1) // copy(dst,src)
	fmt.Println(nums1, nums2)


	// slice operator used for getting element or portion of element
	slice1 := []int{1,2,3,4,5}
	fmt.Println(slice1[0:5])
	fmt.Println(slice1[:3])
	fmt.Println(slice1[1:3])
	fmt.Println(slice1[1:])


	// slice build package

	arr3 := []int{1,2}
	arr4 := []int{1,2}

	fmt.Println(slices.Equal(arr3,arr4))
	fmt.Println(slices.Concat(arr3,arr4))


	// 2D slicess

	var mulnums = [][]int{{1,2,3,4},{5,6,7,8}}
	fmt.Println(mulnums)




}