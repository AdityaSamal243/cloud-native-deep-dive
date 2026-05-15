package main

import "fmt"

func main(){
	const pi = 3.14
	const name = "aditan"
	// name = "ass" // This will cause a compile-time error because constants cannot be reassigned
	fmt.Println(pi)
	fmt.Println(name)
	// to assign multiple constants at once
	const(
		country = "india"
		state = "odisha"
	)
	fmt.Println(state,",",country)
}