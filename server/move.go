package server

import (
	"fmt"
	"math"
	"strconv"
	"truco/app/common"
)

type InfoPlayer struct {
	id     int
	points int
}

type Move struct {
	winner      InfoPlayer
	loser       InfoPlayer
	points      int
	typeMove    int
	cardsPlayed []Card
}

func (move *Move) start_move(player1 *Player, player2 *Player) bool {
	move.askPlayerForWait(player2)
	option := move.askPlayerForMove(player1)
	move.askPlayerForWait(player1)
	move.getOpponentMove(option, player2)
	if len(move.cardsPlayed) == 2 {
		result := move.cardsPlayed[0].compareCards(move.cardsPlayed[1])
		return move.assingWinner(result, player1, player2)
	}
	return false
}

func (move *Move) assingWinner(result int, player1 *Player, player2 *Player) bool {
	if result == 1 || result == 0 {
		return move.process_winner(player1, player2)
	} else {
		return move.process_winner(player2, player1)
	}
}

func (move *Move) getMaxPoints() int {
	return int(math.Max(float64(move.loser.points), float64(move.winner.points)))
}

func (move *Move) process_winner(winner *Player, loser *Player) bool {
	move.winner.id = winner.id
	move.winner.points = 0
	move.loser.id = loser.id
	move.loser.points = 0
	// hay que settear a cero por cada ronda
	winner.winsPerPlay += 1
	if move.typeMove == 3 || winner.winsPerPlay >= 2 {
		fmt.Println("asdinos puntos partida a ganador")
		move.winner.points = 1
		winner.points += 1
		fmt.Println("Puntos jugador ganador", winner.points, winner.name)
	} else {
		move.winner.points = 0
	}

	fmt.Println("ganador primera jugada ", move.winner)
	msgwinner := "Ganaste la jugada " + strconv.Itoa(move.typeMove)
	msgLoser := "Perdiste la jugada" + strconv.Itoa(move.typeMove)
	sendInfoPlayers(winner, loser, msgwinner, msgLoser)
	fmt.Println("jugadas ganadas ", winner.winsPerPlay)
	if winner.winsPerPlay == 2 {
		fmt.Println("termino jugada")
		return true
	} else {
		fmt.Println("No termino jugada")
		return false
	}
}

func (move *Move) askPlayerForWait(player *Player) {
	common.Send(player.socket, "Espera a que juegue tu oponente...")
	message, _ := common.Receive(player.socket)
	fmt.Println(message)
}

func (move *Move) canSingEnvido() bool {
	return move.typeMove == 1
}

func (move *Move) handleEnvido(player *Player) {
	common.Send(player.socket, "cantaste ENVIDO")
	fmt.Println("cantaste ENVIDO")
}

func (move *Move) handleThrowACard(player *Player) {

	message := "Que carta queres tirar? "
	for index, card := range player.cards {
		number := strconv.Itoa(index+1) + ") "
		message += number
		message += card.getFullName()
	}
	common.Send(player.socket, message+". Seleccione un numero:")

	jugada, _ := common.Receive(player.socket)
	option, _ := strconv.Atoi(jugada)
	fmt.Println("Carta seleccionada ", player.cards[option-1].getFullName())
	move.cardsPlayed = append(move.cardsPlayed, player.cards[option-1])
	fmt.Println("cartas", player.cards)
	player.removeCardSelected(option - 1)
	fmt.Println("cartas", player.cards)
}

func (move *Move) sendInfoMove(player *Player) int {
	messageEnvido := ""
	if move.canSingEnvido() {
		messageEnvido = "2) cantar envido"
	}
	message := "Es tu turno, podes hacer las siguientes jugadas: "
	command := "1) tirar una carta, " + messageEnvido + "3) cantar truco. Elija un numero"
	common.Send(player.socket, message+command)

	jugada, _ := common.Receive(player.socket)
	option, _ := strconv.Atoi(jugada)
	return option
}

func (move *Move) askPlayerForMove(player *Player) int {
	fmt.Println("EN LA rpimer jufada")
	fmt.Println(player)
	option := move.sendInfoMove(player)
	fmt.Println("option elegida: ", option)
	switch option {
	case 1:
		fmt.Println("opcion tirar una carta")
		move.handleThrowACard(player)
	case 2:
		move.handleEnvido(player)
	case 3:
		common.Send(player.socket, "cantaste TRUCO")
		fmt.Println("cantaste TRUCO")
	}
	//chequear numeros

	return option
}

func (move *Move) getOpponentMove(action int, player *Player) {
	switch action {
	case 1:
		message := "Tu oponente tiro una carta " + move.cardsPlayed[0].getFullName()
		common.Send(player.socket, message)
		messageClient, _ := common.Receive(player.socket)
		fmt.Println(messageClient)
		move.askPlayerForMove(player)
	case 2:
		common.Send(player.socket, "Tu oponente canto envido")
		//abrir caso de quiero, real envido, falta envido y decir los puntos
	case 3:
		common.Send(player.socket, "Tu oponente canto truco")
		//abrir caso de quiero, re truco, vale 4
		// y tirar cartas
	}
}
