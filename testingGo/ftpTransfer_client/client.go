package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	// Connect straight to your running FTP server socket
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()

	// FIX: Use io.Copy to pipe raw network bytes directly to the terminal screen instantly
	go func() {
		_, err := io.Copy(os.Stdout, conn)
		if err != nil {
			return
		}
	}()

	fmt.Println("Connected! Type your commands (ls, cd <dir>, get <file>):")
	inputScanner := bufio.NewScanner(os.Stdin)
	for inputScanner.Scan() {
		text := inputScanner.Text()
		if text == "exit" {
			break
		}
		// Send the command down the wire with the network-standard line break
		fmt.Fprintf(conn, "%s\r\n", text)
	}
}