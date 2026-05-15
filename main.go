package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("Hello prosperity!!!. Listening on port :6379")

	//Start listening on port :6376
	l, err := net.Listen("tcp", ":6379")
	if err != nil {
		fmt.Println(err)
		return
	}

	//Accept the connection
	conn, err := l.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}

	//Defer here executes the function conn.close()
	// at the end after all the other functions are done executing
	//, this is to ensure that the connection is closed after all
	// the other functions are done executing. And its written
	// right after opening the connection for readability.
	defer conn.Close()

}
