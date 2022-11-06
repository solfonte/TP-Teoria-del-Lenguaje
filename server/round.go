package server

import (
	"fmt"
	"math/rand"
	"strconv"
	"truco/app/common"
)

type Round struct {
	players       []Player
	moves         int
	hand          int
	currentPlayer int
	championId    int
	envido        bool
}

func (Round *Round) initialize(players []Player) {
	Round.players = players
	Round.moves = 0
	Round.hand = rand.Int() % len(players)
	Round.currentPlayer = Round.hand
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
	common.Send(Round.players[Round.currentPlayer].socket, "cantaste ENVIDO")
	fmt.Println("cantaste ENVIDO")
}

func (Round *Round) handleThrowACard() int {
	common.Send(Round.players[Round.currentPlayer].socket, "Que carta queres tirar?")
	fmt.Println("dasdas")
	common.Send(Round.players[Round.currentPlayer].socket, "1) "+Round.players[Round.currentPlayer].cards[0].getFullName())
	fmt.Println("dasdas")
	common.Send(Round.players[Round.currentPlayer].socket, "2) "+Round.players[Round.currentPlayer].cards[1].getFullName())
	fmt.Println("dasdas")
	common.Send(Round.players[Round.currentPlayer].socket, "3) "+Round.players[Round.currentPlayer].cards[2].getFullName())
	fmt.Println("dasdas")

	common.Send(Round.players[Round.currentPlayer].socket, "Seleccione: ")

	jugada, _ := common.Receive(Round.players[Round.currentPlayer].socket)
	option, _ := strconv.Atoi(jugada)
	return option
}

func (Round *Round) askPlayerForMove() {
	fmt.Println("EN LA ROUND")
	fmt.Println(Round.players[Round.currentPlayer])
	for i := 0; i <= len(Round.players) && i != Round.currentPlayer; i++ {
		common.Send(Round.players[i].socket, "Espera a que juegue tu oponente...")
	}
	common.Send(Round.players[Round.currentPlayer].socket, "Podes hacer las siguientes jugadas:")
	common.Send(Round.players[Round.currentPlayer].socket, "1) tirar una carta")
	fmt.Println("SADSADSA") //Investigar como hacer para hacer multiples sends sin que se trabe
	if Round.canSingEnvido() {
		common.Send(Round.players[Round.currentPlayer].socket, "2) cantar envido")
		fmt.Println("SADSADSA") //Investigar como hacer para hacer multiples sends sin que se trabe
	}

	common.Send(Round.players[Round.currentPlayer].socket, "3) cantar truco")
	fmt.Println("SADSADSA") //Investigar como hacer para hacer multiples sends sin que se trabe
	common.Send(Round.players[Round.currentPlayer].socket, "Seleccione: ")

	jugada, _ := common.Receive(Round.players[Round.currentPlayer].socket)
	option, _ := strconv.Atoi(jugada)

	switch option {
	case 1:
		Round.handleThrowACard()
	case 2:
		Round.handleEnvido()
	case 3:
		common.Send(Round.players[Round.currentPlayer].socket, "cantaste TRUCO")
		fmt.Println("cantaste TRUCO")
	}
}
