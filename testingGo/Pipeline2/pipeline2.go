package main

import (
	"fmt"
)

func counter(out chan<- int) {
	for i := 1; i <= 5; i++ {
		out <- i
	}
	close(out)
}
func square(out chan<- int, in <-chan int) {
	for v := range in {
		out <- v * v
	}
	close(out)

}
func printer(in <-chan int) {
	for x := range in {
		fmt.Println(x)
	}
}

func main() {
	naturals := make(chan int)
	squarer := make(chan int)
	go counter(naturals)         //counter ka kaam hai send krna data to channel
	go square(squarer, naturals) //squarer function has both revieving type channel and sending type channel.
	printer(squarer)          //printer takes recieving channel

}
