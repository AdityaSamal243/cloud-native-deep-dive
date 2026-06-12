package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func main(){
	listner,err := net.Listen("tcp","localhost:8000")
	fmt.Println("server started at port 8000")
	if err!=nil{
		log.Fatal(err)
	}
	for{
		conn,err := listner.Accept()
		fmt.Println("server started listening to client")
		if err !=nil{
			log.Fatal(err)
			continue
		}
		go handleconn(conn)  //multiple client can get response instead of waiting for first proccess to finish
	}

}
func handleconn(c net.Conn){
      input := bufio.NewScanner(c)
	  for input.Scan(){ //what does this line do ?  / reads one line at a time
		go echo(c,input.Text(),1*time.Second)
	  }
	  c.Close() // ignore erros from input.Err()
}

func echo(c net.Conn, shout string, delay time.Duration){
	fmt.Fprintln(c,"\t",strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c,"\t",shout)
	time.Sleep(delay)
	fmt.Fprintln(c,"\t",strings.ToLower(shout))
	
}


//what is exactly happening under the hood?
//1. server starts running on port 8000
//2. accepts connection from client
//3. recieves input from client via bufio newScanner package 
//4. created a function that takes this input as arguement
//5. the function echo logic is to send output to client in 3 formats.
