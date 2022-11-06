package client

//  The net module lets you make network connections and transmit data.

import (
	"fmt"
	"net"
	"truco/app/common"
)

const (
	SERVER_HOST = "localhost"
	SERVER_PORT = "9988"
	SERVER_TYPE = "tcp"
)

func Start() {

	// connect to server
	socket, err := net.Dial(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
	if err != nil {
		fmt.Println("Fail connect to server")
		return
	}
	runClient(socket)
	fmt.Println("SALI DEL RUNCLIENT")
}

func runClient(socket net.Conn) {
	fmt.Println("entre a client run")
	sendMenuResponses(socket)
	//este receive deberia bloquearse esperando a que empiece la partida.
	messageServer, _ := common.Receive(socket)
	fmt.Println("Message server: ", messageServer)
}
