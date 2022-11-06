package main

//  The net module lets you make network connections and transmit data.

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
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
		fmt.Println("Fail connect to server")
		return
	}
	runClient(socket)
	fmt.Println("SALI DEL RUNCLIENT")
}

func runClient(socket net.Conn) {
	fmt.Println("entre a client run")
	sendMatchParameters(socket)
	//este receive deberia bloquearse esperando a que empiece la partida.
	messageServer, _ := common.Receive(socket)
	fmt.Println("Message server: ", messageServer)
}

func sendMatchParameters(socket net.Conn) {
	reader := bufio.NewReader(os.Stdin)
	// pido nombre
	messageServer, _ := common.Receive(socket)
	fmt.Println("Message server: ", messageServer)
	messageClient, _ := reader.ReadString('\n')
	common.Send(socket, messageClient)
	// bienvenida
	messageServer, _ = common.Receive(socket)
	fmt.Println("Message server: ", messageServer)
	common.Send(socket, "ok")
	messageServer, _ = common.Receive(socket)
	fmt.Println("Message server: ", messageServer)
	common.Send(socket, "ok") //TODO: esto hay que sacarlo porque lo pusimos
	messageServer, _ = common.Receive(socket)
	fmt.Println("Message server: ", messageServer)

	// responde el cliente

	for !strings.HasPrefix(messageServer, "OK") {
		messageClient, _ = reader.ReadString('\n')
		common.Send(socket, messageClient)
		messageServer, _ = common.Receive(socket)
		fmt.Println("Message server: ", messageServer)
	}
	// consultar
	// se creo partida o se esta buscando partida

	for !strings.HasPrefix(messageServer, "Seleccione") {
		messageServer, _ = common.Receive(socket)
		fmt.Println("Message server: ", messageServer)
	}

	messageClient, _ = reader.ReadString('\n')

	common.Send(socket, messageClient)

	if messageClient == "1" {
		for !strings.HasPrefix(messageServer, "Seleccione") {
			messageServer, _ = common.Receive(socket)
			fmt.Println("Message server: ", messageServer)
		}

		messageClient, _ = reader.ReadString('\n')

		common.Send(socket, messageClient)
	}
}
