package server

import (
	"fmt"
	"math"
	"strconv"
	"truco/app/common"
)

const (
	opponentMessageForEnvido = "Tu oponente canto envido. Tus opciones son: (1) Quiero (2) Quiero envido envido (3) No quiero"
	QUIERE_ENVIDO = 8
	QUIERE_ENVIDO_ENVIDO = 9
	NO_QUIERE_ENVIDO = 10
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
	alreadySangEnvido bool
}/*
func (move *Move) handleOpponentResponse(option int, actual *Player, opponent *Player){
	if option == QUIERE_ENVIDO{
		fmt.Println("entra a quiere envido")
		//TODO: el oponent es el q no es mano???? importante
		envidoWinner := actual.verifyEnvidoWinnerAgainst(opponent)
		envidoWinner.sumPoints(2)
	}else if option == NO_QUIERE_ENVIDO {
		actual.sumPoints(1)
	}

}*/
func (move *Move) handleResult(option1 int, option2 int, actual *Player, opponent *Player) bool {
	if option1 == QUIERE_ENVIDO || option2 == QUIERE_ENVIDO{
		fmt.Println("entra a quiere envido")
		//TODO: el oponent es el q no es mano???? importante
		envidoWinner := actual.verifyEnvidoWinnerAgainst(opponent)
		envidoWinner.sumPoints(2)
		return false
	}else if option1 == NO_QUIERE_ENVIDO {
		opponent.sumPoints(1)
		return false
	}else if option2 == NO_QUIERE_ENVIDO {
		actual.sumPoints(1)
		return false
	}
	return true
}


func (move *Move) start_move(player1 *Player, player2 *Player) bool {
	var i int = 0
	var moveFinished bool
	for !moveFinished {
		fmt.Println("cant de veces uque se ejecuta la misma ronda: " , i)
		move.askPlayerForWait(player2)
		option1 := move.askPlayerForMove(player1, false, 0)
		move.askPlayerForWait(player1)
		option2 := move.askPlayerForMove(player2, true, option1)
		moveFinished = move.handleResult(option1, option2, player1, player2)
		i += 1
	}
	if len(move.cardsPlayed) == 2 {
		fmt.Println("cards played dos/?")
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
	message, _ := common.Receive(player.socket) //TODO: porque esra aca este receive?
	fmt.Println(message)
}

func (move *Move) canSingEnvido() bool {
	return move.typeMove == 1
}

func (move *Move) handleEnvido(player *Player) {
	common.Send(player.socket, "cantaste ENVIDO")
}

func (move *Move) handleThrowACard(player *Player) {

	message := "Que carta queres tirar? "
	playerCards := player.getCards()
	for index, card := range playerCards {
		number := strconv.Itoa(index+1) + ") "
		message += number
		message += card.getFullName() + " "
	}
	common.Send(player.socket, message+". Seleccione un numero:")

	jugada, _ := common.Receive(player.socket)
	option, _ := strconv.Atoi(jugada)

	move.cardsPlayed = append(move.cardsPlayed, playerCards[option-1])
	player.removeCardSelected(option - 1)
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

func (move *Move) askPlayerForMove(player *Player, moveAsOpponent bool, action int) int {
	var option int;
	if moveAsOpponent{
		option = move.getOpponentMove(action, player)
	}else {
		option = move.sendInfoMove(player)
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
	}

	return option
}

func (move *Move) getOpponentMove(action int, player *Player) int {
	switch action {
	case 1:
		message := "Tu oponente tiro una carta " + move.cardsPlayed[0].getFullName()
		common.Send(player.socket, message)
		messageClient, _ := common.Receive(player.socket) //TODO: porque esta este receive
		fmt.Println(messageClient)
		move.askPlayerForMove(player, false, 0)
		return 1
	case 2:
		common.Send(player.socket, opponentMessageForEnvido)
		jugada, _ := common.Receive(player.socket)
		option, _ := strconv.Atoi(jugada)

		if option == 1 {
			fmt.Println("el oponente quiere envido")
			return QUIERE_ENVIDO
		}else if option == 2{
			fmt.Println("el oponente quiere envido envido")
			return QUIERE_ENVIDO_ENVIDO
		}else{
			fmt.Println("el oponente no quiere envido")
			return NO_QUIERE_ENVIDO
		}
	case 3:
		common.Send(player.socket, "Tu oponente canto truco")
		//abrir caso de quiero, re truco, vale 4
		// y tirar cartas
		return 3
	}
	return -1
}
