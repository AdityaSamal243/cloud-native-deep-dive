package main

import (
	"flag"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	//port no. must be user defined
	port := flag.String("port", "8000", "port number")

	flag.Parse() //it parse the flags from command line

	listener, err := net.Listen("tcp", ":"+*port) //listen on user defined port
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept() //program is in blocked state untill tcp handshake is done (nc localhost:8000)
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn) //The handleConn function handles one complete client connection
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n")) //io.WriteString writes over tcp to the client
		if err != nil {
			return //client disconnected
		}
		time.Sleep(time.Second)
	}
}
