package server

import (
	"fmt"
	"net"
	"os"
)

type Acceptor struct {
	net.Listener
	players []Player
}

func start_receiver(acceptor Acceptor) {
	matchManager := &MatchManager{matches: []Match{}}

	for {
		// acept diferent connections
		peer, error := acceptor.Accept()
		if error != nil {
			fmt.Println("Error accepting: ", error.Error())
			os.Exit(1)
		}
		newPlayer := Player{id: len(acceptor.players) + 1, socket: peer}
		acceptor.players = append(acceptor.players, newPlayer)

		matchManager.delete_finish_matches()
		matchManager.process_player(&newPlayer)

		fmt.Println("client connected")

	}

}

func stop_receiver(acceptor Acceptor) {
	acceptor.Close()
	for _, player := range acceptor.players {
		player.stop()
	}
}
