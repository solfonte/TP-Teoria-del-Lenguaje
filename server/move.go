package server

import (
	"fmt"
	"math"
	"net"
	"strconv"
	"truco/app/common"
)

const (
	opponentMessageForEnvido = "Tu oponente canto envido. Tus opciones son: (1) Quiero (2) Quiero envido envido (3) No quiero"
	TIRAR_CARTA              = 4
	CANTAR_ENVIDO            = 5
	CANTAR_TRUCO             = 6
	QUERER_ENVIDO            = 7
	QUERER_ENVIDO_ENVIDO     = 8
	NO_QUERER_ENVIDO_ENVIDO  = 9
	NO_QUERER_ENVIDO         = 10
	IRSE_AL_MAZO             = 11
	VER_MIS_CARTAS           = 12
	WAIT                     = 80
	STOP                     = 81
)

const (
	CANTO_TRUCO      = 20
	ACEPTAR_TRUCO    = 21
	RECHAZAR_TRUCO   = 22
	TERMINAR_PARTIDA = 0
)

type InfoPlayer struct {
	id     int
	points int
}

type Move struct {
	winner              InfoPlayer
	loser               InfoPlayer
	points              int
	typeMove            int
	cardsPlayed         []Card
	alreadySangEnvido   bool
	trucoState          int //20 canto truco, 21 se acepto truco, 22 se rechaza truco
	alreadyAceptedTruco bool
	alreadySangTruco    bool
}

func (move *Move) canSingEnvido() bool {
	return move.typeMove == 1 && !move.alreadySangEnvido && !move.alreadyAceptedTruco
}

// luego sumar aca mismo otros tipo re truco y eso
func (move *Move) setAlreadySangTruco(player1 *Player, player2 *Player) {
	move.alreadySangTruco = (player1.hasSagnTruco || player2.hasSagnTruco)
}

func (move *Move) definePlayerPossibleOptions(opponentOption int) []int {
	var options []int
	if opponentOption == TIRAR_CARTA {
		options = append(options, TIRAR_CARTA)
		if move.canSingEnvido() {
			options = append(options, CANTAR_ENVIDO)
		}
		if !move.alreadySangTruco {
			options = append(options, CANTAR_TRUCO)
		}
	} else if opponentOption == CANTAR_ENVIDO {
		options = append(options, QUERER_ENVIDO)
		options = append(options, QUERER_ENVIDO_ENVIDO)
		options = append(options, NO_QUERER_ENVIDO)
	} else if opponentOption == QUERER_ENVIDO || opponentOption == NO_QUERER_ENVIDO || opponentOption == NO_QUERER_ENVIDO_ENVIDO {
		options = append(options, TIRAR_CARTA)
		options = append(options, CANTAR_TRUCO)
	} else if opponentOption == QUERER_ENVIDO_ENVIDO {
		options = append(options, TIRAR_CARTA)
		options = append(options, NO_QUERER_ENVIDO)
	} else if opponentOption == CANTAR_TRUCO {
		options = append(options, ACEPTAR_TRUCO)
		options = append(options, RECHAZAR_TRUCO)
	} else if opponentOption == ACEPTAR_TRUCO {
		options = append(options, TIRAR_CARTA)
	} else {
		options = append(options, TIRAR_CARTA)
		if move.canSingEnvido() {
			options = append(options, CANTAR_ENVIDO)
		}
		if !move.alreadySangTruco {
			options = append(options, CANTAR_TRUCO)
		}
	}
	return options
	//TODO: DESPUES SE AGREGA TRUCO
}

func (move *Move) finish_round(winner *Player, loser *Player, finish *bool) bool {
	move.winner.id = winner.id
	move.winner.points = 0
	move.loser.id = loser.id
	move.loser.points = 0
	// hay que settear a cero por cada ronda
	move.winner.points = 1
	winner.points += 1
	*finish = true
	fmt.Println("ganador primera jugada ", move.winner, "\n\nPUNTOS GANADOR: ", move.winner.points)
	msgwinner := "Ganaste la jugada " + strconv.Itoa(move.typeMove)
	msgLoser := "Perdiste la jugada" + strconv.Itoa(move.typeMove)
	sendInfoPlayers(winner, loser, msgwinner, msgLoser)
	fmt.Println("jugadas ganadas ", winner.winsPerPlay)

	fmt.Println("termino jugada")
	*finish = true
	return true
}

