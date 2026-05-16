package main

// used for iterating over arrays, slices, maps, strings and channels

import "fmt"
func main(){
    nums := []int{6,7,8}
    
	// for i:=0;i<len(nums);i++{
	// 	fmt.Println(nums[i])
	// }
    sum :=0
	for i, num :=range nums{   //what is _ here ? it is the index
		fmt.Println(num,i)
		sum+=num
	}
	fmt.Println(sum)


	// to iterate over map

	m :=map[string]string{}
	m["name"]="aditan"
	m["age"]="thirty"
	m["city"]="delhi"

	for k, v := range m{
		fmt.Println(k,v)
	}
	for k :=range m{
		fmt.Println(k)
	}


	// to iterate over string

	s1 :="aditan"
	for i, ch := range s1{
		// fmt.Println(i ,ch)    //it gives ascii values of characters
		fmt.Println(i ,string(ch)) // to get character from ascii value

		// unicode code point rune
		// starting byte of rune
		// ascii contains 255 --> each character takes 1 byte space 
		// if anything greater than 255 --> like if 300 it takes 2 byte so index will skip by 1
	}
}