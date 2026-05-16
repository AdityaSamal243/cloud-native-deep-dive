package main

import (
	"fmt"
	"maps"
)

func main(){
	fmt.Println("maps in go")
	// maps are key-value pairs
	// declaration

	m :=make(map[string]int) // map of string keys and int values
	m["adi"]=1
	m["aditan"]=2
	fmt.Println(m)
	fmt.Println(m["adi"], m["aditan"])
	fmt.Println(len(m))

	//what if we access a key that does not exist
	// key,value does not exist in map then it return 0.
	fmt.Println(m["unknown"]) // returns zero value of int which is 0

	m1 :=make(map[string]string);
	m1["age"]="thirty"
	fmt.Println(m1)
	fmt.Println(m1["lol"])  //return empty string lol not a key


	// to delete element from map
	delete(m,"adi");
	fmt.Println(m)

	// clear the map
	clear(m1)
	fmt.Println(m1)


	// declare map without make

	m2 := map[string]int{}
	m2["a"]=1
	m2["b"]=2
	m2["c"]=3
	m2["d"]=4
	fmt.Println(m2)

	// to check if element in map there or not

	_, ok := m2["a"]
	fmt.Println(ok) // returns true

	if ok{
	    fmt.Println("present in map")
	}else{
         fmt.Println("not pesent")
	}

	m3 := map[string]int{}
	m3["a"]=1
	m3["b"]=2
	m3["c"]=3
	m3["e"]=4

	//check if m2 and m3 are equal
	// fmt.Println(m2 == m3) // maps cannot be compared directly
	fmt.Println(maps.Equal(m2,m3)) 
	// if any key/value is different then it returns false
    




	
	
}