package main

import (
	"fmt"
	"time"
)
// using struct we do OOP in go
// struct in go is similiar to classes and objects in other programming languages.
// to create a complex data type we use struct
// order struct
// the below is blueprint from this we could create multiple struct 
type Order struct {
	id string
	amount float32
	status string
	createdAt time.Time  // nanosecond precision
}


func main(){
	//instnace of order struct
	// ord := Order{
       var myord Order = Order{
		id: "1",
		amount: 100,
		status: "delivered",
		
	   }
	   myOrd2 := Order{
		id: "2",
		amount: 5000,
		status: "shipped",
		createdAt: time.Now(),
	   }
	   fmt.Println("Order Struct:",myord)
	   fmt.Println("Order Struct 2:",myOrd2)

	   //    now if i want to add an field after creation of instance

	   myord.createdAt = time.Now()
	   fmt.Println("Order Struct with createdAt:",myord)
       
	   // to get a field
	   fmt.Println("Order Id:",myord.id)
}