package server

import (
	"fmt"
	"net"
	"os"
)

type Receiver struct {
	listenerSocket net.Listener
}

func start_receiver(receiver Receiver) {
	for {
		// acept diferent connections
		conn, error := receiver.listenerSocket.Accept()
		if error != nil {
			fmt.Println("Error accepting: ", error.Error())
			os.Exit(1)
		}
		fmt.Println("client connected")
		go process_client(conn)
	}

}

func process_client(connection net.Conn) {
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
