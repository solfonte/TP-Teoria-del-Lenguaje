package client

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"truco/app/common"
)

func sendPlayerName(socket net.Conn) {
	reader := bufio.NewReader(os.Stdin)
	messageServer, _ := common.Receive(socket)
	fmt.Println(messageServer)
	messageClient, _ := reader.ReadString('\n')
	common.Send(socket, messageClient)
}

func processMenuOptions(socket net.Conn, messageServer string) {
	reader := bufio.NewReader(os.Stdin)
	for !strings.HasPrefix(messageServer, "OK") {
		messageClient, _ := reader.ReadString('\n')
		common.Send(socket, messageClient)
		messageServer, _ = common.Receive(socket)
		fmt.Println(messageServer)
	}
	fmt.Println("buscando o creando partida")
}

func sendMenuResponses(socket net.Conn) {
	sendPlayerName(socket)
	i := 0
	messageServer := ""
	for i < 2 {
		messageServer, _ = common.Receive(socket)
		fmt.Println(messageServer)
		common.Send(socket, "Ok")
		i++
	}
	messageServer, _ = common.Receive(socket)
	fmt.Println(messageServer)
	processMenuOptions(socket, messageServer)
}

func startGame(socket net.Conn) {
	i := 0
	// estas son tus cartas
	for i < 2 {
		messageServer, _ := common.Receive(socket)
		fmt.Println("Message server: ", messageServer)
		common.Send(socket, "Ok")
		i++
	}
	

}
