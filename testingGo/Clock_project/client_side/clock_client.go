package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("no arguements passed")
	}
	args := os.Args[1:]

	for _, arg := range args {
		value := strings.Split(arg, "=")
		label := value[0]
		host := value[1]

		go connectAndStream(label, host) //for server connection
	}
	select {}

}

func connectAndStream(label, host string) {
	conn, err := net.Dial("tcp", host)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	fmt.Printf("------connected to %s-----------\n", label)
	_, err = io.Copy(os.Stdout, conn) //io.Copy copies data from server to client
	if err != nil {
		log.Fatal(err)
	}
}
