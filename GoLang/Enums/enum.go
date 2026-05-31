package main

import (
	"fmt"
)

//enum = enumerated types
//real world examples
//enum is implemented using const in go
//in go we can create custom type
//for enum we use custom type

type OrderStatus int

type Orderstat string

const(
	Recieved Orderstat = "received"
	Confirmed = "confirmed"
	Prespared = "prepared"
	Delivered = "delivered"

)

// const(
// 	Received OrderStatus = iota //iota is predeclared identifier  representing untyped integer //increment hote rahega
//     Confirmed
// 	Prepared
// 	Shipped
// 	Delivered
// )

//if the above is present in any other file/package we can import this 
//as well as we can import the function
// for string we can't remember all the status 
//but using enums we can hover and just go to the type i.e orderstatus and check the values present

// before above implementation
func changeOrderStatus(status string ){
	fmt.Println("changing order status to",status)
}

//after above implementation 

func ChangeStatus(status Orderstat){
	fmt.Println("changing order status to", status)
}

func main(){
	// changeOrderStatus("recieved")
	// changeOrderStatus("confirmed")

	ChangeStatus(Recieved)
	ChangeStatus(Delivered)

    //1st problems
	//it is not a problem
	//but there is a problem, suppose  we have multiple status
	//and in real world projects many such calls are made 
	// and if there is a typo while passing status this could be a problem

	//order status are related valiues so they must be grouped 







}