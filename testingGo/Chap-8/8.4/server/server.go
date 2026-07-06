// modify server to use to sync.waitGroup per connection
//to count number of active echo goroutines. when it falls to 0.
//close the write half of the tcp connection.
//modified client wait for final echoes of multiple concurrent shouts even after input is closed
// refer 8.3

//task1: a simple client and server that can handle multiple connections. the client sends a line of text to the server and the server eachoes it back to client.
//task2: tcp connection has two halves closeRead and closeWrite methods.
// modify main goroutine of client to close only the writehalf of the connection so that program will be able to print
//

package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8001")
	fmt.Println("server started listeneing on port 8001")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		fmt.Println("connection accepted from client")
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConn(conn)
	}

}

func handleConn(c net.Conn) {
	var wg sync.WaitGroup
	defer c.Close()              // to gracefully close  the connection from server when the function returns
	input := bufio.NewScanner(c) // receives input from client via bufio newScanner package
	for input.Scan() {
		// Process each line of input
		wg.Add(1)
		go echo(c, input.Text(), 5*time.Second, &wg)
	}
	wg.Wait()
	if tcpConn, ok := c.(*net.TCPConn); ok {
		tcpConn.CloseWrite()
		fmt.Println("server write closed")
	}
}

func echo(c net.Conn, shout string, delay time.Duration, wg *sync.WaitGroup) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
	wg.Done()
}