func (move *Move) handleResult(option1 int, option2 int, actual *Player, opponent *Player, finish *bool) bool {
	if option1 == QUERER_ENVIDO || option2 == QUERER_ENVIDO {
		//TODO: el oponent es el q no es mano???? importante
		envidoWinner := actual.verifyEnvidoWinnerAgainst(opponent)
		envidoWinner.sumPoints(2)
		return false
	} else if option1 == NO_QUERER_ENVIDO {
		opponent.sumPoints(1)
		return false
	} else if option2 == NO_QUERER_ENVIDO {
		actual.sumPoints(1)
		return false
	} else if option1 == CANTAR_TRUCO || option2 == CANTAR_TRUCO {
		fmt.Println("alguno canto truco")
		return false
	} else if option1 == TIRAR_CARTA && option2 == TIRAR_CARTA {
		result := move.cardsPlayed[0].compareCards(move.cardsPlayed[1])
		return move.assingWinner(result, actual, opponent, finish)
	} else if option1 == CANTAR_ENVIDO || option2 == CANTAR_ENVIDO {
		fmt.Print("alguno pidio envido")
		return false
	} else if option1 == ACEPTAR_TRUCO || option2 == ACEPTAR_TRUCO {
		fmt.Println("alguno quiere truco")
		return false
	} else if option2 == RECHAZAR_TRUCO {
		fmt.Println("No quiere truco")
		return move.finish_round(actual, opponent, finish)
	} else if option1 == RECHAZAR_TRUCO {
		fmt.Println("No quiere truco")
		return move.finish_round(opponent, actual, finish)
	}
	return true
}

func (move *Move) start_move(player1 *Player, player2 *Player, playerError *PlayerError, finish *bool) int {
	var moveFinished bool
	err := 0
	var option1 int = 0
	var option2 int = 0

	for !moveFinished && err != -1 {
		waitingChannelPlayer2 := make(chan int)
		msg := ""
		go move.askPlayerForWait(waitingChannelPlayer2, player2, playerError, "")
		move.setAlreadySangTruco(player1, player2)
		if err != -1 {
			moveFinished = move.handleResult(option1, option2, player1, player2, finish)
			options := move.definePlayerPossibleOptions(option2)
			option1, err = move.askPlayerForMove(player1, options, playerError, &msg)
		}
		waitingChannelPlayer2 <- STOP //TODO: no estamos captando los errores
		if err != -1 {
			moveFinished = move.handleResult(option1, option2, player1, player2, finish)
		}
		waitingChannelPlayer1 := make(chan int)

		if err != -1 {
			go move.askPlayerForWait(waitingChannelPlayer1, player1, playerError, "")
		}
		if err != -1 {
			options := move.definePlayerPossibleOptions(option1)
			option2, err = move.askPlayerForMove(player2, options, playerError, &msg)
		}
		waitingChannelPlayer1 <- STOP
		if err != -1 {
			moveFinished = move.handleResult(option1, option2, player1, player2, finish)
		}
		option1 = 0
	}
	//TODO:err podria ser bool
	return err
}

func (move *Move) assingWinner(result int, player1 *Player, player2 *Player, finish *bool) bool {
	if result == 1 || result == 0 {
		return move.process_winner(player1, player2, finish)
	} else {
		return move.process_winner(player2, player1, finish)
	}
}

func (move *Move) getMaxPoints() int {
	return int(math.Max(float64(move.loser.points), float64(move.winner.points)))
}

