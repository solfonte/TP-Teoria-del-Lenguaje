package main

//  The net module lets you make network connections and transmit data.

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"truco/app/common"
)

const (
	SERVER_HOST = "localhost"
	SERVER_PORT = "9988"
	SERVER_TYPE = "tcp"
)

func main() {

	// connect to server
	socket, err := net.Dial(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
	if err != nil {
		fmt.Println("error")
		panic(err)
	}

	// // sending some data to server
	// _, err = conn.Write([]byte("Este es un mensaje de prueba"))
	// // buffer where the data that the server sendas will be store
	// buffer := make([]byte, 1024)
	// // read what the server send
	// dataLen, error := conn.Read(buffer)
	// if error != nil {
	// 	fmt.Println("Error reading:", error.Error())
	// }
	// fmt.Println("Received: ", string(buffer[:dataLen]))
	// // close connection
	fmt.Println("entre a client run")
	reader := bufio.NewReader(os.Stdin)
	for {
		messageServer, error := common.Receive(socket)
		if error != nil {
			fmt.Println("Error reciving from server: ", error.Error())
			socket.Close()
		}
		fmt.Println("Message server: ", messageServer)
		messageClient, _ := reader.ReadString('\n')
		if messageClient == "close" {
			fmt.Println("Stop playing")
			return
		}
		common.Send(socket, messageClient)
	}
}

// func runClient(socket net.Conn) {
// 	fmt.Println("entre a client run")
// 	reader := bufio.NewReader(os.Stdin)
// 	for {
// 		messageServer, error := common.Receive(socket)
// 		if error != nil {
// 			fmt.Println("Error reciving from server: ", error.Error())
// 			socket.Close()
// 		}
// 		fmt.Println("Message server: ", messageServer)
// 		messageClient, _ := reader.ReadString('\n')
// 		if messageClient == "close" {
// 			fmt.Println("Stop playing")
// 			return
// 		}
// 		common.Send(socket, messageClient)
// 	}
// }
