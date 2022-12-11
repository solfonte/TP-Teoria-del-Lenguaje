package server

import (
	"fmt"
	"net"
)

type Acceptor struct {
	net.Listener
	players      []Player
	matchManager *MatchManager
}

func start_receiver(acceptor Acceptor) {

	finishChannelWaitingPlayers := make(chan bool)
	finishChannelStartMatches := make(chan bool)
	finishChannelFinishMatch := make(chan bool)
	go acceptor.matchManager.processWaitingPlayers(finishChannelWaitingPlayers)
	go acceptor.matchManager.startMatches(finishChannelStartMatches)
	go acceptor.matchManager.deleteFinishMatches(finishChannelFinishMatch)

	for {
		// acept diferent connections
		peer, err := acceptor.Accept()
		if err != nil {
			finishChannelWaitingPlayers <- true
			finishChannelStartMatches <- true
			finishChannelFinishMatch <- true
			return
		}
		newPlayer := Player{id: len(acceptor.players) + 1, socket: peer, lastMove: 0, connected: true}
		acceptor.players = append(acceptor.players, newPlayer)

		acceptor.matchManager.process_player(&newPlayer)
		fmt.Println("client connected")
	}
}

func stop_receiver(acceptor Acceptor) {
	acceptor.Close()
	for _, player := range acceptor.players {
		if !acceptor.matchManager.playerFinish(player) {
			player.stop()
		}

	}
}
