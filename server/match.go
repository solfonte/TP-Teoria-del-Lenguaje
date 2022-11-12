package server

import (
	"fmt"
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
}

func (match *Match) clearCards(players map[int]*Player) {
	for _, p := range players {
		p.cards = []Card{}
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
		if len(match.players) == match.maxPlayers {
			fmt.Println("Arranco la partida")
			match.waiterPlayerId = player.id
			match.started = true
			match.beginGame()
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

func (match *Match) beginGame() {
	match.deal_cards(match.players)
	fmt.Println("Entre a comenzo juego")

	var round = Round{}
	round.initialize(match.players)
	for _, player := range match.players {
		fmt.Println("primer carat ", player.cards[0].suit)
		startGame(*player)
	}
	for match.points <= 15 {
		match.points += round.startRound(match.initialPlayerId, match.waiterPlayerId)
		match.changeInitialPlayerForRounds()
		match.deal_cards(match.players)
	}
}
