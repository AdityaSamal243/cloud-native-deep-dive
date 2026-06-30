package main

import (
	"fmt"
	"time"
)

func main() {
	cakeBuffer := make(chan string, 1)
	go func() { //baker
		for i := 1; i <= 5; i++ {
			cakeName := fmt.Sprintf("cake %d", i)
			fmt.Printf("[baker] started baking %s...\n", cakeName)
			time.Sleep(200 * time.Millisecond)
			fmt.Printf("[Baker] %s...placing on counter\n", cakeName)

			cakeBuffer <- cakeName

			//current state

			fmt.Printf("counter status: %d/%d cakes waiting\n", len(cakeBuffer), cap(cakeBuffer))

		}
		close(cakeBuffer)
		fmt.Println("shop closed")
	}()

	for cake := range cakeBuffer {
		fmt.Println("[Icer] Picked up from counter", cake)
		time.Sleep(190000 * time.Millisecond)
		fmt.Printf("[Icer] finised icing %s..\n", cake)
	}

}
