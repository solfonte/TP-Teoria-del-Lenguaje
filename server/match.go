package server

import (
	"fmt"
)

type Match struct {
	duration   int
	maxPlayers int
	players    []Player
	started    bool
	rounds     []Round
}

func deal_cards(players []Player) {

	var cardDealer = CardDealer{}
	cardDealer.initialize()

	for _, p := range players {

		cardDealer.assignCards(&p)
	}
}

func (match *Match) addPlayerToMatch(player Player) {
	if match != nil {
		match.players = append(match.players, player)
		if len(match.players) == match.maxPlayers {
			match.started = true
			beginGame(match.players)
		}
	}
}

func beginGame(players []Player) {
	deal_cards(players)
	fmt.Println("Entre a comenzo juego")

	var round = Round{}
	round.initialize(players)
	for _, player := range players {
		startGame(player)
	}
	round.askPlayerForMove()
}
