package client

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"truco/app/common"
)

const (
	PING_MESSAGE = "Listo para jugar?"
)

func checkErrorServer(err error) {
	if err != nil {
		panic("error in server, Close connection")
	}
}

func sendPlayerName(socket net.Conn) {
	reader := bufio.NewReader(os.Stdin)
	messageServer, err := common.Receive(socket)
	checkErrorServer(err)
	fmt.Println(messageServer)
	messageClient, _ := reader.ReadString('\n')
	common.Send(socket, messageClient)
}

func processMenuOptions(socket net.Conn, messageServer string) {
	reader := bufio.NewReader(os.Stdin)
	for {
		messageClient, _ := reader.ReadString('\n')
		common.Send(socket, messageClient)
		messageServer, err := common.Receive(socket)
		checkErrorServer(err)
		fmt.Println(messageServer)
		if strings.Contains(messageServer, "OK") {
			return
		}

	}

}

func sendMenuResponses(socket net.Conn) {
	i := 0
	messageServer := ""
	for i < 2 {
		messageServer, err := common.Receive(socket)
		checkErrorServer(err)
		fmt.Println(messageServer)
		common.Send(socket, "Ok")
		i++
	}
	sendPlayerName(socket)
	messageServer, err := common.Receive(socket)
	checkErrorServer(err)
	fmt.Println(messageServer)
	processMenuOptions(socket, messageServer)
}

func isPingMessage(message string) bool {
	return strings.Contains(message, PING_MESSAGE)
}

func startGame(socket net.Conn) {
	i := 0

	for i < 2 {
		messageServer, err := common.Receive(socket)
		checkErrorServer(err)
		if (!isPingMessage(messageServer)){
			fmt.Println(common.BWhite+messageServer+common.NONE)
			i++
		}
		common.Send(socket, "Ok")
	}
}
