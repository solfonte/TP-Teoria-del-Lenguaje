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
	SERVER_PORT = "9994"
	SERVER_TYPE = "tcp"
)

func Start() {

	// connect to server
	socket, err := net.Dial(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
	if err != nil {
		fmt.Println("Fail connect to server")
		return
	}
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("", err)

		}
		socket.Close()
	}()
	runClient(socket)
}

func runClient(socket net.Conn) {
	sendMenuResponses(socket)
	//este receive deberia bloquearse esperando a que empiece la partida.
	startGame(socket)
	//loop juego
	processGameloop(socket)
}

func processGameloop(socket net.Conn) {
	// loop de server manda algo cliente responde
	promptReader := bufio.NewReader(os.Stdin)
	// falta desconexion tanto cliente como servidor
	//handlear ctrl c
	for {
		messageServer, err := common.Receive(socket)
		checkErrorServer(err)
		fmt.Println(messageServer)
		if strings.Contains(messageServer, "Espera a que juegue tu oponente...") {
			common.Send(socket, "OK")
		} else if strings.Contains(messageServer, "Tu oponente tiro una carta") {
			common.Send(socket, "OK")
		} else if strings.Contains(messageServer, "Ganaste") || strings.Contains(messageServer, "Perdiste") {
			common.Send(socket, "OK")
		} else if strings.Contains(messageServer, "Estas son tus cartas") {
			common.Send(socket, "OK")
		} else if strings.Contains(messageServer, "Tu oponente se desconecto") {
			common.Send(socket, "OK")
		} else if strings.Contains(messageServer, "Ganaste el envido") || strings.Contains(messageServer, "Perdiste el envido") {
			common.Send(socket, "OK")
		} else if strings.Contains(messageServer, "cantaste ") {
			common.Send(socket, "OK")
		} else if strings.Contains(messageServer, "Tu oponente") {
			common.Send(socket, "OK")
		} else if strings.Contains(messageServer, "aceptaste") || strings.Contains(messageServer, "Rechazaste") {
			common.Send(socket, "OK")
		} else if strings.Contains(messageServer, "Bienvenido") {
			common.Send(socket, "OK")
		} else {
			messageClient, _ := promptReader.ReadString('\n')
			common.Send(socket, messageClient)
		}
	}
}
