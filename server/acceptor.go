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
	matchManager := &MatchManager{matches: []Match{}}

	fmt.Printf(("estoy en receiver\n"))
	for {
		// acept diferent connections
		peer, error := acceptor.listenerSocket.Accept()
		if error != nil {
			//TODO: cuando se cierra el aceptador, no es un error pero sale por aca y figura como que el servidor "fallo"
			fmt.Println("Error accepting: ", error.Error())
			os.Exit(1)
		}
		newPlayer := Player{id: 1, socket: peer}
		acceptor.players = append(acceptor.players, newPlayer)
		go matchManager.process_player(&newPlayer)

		fmt.Println("client connected")

	}

}

func stop_receiver(acceptor Acceptor) {
	acceptor.listenerSocket.Close()
}
