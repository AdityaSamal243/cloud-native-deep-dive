package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Data division worker starting...")
	
	input := make(chan int, 20)
	for i := 0; i < 20; i++ {
		input <- i
	}
	close(input) // 🛠️ CRITICAL: Close input so the divisor knows when work is done!

	output1 := make(chan int, 5)
	output2 := make(chan int, 5)
	output3 := make(chan int, 5)
	output4 := make(chan int, 5)

	// Launch the optimized divisor
	go divisor(input, output1, output2, output3, output4)

	// 🛠️ FIXING MAIN: To read multiple active streams without blocking sequentially,
	// use a WaitGroup to read from all outputs concurrently.
	var wg sync.WaitGroup
	outputs := []chan int{output1, output2, output3, output4}

	for i, out := range outputs {
		wg.Add(1)
		go func(id int, ch chan int) {
			defer wg.Done()
			for val := range ch {
				fmt.Printf("[Output %d] Received: %d\n", id+1, val)
			}
		}(i, out)
	}

	wg.Wait()
	fmt.Println("All data divided and processed cleanly!")
}

func divisor(input <-chan int, outputs ...chan<- int) {
	fmt.Println("Started data division module...")
	
	// 🛠️ Ensure all output channels are closed when input runs dry
	defer func() {
		for _, out := range outputs {
			close(out)
		}
	}()

	index := 0
	numOutputs := len(outputs)

	// Continuously pull items from input until it is closed
	for val := range input {
		// Evenly distribute values in a Round-Robin rotation (0, 1, 2, 3, 0, 1...)
		outputs[index] <- val
		index = (index + 1) % numOutputs
	}
}