//requirements
// create ftp server listens on port 8000  so that client can establish connection
//the server should be able to interpret commands cd,ls,get
//can use standard ftp command line client 

//ftp(allows remote file mgmt) is application layer protocol - transfer files between client and server
//1st step is to create a client server connection using tcp

//control connection(port 21)
//data connection(port 20) - used for file transfer

//it must be isolated 

package main

import(
	"fmt"
	"net"
	"os"
	"io"
	"bufio"
	"strings"
)

// a struct to represent ftp server also client will have private session.
// why this so no session overwrites the same pwd. suppose a did cd and b did ls at same time then both will overwrite same pwd
type FTPServer struct{
     conn net.Conn
	 pwd string
}

func main(){
	conn,err := net.Listen("tcp","localhost:8000")  //opens the tcp at port 8000  
	if err !=nil{
		fmt.Println(err)
		return
	}
	for{
		c,err := conn.Accept() // waits for client connection
		if err !=nil{
			fmt.Println(err)
			continue
		}
		myFTP := FTPServer{
			conn: c,
			pwd: os.Getenv("HOME"), //initially set to home directory
		}
		go handleConn(&myFTP) //handleConn will handle the client connection and interpret the commands
	}
}
// i need to keep track of client conn and the pwd for use of cd,ls,get cmd .

func handleConn(ftp *FTPServer){
	defer ftp.conn.Close()
	io.WriteString(ftp.conn,"220 Welcome to FTP Server\r\n") //220 is the status code for service ready
    
	scanner := bufio.NewScanner(ftp.conn)  //reads the client input in 8000 port
	for scanner.Scan(){
		input := scanner.Text()
		fields := strings.Fields(input)
		if len(fields) == 0{
			continue
		}
		cmd := fields[0]

		switch cmd{
		case "ls":
			

		}
	}
	
}