package server

import (
	"fmt"
	"truco/app/common"
)

type Match struct {
	duration        int
	maxPlayers      int
	players         map[int]*Player
	started         bool
	rounds          []Round
	points          int
	initialPlayerId int
	waiterPlayerId  int
	finish          bool
	readyToStart    bool
}

type PlayerError struct {
	player *Player
	err    error
}

func (match *Match) clearCards(players map[int]*Player) {
	for _, p := range players {
		p.clearCards()
	}
}

func (match *Match) deal_cards(players map[int]*Player) {
	match.clearCards(players)
	var cardDealer = CardDealer{}
	cardDealer.initialize()

	for _, p := range players {

		cardDealer.assignCards(p)
	}
}

func (match *Match) addPlayerToMatch(player *Player) {
	if match != nil {
		match.players[player.id] = player
		fmt.Println("agrego jugador :", player.name)
		if len(match.players) == match.maxPlayers {
			fmt.Println("Arranco la partida")
			match.waiterPlayerId = player.id
			match.readyToStart = true
			//match.beginGame()
		} else {
			// CREO ALGUIEN LA PARTIDA
			fmt.Println("alguiien creo la partida")
			match.initialPlayerId = player.id
		}
	}
}

func (match *Match) changeInitialPlayerForRounds() {
	newInitialPlayer := -1
	for key := range match.players {
		if key != match.initialPlayerId {
			newInitialPlayer = key
		}
	}
	match.waiterPlayerId = match.initialPlayerId
	match.initialPlayerId = newInitialPlayer
}

func (match *Match) handle_disconnection_player(playerError PlayerError) {
	fmt.Println("--------------------desconecto jugadores-------------------------")

	msg := "Tu oponente se desconecto"
	if playerError.player.id == match.initialPlayerId {
		sendOtherPlayDisconnection(*match.players[match.waiterPlayerId], msg)
	} else {
		sendOtherPlayDisconnection(*match.players[match.initialPlayerId], msg)
	}
	match.DisconnectMatch()
	match.finish = true
}

func (match *Match) DisconnectMatch() {
	fmt.Println("--------termiandoooooooooooooooo---")

	for _, player := range match.players {
		player.stop()
	}
}

func (match *Match) handleConnections(stop *bool, playerError *PlayerError) {
	for *stop == false {
		if playerError.err != nil {
			fmt.Println("desconecto")
			match.handle_disconnection_player(*playerError)
			return
		}
	}
}

func (match *Match) beginGame() {
	defer match.DisconnectMatch()
	match.deal_cards(match.players)
	fmt.Println("Entre a comenzo juego")

	var round = Round{}
	var playerError = PlayerError{err: nil, player: nil}
	var stop bool = false
	round.initialize(match.players)
	for _, player := range match.players {
		fmt.Println("comenzo juego")
		startGame(*player)
	}
	go match.handleConnections(&stop, &playerError)
	numberRound := 0
	for match.points < 6 {
		for _, player := range match.players {
			player.setHasSangTruco(false)
			player.setHasSangRetruco(false)
			player.setNotifyTruco(false)
			player.setNotifyRetruco(false)
		}
		sendInfoCards(*match.players[match.initialPlayerId], &playerError)
		sendInfoCards(*match.players[match.waiterPlayerId], &playerError)
		match.points = round.startRound(match.initialPlayerId, match.waiterPlayerId, &playerError, numberRound)

		fmt.Println("puntos que va el partido: ", match.points)
		numberRound += 1
		match.changeInitialPlayerForRounds()
		match.deal_cards(match.players)
	}
	match.process_winner_and_loser()
	match.FinishMatch()

}

func (match *Match) FinishMatch() {
	for _, player := range match.players {
		common.Send(player.socket, common.FinishGame)
		//common.Receive(player.socket)
	}
	match.finish = true
}

func (match Match) process_winner_and_loser() {
	if match.players[match.initialPlayerId].points >= match.players[match.waiterPlayerId].points {
		sendInfoPlayers(match.players[match.initialPlayerId],
			match.players[match.waiterPlayerId],
			common.WinMatchMessage,
			common.LoseMatchMessage)
	} else {
		sendInfoPlayers(match.players[match.waiterPlayerId],
			match.players[match.initialPlayerId],
			common.WinMatchMessage,
			common.LoseMatchMessage)
	}

}
