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
	done := make(chan struct{}) //channel struct type is used just for signal purpose

	go func(){
		io.Copy(os.Stdout,conn)
		log.Println("done")
		done <- struct{}{} // empty signal to main go routine
	}()
   
	mustCopy(conn,os.Stdin)

	tcpConn,ok := conn.(*net.TCPConn)
	if ok{
		tcpConn.CloseWrite()
	}else{
		conn.Close()
	}
	<- done 
	//after signal reciewd from backgound it terminates
    
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