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
	SERVER_HOST  = "localhost"
	SERVER_PORT  = "9973"
	SERVER_TYPE  = "tcp"
	QUIT         = "Q"
	TIME_TO_WAIT = 6
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
		common.FinishGame,
		common.WinMatchMessage,
		common.LoseMatchMessage}
	find := false
	for _, msg := range specificMessages {
		if strings.Contains(message, msg) {
			if !strings.Contains(message, common.WaitingOptionsPlayer) && !strings.Contains(message, "Es tu turno,") {
				find = true
				break
			}
		}
	}
	return find
}

func processGameloop(socket net.Conn) {

	reader := bufio.NewReader(os.Stdin)
	ch := make(chan string)
	go ProcessResponseClient(ch, reader)
	for {
		messageServer, err := common.Receive(socket)
		checkErrorServer(err)
		fmt.Println(messageServer)
		if contains_message(messageServer) {
			common.Send(socket, "OK")
			if strings.Contains(messageServer, common.FinishGame) {
				return
			}
		} else {
			finish := false
			for !finish {
				select {
				case stdin, ok := <-ch:
					if !ok {
						break
					} else {
						if strings.TrimSpace(stdin) == QUIT {
							return
						}
						common.Send(socket, stdin)
						finish = true
					}
				case <-time.After(TIME_TO_WAIT * time.Second):
					if strings.Contains(messageServer, "Mientras esperas a que sea tu turno") {
						common.Send(socket, common.ACK)
						finish = true
					}
				}
			}
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
		ch <- s
	}
}