func (move *Move) process_winner(winner *Player, loser *Player, finish *bool) bool {
	move.winner.id = winner.id
	move.winner.points = 0
	move.loser.id = loser.id
	move.loser.points = 0
	// hay que settear a cero por cada ronda
	winner.winsPerPlay += 1
	if move.typeMove == 3 || winner.winsPerPlay >= 2 {
		if winner.hasSagnTruco || loser.hasSagnTruco {
			move.winner.points = 2
			winner.points += 2
		} else {
			move.winner.points = 1
			winner.points += 1
		}
	} else {
		move.winner.points = 0
	}

	fmt.Println("ganador primera jugada ", move.winner, "\n\nPUNTOS GANADOR: ", move.winner.points)
	msgwinner := "Ganaste la jugada " + strconv.Itoa(move.typeMove)
	msgLoser := "Perdiste la jugada" + strconv.Itoa(move.typeMove)
	sendInfoPlayers(winner, loser, msgwinner, msgLoser)
	fmt.Println("jugadas ganadas ", winner.winsPerPlay)
	if winner.winsPerPlay == 2 {
		fmt.Println("termino jugada")
		*finish = true
	} else {
		fmt.Println("No termino jugada")
		*finish = false
	}
	// termino jugada
	return true
}

func receiveWaitingRequests(waitingChannel chan<- int, socket net.Conn) {
	common.Send(socket, "Ingresas 11) Irse al mazo, 12) Consultar Cartas. Poner enter si no queres hacer nada")
	message, err := common.Receive(socket)
	option, _ := strconv.Atoi(message)
	if err != nil {
		waitingChannel <- -1
	}
	waitingChannel <- option
	return
}

func (move *Move) handleWaitingOptions(status int, player *Player) {
	if status == VER_MIS_CARTAS {
		//tenemos las card played pero faltaria decir quien tiro que
		message := "Estas son tus cartas actuales: "
		for _, card := range player.getCards() {
			message += card.getFullName() + " | "
		}
		common.Send(player.socket, message)
		common.Receive(player.socket) //receive de patch (ok)
	}
	return
}

func (move *Move) askPlayerForWait(waitingChannel chan int, player *Player, playerError *PlayerError, msg string) int {
	common.Send(player.socket, msg+"Espera a que juegue tu oponente...")
	_, err := common.Receive(player.socket)
	if err != nil {
		playerError.player = player
		playerError.err = err
		return -1
	}

	status := WAIT
	for status != STOP && status != -1 {
		move.handleWaitingOptions(status, player)
		go receiveWaitingRequests(waitingChannel, player.socket)
		status = <-waitingChannel
	}
	if err != nil {
		playerError.player = player
		playerError.err = err
		return -1
	}

	return 0
}

func (move *Move) handleEnvido(player *Player) {
	//common.Receive(player.socket)
	move.alreadySangEnvido = true
}

func (move *Move) handleTruco(player *Player) {
	if move.trucoState == CANTO_TRUCO {
		common.Send(player.socket, "aceptaste TRUCO")
		common.Receive(player.socket)
		move.trucoState = ACEPTAR_TRUCO
		//luego hacer lo de tirar carta
	} else {
		common.Send(player.socket, "cantaste TRUCO")
		common.Receive(player.socket)
		move.trucoState = CANTO_TRUCO
		move.alreadySangTruco = true
		player.hasSagnTruco = true
	}
}

func containsOption(option int, options []int) bool {
	var result bool = false
	for _, x := range options {
		if x == option {
			result = true
			break
		}
	}
	return result
}

func (move *Move) handleThrowACard(player *Player, playerError *PlayerError, msg *string) int {

	message := "Que carta queres tirar? "
	playerCards := player.getCards()

	var maxOptionsSelected []int
	for index, card := range playerCards {
		number := strconv.Itoa(index+1) + ") "
		message += number
		message += card.getFullName() + " "
		maxOptionsSelected = append(maxOptionsSelected, index+1)
	}
	option := 0
	msgError := ""
	for !containsOption(option, maxOptionsSelected) {
		common.Send(player.socket, msgError+message+". Seleccione un numero:")

		jugada, err := common.Receive(player.socket)
		if err != nil {
			fmt.Println("entre a error de receive")
			playerError.player = player
			playerError.err = err
			return -1
		}
		option, _ = strconv.Atoi(jugada)
		msgError = "Error: no elegiste una opcion valida. "
	}

	move.cardsPlayed = append(move.cardsPlayed, playerCards[option-1])
	*msg = "Tiraste la carta " + playerCards[option-1].getFullName() + "."
	player.removeCardSelected(option - 1)
	return 0
}

