package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	fmt.Println("connected to server")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close() // to handle the termination smoothly
   
	go mustCopy(conn,os.Stdin)
	mustCopy(os.Stdout,conn)
    
}
func mustCopy(dst io.Writer, src io.Reader){
     _,err := io.Copy(dst,src)
	 if err !=nil{
		log.Fatal(err)
	 }
}


//core idea
//1. logic to connect to server
//2. write the logic to send string to server