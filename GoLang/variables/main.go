package main

import "fmt"

func main() {
	// var name string = "tanu"
	// var age int = 22


	// go can automatically infer the type of variable based on the value assigned to it
	// var name, age = "tan", 62

	// short hand syntax
	// name := "tan"
	// age := 62

	// but there is a need to use complete syntax if variable is declared without assignment
	var name string
	var age int
    var price float64
	name = "tan"
	age = 62
	price = 5.333333

	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(price)
}