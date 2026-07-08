// try-send: if case has send operation. then select block is called try-send.
// if the send operation is blocked(cannot send into channel) . then default branch is executed

package main

import(
	"fmt"
)
type Book struct{
	id int
}

func main(){
	bookshelf := make(chan Book,3)

	//try-send: we are tring sending into channel
	for i:=0;i<cap(bookshelf)*2;i++{
		select{
		case bookshelf <- Book{id : i}:
           fmt.Println("succeeded to put book",i)
		default:
			fmt.Println("failed to put book")		 
		}
	}
	//try-receive
	for i :=0; i < cap(bookshelf)*2 ; i++{
		select{
		case book := <-bookshelf:
			fmt.Println("succeeded to get book",book.id)
		default:
			fmt.Println("failed to get book")
		}
	}
}