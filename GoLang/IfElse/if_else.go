package main

import "fmt"
func main(){
	age :=12
	if age >=18{
		fmt.Println("You are an adult")
	} else {
		fmt.Println("You are a minor")
	}

	var role = "admin"
	var hasPerm = false

	if role == "admin" && hasPerm {
		fmt.Println("You have admin privileges")
	}else {
		fmt.Println("no admin privileges")
	}

    // we can directly declare and initialize a variable in the if statement
	if name := "aditan"; name == "adi" {
		fmt.Println("Hello, adi!")
	}else{
		fmt.Println("Hello, stranger!")
	}
	// go does not have ternary operator, we can use if-else instead
}