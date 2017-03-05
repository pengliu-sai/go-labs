package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	l, err := net.Listen("tcp", "localhost:9000")
	if err != nil {
		fmt.Println("Error listening:", err)
		os.Exit(1)
	}
	defer l.Close()
	fmt.Println("Listening on " + l.Addr().String())

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err)
			os.Exit(1)
		}

		//logs an incoming message
		fmt.Printf("Received message %s -> %s \n", conn.RemoteAddr(), conn.LocalAddr())

		// Handle connections in a new goroutine.
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	defer conn.Close()
	for {
		io.Copy(conn, conn)
	}
}
