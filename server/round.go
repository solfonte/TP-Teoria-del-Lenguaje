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
}

func (Round *Round) initialize(players []*Player) {
	Round.players = players
	Round.moves = 0
	Round.hand = rand.Int() % len(players)
	Round.currentPlayer = players[Round.hand]
	Round.championId = -1
	Round.envido = false
}

func (Round *Round) startRound() {
	Round.askPlayerForMove()
	//recibir
	//leer y mandar la jugada que sea
}

func (Round *Round) canSingEnvido() bool {
	return (Round.moves <= 1 || !Round.envido)
}

func (Round *Round) handleEnvido() {
	common.Send(Round.currentPlayer.socket, "cantaste ENVIDO")
	fmt.Println("cantaste ENVIDO")
}

func (Round *Round) handleThrowACard() Card {
	card1 := "1) " + Round.currentPlayer.cards[0].getFullName()
	card2 := " 2) " + Round.currentPlayer.cards[1].getFullName()
	card3 := " 3) " + Round.currentPlayer.cards[2].getFullName()
	message := "Que carta queres tirar? "
	common.Send(Round.currentPlayer.socket, message+card1+card2+card3+". Seleccione un numero:")

	jugada, _ := common.Receive(Round.currentPlayer.socket)
	option, _ := strconv.Atoi(jugada)
	fmt.Println("Carta seleccionada ", Round.currentPlayer.cards[option-1].getFullName())
	return Round.currentPlayer.cards[option-1]
}

func (Round *Round) askPlayerForMove() {
	fmt.Println("EN LA ROUND")
	fmt.Println(Round.currentPlayer)

	for i := 0; i < len(Round.players); i++ {
		if Round.players[i].id != Round.currentPlayer.id {
			fmt.Println("entre")
			common.Send(Round.players[i].socket, "Espera a que juegue tu oponente...")
		}
	}
	messageEnvio := ""
	if Round.canSingEnvido() {
		messageEnvio = "2) cantar envido"
		fmt.Println("entre a handle envido") //Investigar como hacer para hacer multiples sends sin que se trabe
	}
	message := "Es tu turno, podes hacer las siguientes jugadas: "
	command := "1) tirar una carta, " + messageEnvio + "3) cantar truco. Elija un numero"
	common.Send(Round.currentPlayer.socket, message+command)

	jugada, _ := common.Receive(Round.currentPlayer.socket)
	option, _ := strconv.Atoi(jugada)
	fmt.Println("option elegida: ", option)
	switch option {
	case 1:
		fmt.Println("opcion tirar una carta")
		Round.handleThrowACard()
	case 2:
		Round.handleEnvido()
	case 3:
		common.Send(Round.currentPlayer.socket, "cantaste TRUCO")
		fmt.Println("cantaste TRUCO")
	}
	Round.changeTurn()
}
