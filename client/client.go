package client

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
	startGame(socket)
	//loop juego
	processGameloop(socket)
}

func processGameloop(socket net.Conn) {
	// loop de server manda algo cliente responde
	promptReader := bufio.NewReader(os.Stdin)
	for {
		messageServer, _ := common.Receive(socket)
		fmt.Println(messageServer)
		if strings.HasPrefix(messageServer, "Espera a que juegue tu oponente...") {
			fmt.Println("entre a esperar")
			common.Send(socket, "OK")
		} else if strings.Contains(messageServer, "Tu oponente tiro una carta") {
			fmt.Println("Tu oponente tiro una carta")
			common.Send(socket, "OK")
		} else if strings.Contains(messageServer, "la jugada") {
			fmt.Println("gannaste o perdiste")
			common.Send(socket, "OK")
		} else {
			messageClient, _ := promptReader.ReadString('\n')
			common.Send(socket, messageClient)
		}
	}
}
