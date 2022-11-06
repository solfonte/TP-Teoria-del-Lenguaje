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
}

func (Round *Round) initialize(players []Player) {
	Round.players = players
	Round.moves = 0
	Round.hand = rand.Int() % len(players)
	Round.currentPlayer = Round.hand
	Round.championId = -1
}

func (Round *Round) startRound() {
	Round.askPlayerForMove()
	//recibir
	//leer y mandar la jugada que sea
}

func (Round *Round) askPlayerForMove() {
	fmt.Println("EN LA ROUND")
	fmt.Println(Round.players[Round.currentPlayer])
	for i := 0; i <= len(Round.players) && i != Round.currentPlayer; i++ {
		common.Send(Round.players[i].socket, "Espera a que juegue tu oponente")
	}
	for i := Round.currentPlayer + 1; i <= len(Round.players); i++ {
		common.Send(Round.players[i].socket, "Espera a que juegue tu oponente")
	}
	common.Send(Round.players[Round.currentPlayer].socket, "Podes hacer las siguientes jugadas:")
	common.Send(Round.players[Round.currentPlayer].socket, "1) tirar una carta")
	common.Send(Round.players[Round.currentPlayer].socket, "2) cantar envido")
	common.Send(Round.players[Round.currentPlayer].socket, "3) cantar truco")
	common.Send(Round.players[Round.currentPlayer].socket, "Seleccione: ")

	jugada, _ := common.Receive(Round.players[Round.currentPlayer].socket)
	option, _ := strconv.Atoi(jugada)

	switch option {
	case 1:
		common.Send(Round.players[Round.currentPlayer].socket, "tiraste una carta")
		fmt.Println("tiraste una carta")
		//meter aca lo de cual tirar
	case 2:
		common.Send(Round.players[Round.currentPlayer].socket, "cantaste ENVIDO")
		fmt.Println("cantaste ENVIDO")
	case 3:
		common.Send(Round.players[Round.currentPlayer].socket, "cantaste TRUCO")
		fmt.Println("cantaste TRUCO")
	}
}
