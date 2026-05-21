package main
import(
	"fmt"
	"time"
)

// struct helps in OOP 

type Order1 struct{
	id string
	amount float32
	status string
	createdAt time.Time
}
// we create a function below for creation
// it is kind of constructor

func newOrder(id string,amount float32, status string) *Order1 {
	// initial setup
	    myOrder1 := Order1{
		id : id,
		amount : amount,
		status : status,
	  }
	return &myOrder1
}



// In OOP we attach methods to classes same we can do in struct

//a method to change status

// reciever type == this function now will be attached to struct
//pass tsruct as reference
func (O *Order1)changeStatus(status string){

  // why didn't we dereference here ?
  // Struct does the derefrencing work for us
	O.status = status

}

func (O *Order1) getAmount() float32 {
	return O.amount
}


func main(){
      myOrder := Order1{
		id : "1",
		amount : 100,
		// status : "paid",
		createdAt : time.Now(),
	  }
	  // see above instance status is commented and it is a string 
	  //so if we print myOrder the status will show empty string
	  // int = 0, float =0 , string = empty string , bool = false
	  myOrder.changeStatus("confirmed")
	  fmt.Println(myOrder)
	  fmt.Println(myOrder.getAmount())

	  myOrder2 :=Order1{
		id : "2",
		amount : 200,
		status : "delivered",
		createdAt : time.Now(),
	  }
	  myOrder.status = "returned"
	  fmt.Println("order2 :",myOrder2)
	  fmt.Println("order1 :",myOrder)


	  myOrder1 := newOrder("3",300,"paid")
	  fmt.Println("order1 :",myOrder1)


	  //but what if struct instance is used only once in a program no need to do so much 
      
	language := struct{
	  name string
	  isGood bool
	}{"Go",true}
	fmt.Println(language)


}