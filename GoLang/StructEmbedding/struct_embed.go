package main


import (
	"fmt"
	"time"
)

//struct embedding is a way to reuse the fields and methods of one struct in another struct.
// it is a way to achieve composition in go.
// it is a way to achieve code reuse in go.
// it is a way to achieve inheritance in go.

type Customer struct{
	name string
    email string
}

type Order struct{
	id string
	amount float32
	status string
	createdAt time.Time
	Customer // embedding customer struct in order struct
}

func main(){
     
	 Customer1 := Customer{
		name : "John Doe",
		email : "jdoe@gmail.com",
	 }
	 myOrder := Order{
		id : "1",
		amount : 100,
		status : "paid",
		createdAt : time.Now(),
		Customer: Customer1,
	 }
	 fmt.Println("Order Struct:",myOrder)
	 fmt.Println("customer is:",myOrder.Customer)

	 myOrder2 := Order{
		id : "2",
		amount : 200,
        status : "shipped",
		createdAt: time.Now(),
		Customer : Customer{
			name : "tanu",
			email : "tan@gmail.com",
		},
	 }
	 fmt.Println(myOrder2)

	 myOrder2.Customer.name = "aditya"

	 fmt.Println(myOrder2.Customer)

     
}