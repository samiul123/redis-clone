package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	fmt.Println("Listening on port :6379")

	// Create a new server
	l, err := net.Listen("tcp", ":6379")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Listen for connections
	fmt.Println("Accepting connections")
	conn, err := l.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}

	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Printf("Error closing connection %v. Message: %v\n", conn.RemoteAddr(), err.Error())
		}
	}(conn)

	for {
		buf := make([]byte, 1024)

		_, err = conn.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("Error reading from client: ", err.Error())
			os.Exit(1)
		}

		// ignore request and send back a PONG
		_, err := conn.Write([]byte("+OK\r\n"))
		if err != nil {
			fmt.Println("Error writing to client: ", err.Error())
			os.Exit(1)
		}
	}
}
