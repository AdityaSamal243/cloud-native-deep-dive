package main

// important
//makes code scalable
//for example : project to implement payment gateway 
// 1st step : a struct ds is required
// 2nd step --  we need a payment gateway suppose razorpay integrate
// 1 more struct is required 


import "fmt"


type paymenter interface {
	pay(amount float32) //works like contract
	// we can list more methods here like refund
} 


type Payment struct{
	// gateway stripe
	gateway paymenter
}

// Open Close principle is violated

func (p *Payment) makePayment(amount float32){
        // razorpaypaymentGw := razorpay{}
		// razorpaypaymentGw.pay(amount)

		// stripepaymentgw := stripe{}
		// stripepaymentgw.pay(amount)

		p.gateway.pay(amount) 
}

type razorpay struct{}

func (r razorpay) pay(amount float32){
	//logic to make payment razorpay apis
	fmt.Println("making payment using razorpay",amount)
}

// the above is concrete implmentation of razorpay payment gateway


// type stripe struct{}

// func (s stripe) pay(amount float32){
// 	//logic to make payment razorpay apis
// 	fmt.Println("making payment using stripe",amount)
// }


//for unit testing purposes

type fakepay struct{}

func (f fakepay) pay(amount float32){
	//logic to make payment razorpay apis
	fmt.Println("making payment using fakepay for testing",amount)
}

// if in future new payment gateway ios to be added

type Paypal struct{}
// use same method signature present in paymenter interface

func (p Paypal) pay(amount float32){
	fmt.Println("making payment from Paypal",amount)
}



func main(){
    // stripePaymentGw := stripe{}
	// fakepayGw := fakepay{}
	// razorpayGw := razorpay{}
	paypalGw := Paypal{}
	newPayment := Payment{
		gateway: paypalGw,
	}
	newPayment.makePayment(100)

}

// when new requirement comes then problem arises 
// suppose instead of razorpay we will use stripe.
//  or add new payment gateways 

// we did using stripe but we can see we need to modify the code by commenting the razorpay part 
// it will make tetsing also difficult
//so for this purpose we use interfaces

//interface is a contract


// in go we don't need to explicitely mention extends interface
// if methods within an interface is present in any struct then the go compiler undertands
// the struct implements that interface
// it happens implicitely (same method signature must be there i.e name arguemnts etc)