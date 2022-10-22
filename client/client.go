package main

//  The net module lets you make network connections and transmit data.

import (
	"fmt"
	"net"
)

const (
	SERVER_HOST = "localhost"
	SERVER_PORT = "9988"
	SERVER_TYPE = "tcp"
)

func main() {

	// connect to server
	conn, err := net.Dial(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
	if err != nil {
		panic(err)
	}

	// sending some data to server
	_, err = conn.Write([]byte("Este es un mensaje de prueba"))
	// buffer where the data that the server sendas will be store
	buffer := make([]byte, 1024)
	// read what the server send
	dataLen, error := conn.Read(buffer)
	if error != nil {
		fmt.Println("Error reading:", error.Error())
	}
	fmt.Println("Received: ", string(buffer[:dataLen]))
	// close connection
	defer conn.Close()
}