func (move *Move) sendInfoMove(player *Player, options []int, playerError *PlayerError) (int, int) {
	fmt.Println("mando info")
	message := "Es tu turno, podes hacer las siguientes jugadas: "
	for _, possibleOption := range options {
		if possibleOption == TIRAR_CARTA {
			message += "(" + strconv.Itoa(TIRAR_CARTA) + ") Tirar una carta "
		} else if possibleOption == CANTAR_ENVIDO {
			message += "(" + strconv.Itoa(CANTAR_ENVIDO) + ") Cantar envido "
		} else if possibleOption == CANTAR_TRUCO && move.trucoState < ACEPTAR_TRUCO {
			message += "(" + strconv.Itoa(CANTAR_TRUCO) + ") Cantar truco "
		} else if possibleOption == QUERER_ENVIDO {
			message += "(" + strconv.Itoa(QUERER_ENVIDO) + ") Quiero envido "
		} else if possibleOption == QUERER_ENVIDO_ENVIDO {
			message += "(" + strconv.Itoa(QUERER_ENVIDO_ENVIDO) + ") Quiero envido envido "
		} else if possibleOption == NO_QUERER_ENVIDO {
			message += "(" + strconv.Itoa(NO_QUERER_ENVIDO) + ") No quiero envido "
		} else if possibleOption == ACEPTAR_TRUCO {
			message += "(" + strconv.Itoa(ACEPTAR_TRUCO) + ") Quiero truco "
		} else if possibleOption == RECHAZAR_TRUCO {
			message += "(" + strconv.Itoa(RECHAZAR_TRUCO) + ") Rechazar truco "
		}
	}
	//esto va en otra funcion
	option := 0
	msgError := ""
	//TODO: tanto en sendInfo move como en handlethwro se hace el mismo loop, ver de meterlos
	// en una misma funcion.(recibea mensaje y cant de optiones y devuelvan la opcion elegida)
	fmt.Println("todavia no entre al for")

	for !containsOption(option, options) {
		common.Send(player.socket, msgError+message)

		jugada, err := common.Receive(player.socket)
		if err != nil {
			playerError.player = player
			playerError.err = err
			return -1, -1
		}
		msgError = "Error: no elegiste una opcion valida. "
		option, _ = strconv.Atoi(jugada)
		fmt.Println("El jugador "+player.name+" mando la opcion: ", option)
	}
	fmt.Println("sali")
	return option, 0
}

func (move *Move) askPlayerForMove(player *Player, options []int, playerError *PlayerError, msg *string) (int, int) {
	option := 0
	var err int
	if len(move.cardsPlayed) > 0 {
		message := "Tu oponente tiro una carta " + move.cardsPlayed[0].getFullName()
		common.Send(player.socket, message)
		/*chequear errores*/ common.Receive(player.socket)
	}
	if move.trucoState == CANTO_TRUCO {
		message := "Tu oponente canto TRUCO"
		common.Send(player.socket, message)
		/*chequear errores*/ common.Receive(player.socket)
	}
	if move.trucoState == ACEPTAR_TRUCO && !move.alreadyAceptedTruco {
		move.alreadyAceptedTruco = true
		message := "Tu oponente Acepto el TRUCO"
		common.Send(player.socket, message)
		/*chequear errores*/ common.Receive(player.socket)
	}
	if move.trucoState == RECHAZAR_TRUCO && !move.alreadyAceptedTruco {
		move.alreadyAceptedTruco = true
		message := "Tu oponente Rechazo el TRUCO"
		common.Send(player.socket, message)
		/*chequear errores*/ common.Receive(player.socket)
		return TERMINAR_PARTIDA, err
	}

	option, err = move.sendInfoMove(player, options, playerError)

	if err == -1 {
		return -1, -1
	}

	switch option {
	case TIRAR_CARTA:
		err = move.handleThrowACard(player, playerError, msg)
	case CANTAR_ENVIDO:
		move.handleEnvido(player)
		*msg = "Cantaste envido."
	case CANTAR_TRUCO:
		fmt.Println("canto truco")
		move.handleTruco(player)
	case ACEPTAR_TRUCO:
		move.handleTruco(player)
	case RECHAZAR_TRUCO:
		move.trucoState = RECHAZAR_TRUCO
	}
	if option == ACEPTAR_TRUCO && len(move.cardsPlayed)%2 != 0 {
		return TIRAR_CARTA, err
	}
	return option, err
}
