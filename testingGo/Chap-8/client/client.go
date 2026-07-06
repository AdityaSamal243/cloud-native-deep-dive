package main	

import (
	"fmt"
	"log"
	"io"
	"net"
	"os"
)

func main(){
	listener,err := net.Dial("tcp","localhost:8001")
	if err!= nil{
		log.Fatal(err)
		return
	}
	fmt.Println("connected to server")
	defer listener.Close()
	done :=  make(chan struct{}) // created a channel to send signal that echo has been done
	go func(){  // background goroutine sends the text from server to client
		io.Copy(os.Stdout,listener)
		fmt.Println("done")
		done <- struct{}{}
	}()
	mustCopy(listener,os.Stdin) // main goroutine: copy from keyboard to server
	//ctrl+d blocks the terminal

	fmt.Println("standard input closed. shutting down client write half")

	tcpconn,ok :=listener.(*net.TCPConn)
	if ok{
		fmt.Println("ruk ja bhai")
		tcpconn.CloseWrite() //tells the server no more input is going
	}else{
		fmt.Println("not a tcp connection")
	}

	
	<-done

	


}

func mustCopy(dst io.Writer,src io.Reader){
	 _,err := io.Copy(dst,src)
	 if err !=nil{
		log.Fatal(err)
	 }
}