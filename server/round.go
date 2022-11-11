package server

import (
	"fmt"
)

type Round struct {
	players       map[int]*Player
	moves         int
	waitingPlayer *Player
	currentPlayer *Player
	championId    int
	envido        bool
	cardsPlayed   []Card
	points        int
}

func (Round *Round) initialize(players map[int](*Player)) {
	Round.players = players
	Round.moves = 0
	Round.championId = -1
	Round.envido = false
	Round.points = 0
}

func (round *Round) startRound() int {
	completeRound := 1
	finish := false
	round.moves = 0
	fmt.Println("Arranca ronda")
	round.decide_hand_players()
	for completeRound <= 3 || !finish {
		var move = Move{typeMove: completeRound}
		finish = move.start_move(round.currentPlayer, round.waitingPlayer)
		completeRound += 1
		round.currentPlayer = round.players[move.winner.id]
		round.waitingPlayer = round.players[move.loser.id]
		round.points += move.getMaxPoints()
		fmt.Println("Puntos ronda", round.points)
	}
	fmt.Println("Gano ronda ", round.currentPlayer)
	fmt.Println("Puntos ronda", round.points)
	msgWinner := "Ganaste la ronda"
	msgLoser := "Perdiste la ronda"
	sendInfoPlayers(round.currentPlayer, round.waitingPlayer, msgWinner, msgLoser)
	return round.points
}

func (round *Round) decide_hand_players() {
	//Mañana veo
	round.waitingPlayer = round.players[1]
	round.currentPlayer = round.players[2]
}

// func (round *Round) waitingPlayerId() int {
// 	fmt.Println("jugador actual id", round.currentPlayerId)
// 	if round.currentPlayerId == 0 {
// 		fmt.Println("entre a  caso donde cambio a 1")
// 		return 1
// 	} else {
// 		fmt.Println("entre a  caso donde cambio a 0")
// 		return 0
// 	}
// }
