package main

import(
	"fmt"
)

func PrintSlice(items []int){
	for _,item :=range items{
		fmt.Println(item)
	}
}

// what to do if i have string slice I have to create a new function that takes slice of string type

func PrintSliceString(items []string){
	for _,item :=range items{
		fmt.Println(item)
	}
}
//we can see PrintSliceString and PrintSlice are exact same logic 
// and duplicates only name and arguements are different
//so to solve this problem we use generics

//now suppose in future i have to rpint boolean type slice then again i have to create a new function


func PrtSlice[T int | string | bool, V string](items []T, name V){    //any means it can be of any type
	for _,item :=range items{
		fmt.Println(item, name)
	}
}  //no need to duplicate
//but any is not a good practice as we can add any type here
//i want it to be scoped i.e only integers and string type are allowed
// we can also pass multiple types you can see passed T and V string 


//we can use generics in struct also

type Stack[T any] struct{
	// elements []int
	//use of generics
	//instead of any....we can also pipe the types
	elements []T
	
}

func main(){
	nums :=[]int{1,2,3,4,5}
    PrtSlice(nums, "john")

	name :=[]string{"aditan","tan"}
	PrtSlice(name,"john")

	judge :=[]bool{true,false}
	PrtSlice(judge,"john") // to make it valid pipe bool type in PrtSlice function

	myStack := Stack[int]{
		elements: []int{1,2,3,4,5},
	}
	myStack2 := Stack[string]{
		elements: []string{"go","k8s"},
	}	
	fmt.Println(myStack)
	fmt.Println(myStack2)
}