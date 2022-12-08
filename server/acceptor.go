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
	matchManager := &MatchManager{matches: []*Match{}}
	finishChannelWaitingPlayers := make(chan bool)
	finishChannelStartMatches := make(chan bool)
	go matchManager.processWaitingPlayers(finishChannelWaitingPlayers)
	go matchManager.startMatches(finishChannelStartMatches)
	go matchManager.delete_finish_matches()

	for {
		// acept diferent connections
		peer, error := acceptor.Accept()
		if error != nil {
			fmt.Println("Error accepting: ", error.Error())
			finishChannelWaitingPlayers <- true
			finishChannelStartMatches <- true
			os.Exit(1)
		}
		newPlayer := Player{id: len(acceptor.players) + 1, socket: peer, lastMove: 0}
		acceptor.players = append(acceptor.players, newPlayer)

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
