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
		// can be used for host,port and env variables as they are constant
	)
	fmt.Println(state,",",country)
}