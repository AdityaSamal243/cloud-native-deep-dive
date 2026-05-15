package main

import "fmt"
import "time"

func main(){
	// switch statement is used to perform different actions based on many multiple conditions
	// simple switch 
	i :=7

	switch i {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
	case 4:
		fmt.Println("4")
	case 5:
		fmt.Println("5")
	default:
		fmt.Println("other")
	}

	// multiple condittion switch

	switch time.Now().Weekday(){
	case time.Saturday, time.Sunday:
		fmt.Println("weekend")
	default:
		fmt.Println("workday")
	}

	// type switch
	whoAmI := func(i interface{}){
		// it returns the type of variable
		// switch t := i.(type){
		switch i.(type){
		case int:
			fmt.Println("it is an integer")
		case string:
			fmt.Println("it is string")
		case bool:
			fmt.Println("it is a boolean")
		default:
			fmt.Printf("unknown type %T\n", i)  
		}
	}
	whoAmI("aditan")
	whoAmI(42)
	whoAmI(true)
	whoAmI(3.14)
}