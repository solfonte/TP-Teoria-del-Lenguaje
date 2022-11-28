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
	SERVER_PORT = "9977"
	SERVER_TYPE = "tcp"
	QUIT        = "Q"
)

func Start() {

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

	startGame(socket)

	processGameloop(socket)
}

func processGameloop(socket net.Conn) {
	// loop de server manda algo cliente responde
	promptReader := bufio.NewReader(os.Stdin)

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
		} else if strings.Contains(messageServer, "Cantaste ") {
			common.Send(socket, "OK")
		} else if strings.Contains(messageServer, "Tu oponente") {
			common.Send(socket, "OK")
		} else if strings.Contains(messageServer, "Aceptaste") || strings.Contains(messageServer, "Rechazaste") {
			common.Send(socket, "OK")
		} else if strings.Contains(messageServer, "Tus puntos son") {
			common.Send(socket, "OK")
		} else if strings.Contains(messageServer, common.FinishGame) {
			fmt.Println("termino el juego")
			common.Send(socket, "OK")
			return
		} else {
			messageClient, _ := promptReader.ReadString('\n')
			if strings.TrimSpace(messageClient) == QUIT {
				fmt.Println("entre a quit")
				return
			}
			common.Send(socket, messageClient)
		}
	}
}
