package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"truco/app/common"
)

func sendMatchParameters2(socket net.Conn) {
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
	common.Send(socket, "ok") //TODO: esto hay que sacarlo porque lo pusimos como patch para que no se bloquee
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

}
