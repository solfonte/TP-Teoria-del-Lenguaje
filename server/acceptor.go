package server

import (
	"fmt"
	"net"
	"os"
)

type Acceptor struct {
	listenerSocket net.Listener
	players        []Player
}

func start_receiver(acceptor Acceptor) {
	matchManager := MatchManager{}
	fmt.Printf(("estoy en receiver"))
	for {
		// acept diferent connections
		peer, error := acceptor.listenerSocket.Accept()
		if error != nil {
			fmt.Println("Error accepting: ", error.Error())
			os.Exit(1)
		}
		newPlayer := Player{id: 1, name: "pepe", socket: peer}
		acceptor.players = append(acceptor.players, newPlayer)
		add_player(matchManager, newPlayer)

		fmt.Println("client connected")
		//go process_client(peer)
	}

}

// func process_client(connection net.Conn) {
// 	buffer := make([]byte, 1024)
// 	// leo lo que me llega
// 	dataLen, error := connection.Read(buffer)
// 	if error != nil {
// 		fmt.Println("Error reading:", error.Error())
// 		connection.Close()
// 	}
// 	fmt.Println("Received: ", string(buffer[:dataLen]))
// 	_, error = connection.Write([]byte("Llego este mensaje al server: " + string(buffer[:dataLen])))
// 	connection.Close()
// }

func stop_receiver(acceptor Acceptor) {
	acceptor.listenerSocket.Close()
}
