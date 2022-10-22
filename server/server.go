package main

import (
	"fmt"
	"net"
	"os"
)

const (
	HOST = "localhost"
	PORT = "9988"
	TYPE = "tcp"
)

func main() {
	fmt.Println("Server Running...")

	server, error := net.Listen(TYPE, HOST+":"+PORT)
	// there are no exceptions we handle error with if (hay panic tambien investigar)
	if error != nil {
		fmt.Println("Error listening:", error.Error())
		os.Exit(1)
	}
	// delay the execution of the function or method or an anonymous method until the nearby functions returns.
	defer server.Close()
	fmt.Println("Listening on " + HOST + ":" + PORT)
	fmt.Println("Waiting for client...")
	for {
		// acept diferent connections
		conn, error := server.Accept()
		if error != nil {
			fmt.Println("Error accepting: ", error.Error())
			os.Exit(1)
		}
		fmt.Println("client connected")
		go processClient(conn)
	}

}
func processClient(connection net.Conn) {
	buffer := make([]byte, 1024)
	// leo lo que me llega
	dataLen, error := connection.Read(buffer)
	if error != nil {
		fmt.Println("Error reading:", error.Error())
		connection.Close()
	}
	fmt.Println("Received: ", string(buffer[:dataLen]))
	_, error = connection.Write([]byte("Llego este mensaje al server: " + string(buffer[:dataLen])))
	connection.Close()
}
