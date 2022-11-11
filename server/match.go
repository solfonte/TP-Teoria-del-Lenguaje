package server

import (
	"fmt"
)

type Match struct {
	duration   int
	maxPlayers int
	players    map[int]*Player
	started    bool
	rounds     []Round
	points     int
}

func deal_cards(players map[int]*Player) {

	var cardDealer = CardDealer{}
	cardDealer.initialize()

	for _, p := range players {

		cardDealer.assignCards(p)
	}
}

func (match *Match) addPlayerToMatch(player *Player) {
	if match != nil {
		match.players[player.id] = player
		if len(match.players) == match.maxPlayers {
			match.started = true
			match.beginGame()
		}
	}
}

func (match *Match) beginGame() {
	deal_cards(match.players)
	fmt.Println("Entre a comenzo juego")

	var round = Round{}
	round.initialize(match.players)
	for _, player := range match.players {
		fmt.Println("primer carat ", player.cards[0].suit)
		startGame(*player)
	}
	for match.points <= 15 {
		match.points += round.startRound()
	}
}
