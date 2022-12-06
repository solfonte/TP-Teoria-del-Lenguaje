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
	SERVER_PORT = "9951"
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

func contains_message(message string) bool {
	specificMessages := []string{"Espera a que juegue tu oponente...",
		"Tu oponente tiro una carta",
		"Estas son tus cartas",
		"Tu oponente ",
		"Ganaste", "Perdiste", "Cantaste ",
		"Tiraste la carta", "Aceptaste",
		"Rechazaste", "Tus puntos son",
		"Te fuiste al MAZO",
		common.FinishGame, common.WinMatchMessage,
		common.LoseMatchMessage}
	find := false
	for _, msg := range specificMessages {
		if strings.Contains(message, msg) {
			find = true
			break
		}
	}
	return find
}

func processGameloop(socket net.Conn) {
	// loop de server manda algo cliente responde
	reader := bufio.NewReader(os.Stdin)
	ch := make(chan string)
	go ProcessResponseClient(ch, reader)
	for {
		fmt.Println("voy a procesar")
		messageServer, err := common.Receive(socket)
		checkErrorServer(err)
		fmt.Println(messageServer)
		if contains_message(messageServer) {
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
						//fmt.Println("Read input from stdin:", stdin)
						if strings.TrimSpace(stdin) == QUIT {
							fmt.Println("entre a quit")
							return
						}
						//fmt.Println("lo que mando ", stdin)
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
