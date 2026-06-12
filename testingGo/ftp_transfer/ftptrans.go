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

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath"
	"strings"
)

// a struct to represent ftp server also client will have private session.
// why this so no session overwrites the same pwd. suppose a did cd and b did ls at same time then both will overwrite same pwd
type FTPServer struct {
	conn net.Conn
	pwd  string
}

func main() {
	conn, err := net.Listen("tcp", "localhost:8000") //opens the tcp at port 8000
	fmt.Println("server running on port 8000") 
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		c, err := conn.Accept() // waits for client connection
		if err != nil {
			fmt.Println(err)
			continue
		}
		dir, err := os.Getwd()
		if err != nil {
			dir = "." // Fallback if system call fails
		}
		myFTP := FTPServer{
			conn: c,
			pwd:  dir, //initially set to /cloud-native-ddep-dive
		}
		go handleConn(&myFTP) //handleConn will handle the client connection and interpret the commands
	}
}

// i need to keep track of client conn and the pwd for use of cd,ls,get cmd .

func handleConn(ftp *FTPServer) {
	defer ftp.conn.Close()
	io.WriteString(ftp.conn, "220 Welcome to FTP Server\r\n") //220 is the status code for service ready

	scanner := bufio.NewScanner(ftp.conn) //reads the client input in 8000 port
	for scanner.Scan() {
		input := scanner.Text()
		input = strings.TrimSpace(input)
		fields := strings.Fields(input)
		if len(fields) == 0 {
			continue
		}
		cmd := fields[0]

		switch cmd {
		case "ls":
			files, err := os.ReadDir(ftp.pwd)
			if err != nil {
				io.WriteString(ftp.conn, "550 failed to read directory\r\n")
				continue
			}
			for _, file := range files {
				io.WriteString(ftp.conn, file.Name()+"\r\n")
			}
		case "cd":
			// cd /home/ubuntu fields=["cd","/home/ubuntu"]
			if len(fields) < 2 {
				io.WriteString(ftp.conn, "550 no such directory")
				continue
			}
			targetdir := filepath.Join(ftp.pwd, fields[1])

			info, err := os.Stat(targetdir) //validation of directory
			if err != nil {
				io.WriteString(ftp.conn, "550 no such directory\r\n")
				continue
			}
			//check if it is a dirctory or file

			if info.IsDir() {
				ftp.pwd = targetdir
				io.WriteString(ftp.conn, "into the directory.")
			} else {
				io.WriteString(ftp.conn, "550 target path is a file")
			}

		case "get":
			//1st check if fields <2
			if len(fields) < 2 {
				io.WriteString(ftp.conn, "501 missing file\r\n")
				continue
			}
			//joins file to pwd
			targetdir := filepath.Join(ftp.pwd, fields[1])
			res, err := os.Stat(targetdir)
			if err != nil {
				io.WriteString(ftp.conn, "550 no such file\r\n")
				continue
			}

			if res.IsDir() {
				io.WriteString(ftp.conn, "550 target path is directory\r\n")
				continue
			}
			//download over internet
			//continuous token so we will use io package
			file, err := os.Open(targetdir)
			if err != nil {
				io.WriteString(ftp.conn, "failed to open\r\n")
				continue
			}
			// stream the file directly accross netwrok connection wire
			_, err = io.Copy(ftp.conn, file)
			file.Close() // always cleans up resource leaks

			if err != nil {
				fmt.Println("error streaming data:", err)
				io.WriteString(ftp.conn, "failed to get\r\n")
				continue
			}

		}
	}

}
