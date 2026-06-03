package main

import(
	"io"
	"log"
	"net"
	"time"
)


func main(){
	//listen on port 8000
	listener,err := net.Listen("tcp",":8000")
	if err !=nil{
		log.Fatal(err)
	}
	for{
		conn,err := listener.Accept()  //program is in blocked state untill tcp handshake is done (nc localhost:8000)
		if err!=nil{
			log.Print(err)
			continue
		}
		go handleConn(conn) //The handleConn function handles one complete client connection
	}
}


func handleConn(c net.Conn){
	defer c.Close()  // it means if program exited, network error , defer ensures the socket is closed efficiently to avoid memory leaks
	for{
		_,err :=io.WriteString(c, time.Now().Format("15:04:05\n"))   //io.WriteString writes over tcp to the client
		if err!=nil{
			return
		}
		time.Sleep(time.Second)
  }
  //loop ends when write fails=== client disconnects
  //handdleConn closses the connection when client disconnects using defer
  //and wait for the next client to connect.
}