package server

import (
	"fmt"
	"math/rand"
	"strconv"
	"truco/app/common"
)

type Round struct {
	players       []*Player
	moves         int
	hand          int
	currentPlayer *Player
	championId    int
	envido        bool
	cardsPlayed   []Card
}

func (Round *Round) initialize(players []*Player) {
	Round.players = players
	Round.moves = 0
	Round.hand = rand.Int() % len(players)
	Round.currentPlayer = players[Round.hand]
	Round.championId = -1
	Round.envido = false

}

func (round *Round) startRound() int {
	round.askPlayerForMove()
	changeTurn(round)
	common.Send(round.currentPlayer.socket, "The other player send this card"+round.cardsPlayed[0].getFullName())
	message, _ := common.Receive(round.currentPlayer.socket)
	fmt.Println(message)
	round.askPlayerForMove()
	return 0
}

func (Round *Round) canSingEnvido() bool {
	return (Round.moves <= 1 || !Round.envido)
}

func (Round *Round) handleEnvido() {
	common.Send(Round.currentPlayer.socket, "cantaste ENVIDO")
	fmt.Println("cantaste ENVIDO")
}

func (round *Round) handleThrowACard() {
	card1 := "1) " + round.currentPlayer.cards[0].getFullName()
	card2 := " 2) " + round.currentPlayer.cards[1].getFullName()
	card3 := " 3) " + round.currentPlayer.cards[2].getFullName()
	message := "Que carta queres tirar? "
	common.Send(round.currentPlayer.socket, message+card1+card2+card3+". Seleccione un numero:")

	jugada, _ := common.Receive(round.currentPlayer.socket)
	option, _ := strconv.Atoi(jugada)
	fmt.Println("Carta seleccionada ", round.currentPlayer.cards[option-1].getFullName())
	round.cardsPlayed = append(round.cardsPlayed, round.currentPlayer.cards[option-1])
}

func changeTurn(round *Round) {
	fmt.Println("mano antes de cambio ", round.hand)
	if round.hand < len(round.players)-1 {
		fmt.Println("entre a  caso donde cambio a 1")
		round.hand = 1
	} else {
		fmt.Println("entre a  caso donde cambio a 0")
		round.hand = 0
	}
	fmt.Println("mano despues de cambio ", round.hand)
	round.currentPlayer = round.players[round.hand]
}

func (round *Round) askPlayerForMove() int {
	fmt.Println("EN LA ROUND")
	fmt.Println(round.currentPlayer)

	for i := 0; i < len(round.players); i++ {
		if round.players[i].id != round.currentPlayer.id {
			fmt.Println("entre para jugador ", round.players[i])
			common.Send(round.players[i].socket, "Espera a que juegue tu oponente...")
			message, _ := common.Receive(round.players[i].socket)
			fmt.Println(message)
		}
	}
	messageEnvido := ""
	if round.canSingEnvido() {
		messageEnvido = "2) cantar envido"
		fmt.Println("entre a handle envido") //Investigar como hacer para hacer multiples sends sin que se trabe
	}
	message := "Es tu turno, podes hacer las siguientes jugadas: "
	command := "1) tirar una carta, " + messageEnvido + "3) cantar truco. Elija un numero"
	common.Send(round.currentPlayer.socket, message+command)

	jugada, _ := common.Receive(round.currentPlayer.socket)
	option, _ := strconv.Atoi(jugada)
	fmt.Println("option elegida: ", option)
	switch option {
	case 1:
		fmt.Println("opcion tirar una carta")
		round.handleThrowACard()
	case 2:
		round.handleEnvido()
	case 3:
		common.Send(round.currentPlayer.socket, "cantaste TRUCO")
		fmt.Println("cantaste TRUCO")
	}

	return option
}

func (Round *Round) getOpponentMove(move int) {
	switch move {
	case 1:
		common.Send(Round.currentPlayer.socket, "Tu oponente tiro una carta")
		//ask for a move
	case 2:
		common.Send(Round.currentPlayer.socket, "Tu oponente canto envido")
		//abrir caso de quiero, real envido, falta envido y decir los puntos
	case 3:
		common.Send(Round.currentPlayer.socket, "Tu oponente canto truco")
		//abrir caso de quiero, re truco, vale 4
		// y tirar cartas
	}
}
