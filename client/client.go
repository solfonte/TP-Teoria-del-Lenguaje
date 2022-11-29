package client

//  The net module lets you make network connections and transmit data.

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
	"truco/app/common"
)

const (
	SERVER_HOST = "localhost"
	SERVER_PORT = "9963"
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
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("voy a procesar")
		messageServer, err := common.Receive(socket)
		checkErrorServer(err)
		fmt.Println(messageServer)
		ch := make(chan string)
		go ProcessResponseClient(ch, reader)
		if strings.Contains(messageServer, "Espera a que juegue tu oponente...") {
			fmt.Println("entre a aca")
			common.Send(socket, "OK")
		} else if strings.Contains(messageServer, "Tu oponente tiro una carta") {
			common.Send(socket, "OK")
		} else if strings.Contains(messageServer, "Ganaste") || strings.Contains(messageServer, "Perdiste") {
			common.Send(socket, "OK")
		} else if strings.Contains(messageServer, "Estas son tus cartas") {
			fmt.Println("Mando un ok")
			common.Send(socket, "OK")
		} else if strings.Contains(messageServer, "Tu oponente se desconecto") {
			common.Send(socket, "OK")
		} else if strings.Contains(messageServer, "Ganaste el envido") || strings.Contains(messageServer, "Perdiste el envido") {
			common.Send(socket, "OK")
		} else if strings.Contains(messageServer, "Cantaste ") || strings.Contains(messageServer, "Tiraste la carta") {
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
		} else if strings.Contains(messageServer, common.WinMatchMessage) || strings.Contains(messageServer, common.LoseMatchMessage) {
			common.Send(socket, "OK")
		} else {
			finish := false
			for !finish {
				select {
				case stdin, ok := <-ch:
					fmt.Println("lo que mando ", stdin)
					if !ok {
						break
					} else {
						fmt.Println("Read input from stdin:", stdin)

						if strings.TrimSpace(stdin) == QUIT {
							fmt.Println("entre a quit")
							return
						}
						fmt.Println("lo que mando ", stdin)
						common.Send(socket, stdin)
						finish = true
					}
				case <-time.After(6 * time.Second):
					if strings.Contains(messageServer, "Mientras esperas a que sea tu turno") {
						common.Send(socket, "0")
						finish = true
					}

				}

			}
			fmt.Println("sali de procesar")
		}
	}
}

func ProcessResponseClient(ch chan string, reader *bufio.Reader) {
	for {
		s, err := reader.ReadString('\n')
		if err != nil {
			close(ch)
			return
		}
		fmt.Println("recibo: ", s)
		ch <- s
	}
}
