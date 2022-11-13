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
	finish          bool
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
			fmt.Println("tERMINO PARTIDA")
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
	for match.points <= 2 {
		sendInfoCards(*match.players[match.initialPlayerId])
		sendInfoCards(*match.players[match.waiterPlayerId])
		match.points += round.startRound(match.initialPlayerId, match.waiterPlayerId)
		match.changeInitialPlayerForRounds()
		match.deal_cards(match.players)
	}
	match.process_winner_and_loser()
	match.finish = true

}

func (match Match) process_winner_and_loser() {
	if match.players[match.initialPlayerId].points >= match.players[match.waiterPlayerId].points {
		sendInfoPlayers(match.players[match.initialPlayerId],
			match.players[match.waiterPlayerId],
			"Ganaste la partida :)",
			"Perdiste la partida :(")
	} else {
		sendInfoPlayers(match.players[match.waiterPlayerId],
			match.players[match.initialPlayerId],
			"Ganaste la partida :)",
			"Perdiste la partida :(")
	}
}
