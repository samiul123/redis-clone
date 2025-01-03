package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	fmt.Println("Listening on port :6379")

	// Create a new server
	l, err := net.Listen("tcp", ":6379")
	if err != nil {
		fmt.Println(err)
		return
	}

	aof, err := NewAof("database.aof")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer func() {
		if err := aof.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	aof.Read(func(value Value) {
		command := strings.ToUpper(value.array[0].bulk)
		args := value.array[1:]

		handler, ok := handlers[command]
		if !ok {
			fmt.Println("Invalid command: ", command)
			return
		}

		handler(args)
	})

	// Listen for connections
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}

		go handleConnection(conn, aof)
	}
}

func handleConnection(conn net.Conn, aof *Aof) {
	defer func() {
		if err := conn.Close(); err != nil {
			fmt.Println("Error closing connection:", err)
		}
	}()

	for {
		resp := NewResp(conn)
		value, err := resp.Read()
		if err != nil {
			fmt.Println("Read error: ", err)
			return
		}

		if value.dataType != "array" {
			fmt.Println("Invalid request, expected array")
			continue
		}

		if len(value.array) == 0 {
			fmt.Println("Invalid request, expected array length > 0")
			continue
		}

		command := strings.ToUpper(value.array[0].bulk)
		args := value.array[1:]

		// ignore request and send back a PONG
		writer := NewWriter(conn)
		handler, ok := handlers[command]

		if !ok {
			fmt.Println("Invalid command: ", command)
			writer.Write(Value{dataType: "string", str: ""})
			continue
		}

		if command == "SET" || command == "HSET" {
			aof.Write(value)
		}

		result := handler(args)
		writer.Write(result)
	}
}
