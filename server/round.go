package server

import (
	"fmt"
	"math"
	"truco/app/common"
)

type Round struct {
	players       map[int]*Player
	moves         int
	waitingPlayer *Player
	currentPlayer *Player
	envido        bool
	cardsPlayed   []Card
	points        int
	number        int
}

func (Round *Round) initialize(players map[int](*Player)) {
	Round.players = players
	Round.moves = 0
	Round.envido = false
	Round.points = 0
	Round.number = 0
}

func (round *Round) startRound(initialCurrentId int, initialWaitingId int, playerError *PlayerError, roundNumber int) int {
	completeRound := 1
	finish := false
	round.moves = 0
	round.number = roundNumber
	round.decide_hand_players(initialCurrentId, initialWaitingId)
	round.currentPlayer.winsPerPlay = 0
	round.waitingPlayer.winsPerPlay = 0
	var err int

	for completeRound <= 3 && !finish {
		fmt.Println("estoy en ronda ", completeRound)
		var move = Move{typeMove: completeRound, alreadySangEnvido: false} //TODO: no se si hace falta inicializarlo asi pero por ahora para probar
		err = move.start_move(round.currentPlayer, round.waitingPlayer, playerError, &finish)
		if err == -1 {
			return -1
		}
		//TODO: CAMBIAR EL RETURN PARA NO CORTAR EL FOR
		completeRound += 1

		round.decide_hand_players(move.winner.id, move.loser.id)

		fmt.Println("Puntos ronda", round.points)
	}
	fmt.Println("Gano ronda ", round.currentPlayer)
	fmt.Println("Puntos ronda", round.points)

	round.points = round.getMatchPointsPlayers(initialCurrentId, initialWaitingId)
	fmt.Println("Puntos ronda: ", round.points)
	sendInfoPlayers(round.currentPlayer, round.waitingPlayer, common.GetWinningRoundMessage(round.number), common.GetLossingRoundMessage(round.number))
	sendInfoPointsPlayers(round.currentPlayer, round.waitingPlayer)
	return round.points
}

func (round *Round) getMatchPointsPlayers(initialCurrentPlayerId int, initialWaitPlayerId int) int {
	return int(math.Max(float64(round.players[initialCurrentPlayerId].points), float64(round.players[initialWaitPlayerId].points)))
}

func (round *Round) decide_hand_players(initialCurrentPlayerId int, initialWaitPlayerId int) {
	round.waitingPlayer = round.players[initialWaitPlayerId]
	round.currentPlayer = round.players[initialCurrentPlayerId]
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
