// implement try-send mechanism in ticker

//time package has its own pre defined tick functiom 




package main	

import(
	"fmt"
	"time"
)

// func Tick(d time.Duration) <- chan struct{}{ //  return a recieving channel.
// 	c := make(chan struct{},1)  // a buffered channel with size =1
// 	go func(){
//          for{
// 			time.Sleep(d)
// 			select{
// 			case c <- struct{}{}:
// 			default:
// 			}
// 		 }
// 	}()
// 	return c
// }

// func main(){
//  t := time.Now()
// 	for range Tick(time.Second){
// 		fmt.Println(time.Since(t))
// 	}
// }

func main(){
    ticker := time.Tick(2*time.Second)  //creates a channel that ticks every 2 second
    fmt.Println("starting periodic task...")
	for now := range ticker{
        fmt.Println(now.Format("15:14:05"))

	}
}

//"The underlying Ticker cannot be recovered by the garbage collector; it leaks."
// no way to call .Stop() on it.
// remain alive and consume memory resources for the entire lifetime of your application.

//Never use time.Tick inside spawned short-lived goroutines, since if goroutine exits. the time.Tick remains active for entire application lifetime 
//HTTP request handlers, or methods that are created and destroyed dynamically.

//safe alternative : time.NewTicker