package server

import (
	"fmt"
	"math"
	"net"
	"strconv"
	"truco/app/common"
)

const (
	opponentMessageForEnvido    = "Tu oponente canto envido. Tus opciones son: (1) Quiero (2) Quiero envido envido (3) No quiero"
	RETURN_FROM_WAITING_OPTIONS = 0
	TIRAR_CARTA                 = 4
	CANTAR_ENVIDO               = 5
	CANTAR_TRUCO                = 6
	QUERER_ENVIDO               = 7
	CANTAR_ENVIDO_ENVIDO        = 8
	NO_QUERER_ENVIDO_ENVIDO     = 9
	NO_QUERER_ENVIDO            = 10
	QUERER_ENVIDO_ENVIDO        = 13
	IRSE_AL_MAZO                = 11
	VER_MIS_CARTAS              = 12
	WAIT                        = 80
	STOP                        = 81
	PLAY                        = 82
)

const (
	CANTO_TRUCO    = 20
	ACEPTAR_TRUCO  = 21
	RECHAZAR_TRUCO = 22
	CANTAR_RETRUCO = 23
	ACEPTAR_RETRUCO = 24
	RECHAZAR_RETRUCO = 25
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
	envidoState         int
	hasSangFinishRound  bool
}

func (move *Move) canSingEnvido() bool {
	return move.typeMove == 1 && !move.alreadySangEnvido && !move.alreadyAceptedTruco
}

// luego sumar aca mismo otros tipo re truco y eso
func (move *Move) setAlreadySangTruco(player1 *Player, player2 *Player) {
	if player1 == nil {
		fmt.Println("es nul el 1")
	}
	if player2 == nil {
		fmt.Println("es nul el 2")
	}
	move.alreadySangTruco = (player1.hasSagnTruco || player2.hasSagnTruco)
}

func (move *Move) definePlayerPossibleOptions(actualOption int, opponentOption int) []int {
	var options []int
	if actualOption == NO_QUERER_ENVIDO || actualOption == QUERER_ENVIDO_ENVIDO {
		options = append(options, TIRAR_CARTA)
		if move.canSingEnvido() {
			options = append(options, CANTAR_ENVIDO)
		}
		if !move.alreadySangTruco {
			options = append(options, CANTAR_TRUCO)
		}
		return options
	}

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
		options = append(options, CANTAR_ENVIDO_ENVIDO)
		options = append(options, NO_QUERER_ENVIDO)
	} else if opponentOption == QUERER_ENVIDO || opponentOption == NO_QUERER_ENVIDO || opponentOption == NO_QUERER_ENVIDO_ENVIDO {
		options = append(options, TIRAR_CARTA)
		if !move.alreadySangTruco {
			options = append(options, CANTAR_TRUCO)
		}
	} else if opponentOption == CANTAR_ENVIDO_ENVIDO {
		options = append(options, QUERER_ENVIDO_ENVIDO)
		options = append(options, NO_QUERER_ENVIDO)
	} else if opponentOption == CANTAR_RETRUCO && move.trucoState == ACEPTAR_RETRUCO {
		options = append(options, TIRAR_CARTA)
	} else if opponentOption == CANTAR_TRUCO {
		options = append(options, ACEPTAR_TRUCO)
		options = append(options, RECHAZAR_TRUCO)
		options = append(options, CANTAR_RETRUCO)
	} else if opponentOption == CANTAR_RETRUCO {
		options = append(options, ACEPTAR_RETRUCO)
		options = append(options, RECHAZAR_RETRUCO)
	} else if opponentOption == ACEPTAR_TRUCO || opponentOption == ACEPTAR_RETRUCO {
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

	options = append(options, IRSE_AL_MAZO)
	return options
}

func (move *Move) finish_round(winner *Player, loser *Player, finish *bool) bool {
	move.winner.id = winner.id
	move.winner.points = 0
	move.loser.id = loser.id
	move.loser.points = 0

	if move.hasSangFinishRound && move.trucoState != ACEPTAR_TRUCO && move.trucoState != ACEPTAR_RETRUCO && move.envidoState != QUERER_ENVIDO && move.envidoState != QUERER_ENVIDO_ENVIDO && move.trucoState != RECHAZAR_RETRUCO && move.trucoState != CANTAR_RETRUCO {
		move.winner.points = 1
		winner.points += 1
		fmt.Println("@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@   1")
	} else {
		if move.trucoState == ACEPTAR_TRUCO || move.envidoState == QUERER_ENVIDO {
			move.winner.points = 2
			winner.points += 2
			fmt.Println("@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@  2")

		} else if move.envidoState == QUERER_ENVIDO_ENVIDO {
			move.winner.points = 4
			winner.points += 4
			fmt.Println("@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@  3")

		} else if move.trucoState == RECHAZAR_RETRUCO || (move.hasSangFinishRound && move.trucoState == CANTAR_RETRUCO){
			move.winner.points = 2
			winner.points += 2
			fmt.Println("@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@  4")

		} else if move.trucoState == ACEPTAR_RETRUCO {
			move.winner.points = 3
			winner.points += 3
			fmt.Println("@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@  5")

		} else {
			fmt.Println("@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@  6")
			move.winner.points = 1
			winner.points += 1
		}
	}

	*finish = true
	fmt.Println("ganador primera jugada ", move.winner, "\n\nPUNTOS GANADOR: ", move.winner.points)

	sendInfoPlayers(winner, loser, common.GetWinningMoveMessage(move.typeMove), common.GetLossingMoveMessage(move.typeMove))
	fmt.Println("jugadas ganadas ", winner.winsPerPlay)

	fmt.Println("termino jugada")
	return true
}

func (move *Move) handleEnvidoResult(actualOption int, opponentOption int, actual *Player, opponent *Player, finish *bool) {
	fmt.Println(" me llegan las opciones " + strconv.Itoa(actualOption) + " y " + strconv.Itoa(opponentOption))
	if (actualOption == CANTAR_ENVIDO || opponentOption == CANTAR_ENVIDO) && (actualOption == QUERER_ENVIDO || opponentOption == QUERER_ENVIDO) {
		envidoWinner := actual.verifyEnvidoWinnerAgainst(opponent)
		envidoWinner.sumPoints(2)
		fmt.Println("sume puntos por envido a " + envidoWinner.name)
	} else if (actualOption == CANTAR_ENVIDO || opponentOption == CANTAR_ENVIDO) && (actualOption == NO_QUERER_ENVIDO || opponentOption == NO_QUERER_ENVIDO) {
		playerToSumPoints := actual
		if opponentOption == CANTAR_ENVIDO {
			playerToSumPoints = opponent
		}
		playerToSumPoints.sumPoints(1)
		fmt.Println("sume puntos por envido no querido a " + playerToSumPoints.name)
	} else if (actualOption == CANTAR_ENVIDO_ENVIDO || opponentOption == CANTAR_ENVIDO_ENVIDO) && (actualOption == QUERER_ENVIDO_ENVIDO || opponentOption == QUERER_ENVIDO_ENVIDO) {
		envidoWinner := actual.verifyEnvidoWinnerAgainst(opponent)
		envidoWinner.sumPoints(2)
		fmt.Println("sume puntos por envido envido a " + envidoWinner.name)
	} else if (actualOption == CANTAR_ENVIDO_ENVIDO || opponentOption == CANTAR_ENVIDO_ENVIDO) && (actualOption == NO_QUERER_ENVIDO_ENVIDO || opponentOption == NO_QUERER_ENVIDO_ENVIDO) {
		playerToSumPoints := actual
		if opponentOption == CANTAR_ENVIDO {
			playerToSumPoints = opponent
		}
		playerToSumPoints.sumPoints(1)
		fmt.Println("sume puntos por envido envido a " + playerToSumPoints.name)
	}
}

func envidoRelatedOptions(playerOption int, anotherPlayerOption int) bool {
	options := []int{CANTAR_ENVIDO, QUERER_ENVIDO, QUERER_ENVIDO_ENVIDO, NO_QUERER_ENVIDO_ENVIDO, NO_QUERER_ENVIDO, QUERER_ENVIDO_ENVIDO}

	for _, option := range options {
		if playerOption == option || anotherPlayerOption == option {
			fmt.Println("opcion de envido identificada")
			return true
		}
	}
	return false
}

func (move *Move) handleResult(actualoption int, opponentOption int, actual *Player, opponent *Player, finish *bool) bool {
	fmt.Println("actual: ", actual.name)
	fmt.Println("oponente: ", opponent.name)
	if actualoption == IRSE_AL_MAZO {
		common.Send(opponent.socket, common.OpponetHasSangFinishRound)
		common.Receive(opponent.socket)
		fmt.Println("ACTUAL se fue al MAZO")
		return move.finish_round(opponent, actual, finish)
	} else if opponentOption == IRSE_AL_MAZO {
		fmt.Println("OPONENT se fue al  MAZO")
		common.Send(actual.socket, common.OpponetHasSangFinishRound)
		common.Receive(actual.socket)
		return move.finish_round(actual, opponent, finish)
	} else if envidoRelatedOptions(actualoption, opponentOption) {
		fmt.Println("identifique envido")
		move.handleEnvidoResult(actualoption, opponentOption, actual, opponent, finish)
		return false
	} else if actualoption == CANTAR_TRUCO || opponentOption == CANTAR_TRUCO {
		fmt.Println("alguno canto truco")
		return false
	} else if actualoption == CANTAR_RETRUCO || opponentOption == CANTAR_RETRUCO {
		fmt.Println("alguno canto REtruco")
		return false
	} else if len(move.cardsPlayed) == 2 {
		result := move.cardsPlayed[0].compareCards(move.cardsPlayed[1])
		return move.assingWinner(result, actual, opponent, finish)
	} else if actualoption == ACEPTAR_TRUCO || opponentOption == ACEPTAR_TRUCO {
		fmt.Println("alguno quiere truco")
		return false
	} else if opponentOption == RECHAZAR_TRUCO || opponentOption == RECHAZAR_RETRUCO {
		fmt.Println("No quiere truco")
		return move.finish_round(actual, opponent, finish)
	} else if actualoption == RECHAZAR_TRUCO || actualoption == RECHAZAR_RETRUCO{
		fmt.Println("No quiere truco")
		return move.finish_round(opponent, actual, finish)
	}

	if actualoption == TIRAR_CARTA && move.trucoState == ACEPTAR_RETRUCO {
		opponent.lastMove = 0
	}
	if opponentOption == TIRAR_CARTA && move.trucoState == ACEPTAR_RETRUCO {
		actual.lastMove = 0
	}
	return false
}

func (move *Move) handlePlayersMoves(orderChannel chan int, movesChannel chan int, player *Player, playerError *PlayerError) {
	var moveOrder int = -1
	var opponentOption int = 0
	for moveOrder != STOP {
		moveOrder = <-orderChannel

		if moveOrder == WAIT {
			opponentOption = move.askPlayerToWait(player, playerError)
			fmt.Println("Opponet option ", opponentOption)
			movesChannel <- opponentOption
			if playerError.err != nil {
				fmt.Println("//////////////////////////salgo de handelear al jugador " + player.name + "//////////////////////////////")
				return
			} else {
				opponentOption = <-movesChannel
			}
		} else if moveOrder == PLAY {
			options := move.definePlayerPossibleOptions(player.lastMove, opponentOption)
			actualPlayerOption, _ := move.askPlayerToMove(player, options, playerError)
			if playerError.err != nil {
				fmt.Println("//////////////////////////salgo de handelear al jugador " + player.name + "//////////////////////////////")

				return
			} else {
				player.lastMove = actualPlayerOption
				movesChannel <- actualPlayerOption
			}
		}

	}
}

func isTurnOfPlayer(player *Player) bool {
	fmt.Println("99999999999999999999999999  PLAYER MOVE", player.lastMove)
	return !(player.lastMove == CANTAR_ENVIDO_ENVIDO) && !(player.lastMove == CANTAR_RETRUCO)
}

func (move *Move) start_move(player1 *Player, player2 *Player, playerError *PlayerError, finish *bool) int {
	fmt.Println("entra a start_move")
	move.envidoState = 0
	err := 0
	var moveFinished bool
	var option1 int = 0
	var option2 int = 0
	orderChannel1 := make(chan int)
	orderChannel2 := make(chan int)
	movesChannel1 := make(chan int)
	movesChannel2 := make(chan int)
	//TODO: el player error va con mutex
	go move.handlePlayersMoves(orderChannel1, movesChannel1, player1, playerError)
	go move.handlePlayersMoves(orderChannel2, movesChannel2, player2, playerError)
	fmt.Println("start_move lanza los hilos")
	fmt.Println("--------------------estoy entrando a la jugada " + strconv.Itoa(move.typeMove) + "-----------------------")
	for !moveFinished && playerError.err == nil {

		move.setAlreadySangTruco(player1, player2) //TODO:chequear si va aca
		if isTurnOfPlayer(player1) && !moveFinished && playerError.err == nil {
			orderChannel1 <- PLAY
			orderChannel2 <- WAIT
			option1 = <-movesChannel1
			option2 = <-movesChannel2
			fmt.Println("option jugador1: ", option1)
			fmt.Println("Option jugador2: ", option2)
			moveFinished = move.handleResult(option1, option2, player1, player2, finish)
			fmt.Println("finish move: ", moveFinished)
			movesChannel2 <- option1 //al jugador 2 le mando la jugada del jugador 1
		}
		fmt.Println("finish move: ", moveFinished)
		if isTurnOfPlayer(player2) && !moveFinished && playerError.err == nil {
			fmt.Println("No tengo que entrar si alguien canto irse al mazo")
			orderChannel1 <- WAIT
			orderChannel2 <- PLAY

			option2 = <-movesChannel2
			option1 = <-movesChannel1
			fmt.Println("option jugador1: ", option1)
			fmt.Println("Option jugador2: ", option2)

			//ver si hay que resettear alguna opcion
			moveFinished = move.handleResult(option2, option1, player2, player1, finish)
			movesChannel1 <- option2 //al jugador 1 le mando la jugada del jugador 2
			fmt.Println("finish move: ", moveFinished)
		} else {
			fmt.Println("esta bien que entre aca si alguien tiro irse al mazo")
			option1 = 0
		}
	}
	fmt.Println("----------------------salgo del for de start move---------------------------------")
	orderChannel1 <- STOP
	orderChannel2 <- STOP
	player1.lastMove = 0
	player2.lastMove = 0
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
	if !move.hasSangFinishRound {
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
	}

	fmt.Println("ganador primera jugada ", move.winner, "\n\nPUNTOS GANADOR: ", move.winner.points)

	sendInfoPlayers(winner, loser, common.GetWinningMoveMessage(move.typeMove), common.GetLossingMoveMessage(move.typeMove))
	fmt.Println("jugadas ganadas ", winner.winsPerPlay)
	if winner.winsPerPlay == 2 || move.hasSangFinishRound {
		fmt.Println("termino jugada")
		*finish = true
	} else {
		fmt.Println("No termino jugada")
		*finish = false
	}
	// termino jugada
	return true
}

func receiveWaitingRequests(socket net.Conn) (int, error) {
	common.Send(socket, common.WaitingOptionsPlayer)
	message, err := common.Receive(socket)
	if err != nil {
		return -1, err
	}
	fmt.Println("el hilo de receive waiting requests recibio " + message)
	fmt.Println(message)
	fmt.Println("pase waiting requests")
	option, _ := strconv.Atoi(message)
	return option, nil
}

func (move *Move) handleWaitingOptions(status int, player *Player, playerError *PlayerError) {
	fmt.Println("STATUS ES " + string(status))

	if status == VER_MIS_CARTAS {
		fmt.Println("STATUS ES VER MIS CARTAS")
		sendInfoCards(*player, playerError)
	}
	if status == IRSE_AL_MAZO {
		common.Send(player.socket, common.SingFinishRound)
		/* chequear error */ common.Receive(player.socket)
	}
	return
}

func (move *Move) handlePlayerActivity(player *Player, playerError *PlayerError) int {
	status := WAIT
	var err error
	for status != RETURN_FROM_WAITING_OPTIONS && status != -1 && status != IRSE_AL_MAZO {
		status, err = receiveWaitingRequests(player.socket)
		move.handleWaitingOptions(status, player, playerError)
		if err != nil {
			fmt.Println("detecte error del q espera")
			playerError.player = player
			playerError.err = err
		}
	}
	fmt.Println("Salgo del for del waiting options ")
	if status == IRSE_AL_MAZO {
		return IRSE_AL_MAZO
	}
	return 0
}

func (move *Move) askPlayerToWait(player *Player, playerError *PlayerError) int {
	common.Send(player.socket, common.WaitPlayerToPlayMessage)
	message, err := common.Receive(player.socket)
	fmt.Println("mESNAJE QUE ME LLEGA EN AK PLAYER TO WAIT: ", message)
	if err != nil {
		playerError.player = player
		playerError.err = err
		return -1
	}

	status := move.handlePlayerActivity(player, playerError)

	if err != nil {
		playerError.player = player
		playerError.err = err
		return -1
	}

	return status
}

func (move *Move) handleEnvido(player *Player) {
	if move.envidoState == CANTAR_ENVIDO {
		common.Send(player.socket, common.AcceptEnvido)
		common.Receive(player.socket)
		move.envidoState = QUERER_ENVIDO
	} else {
		common.Send(player.socket, common.SingEnvido)
		common.Receive(player.socket)
		move.envidoState = CANTAR_ENVIDO
		move.alreadySangEnvido = true
	}

}

func (move *Move) handleTruco(player *Player, option int) {
	if (option == CANTAR_RETRUCO) {
		common.Send(player.socket, common.SingRetruco)
		common.Receive(player.socket)
		move.trucoState = CANTAR_RETRUCO
	} else if option == RECHAZAR_TRUCO {
		move.trucoState = RECHAZAR_TRUCO
	} else if option == RECHAZAR_RETRUCO {
		move.trucoState = RECHAZAR_RETRUCO
	} else if option == ACEPTAR_RETRUCO {
		move.trucoState = ACEPTAR_RETRUCO
		common.Send(player.socket, common.AcceptRetruco)
		common.Receive(player.socket)
		move.alreadySangTruco = true
	}else if option == ACEPTAR_TRUCO {
		common.Send(player.socket, common.AcceptTruco)
		common.Receive(player.socket)
		move.trucoState = ACEPTAR_TRUCO
	} else {
		common.Send(player.socket, common.SingTruco)
		common.Receive(player.socket)
		move.trucoState = CANTO_TRUCO
		move.alreadySangTruco = true
		player.hasSagnTruco = true
	}
}

func (move *Move) handleFinishRound(player *Player) {
	common.Send(player.socket, common.SingFinishRound)
	common.Receive(player.socket)
	move.hasSangFinishRound = true
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

func (move *Move) handleThrowACard(player *Player, playerError *PlayerError) int {
	playerCards := player.getCards()
	message, options := GetCardsToThrow(playerCards)

	option, err := loopSendOptionsToPlayer(options, player, playerError, message)
	move.cardsPlayed = append(move.cardsPlayed, playerCards[option-1])
	sendPlayerCardPlayed(player, playerCards[option-1])
	player.removeCardSelected(option - 1)
	return err
}

func loopSendOptionsToPlayer(options []int, player *Player, playerError *PlayerError, message string) (int, int) {
	option := 0
	msgError := ""
	for !containsOption(option, options) && playerError.err == nil {
		fmt.Println("loop options player: entre a mandarle la info a los jugadores")
		common.Send(player.socket, msgError+message)
		fmt.Println("mande info a cliente")
		jugada, err := common.Receive(player.socket)
		fmt.Println(err)
		if err != nil {
			fmt.Println("ERROR EN EL LOOP")
			playerError.player = player
			playerError.err = err
			return -1, -1
		}
		msgError = "Error: no elegiste una opcion valida. "
		option, _ = strconv.Atoi(jugada)
		fmt.Println("El jugador "+player.name+" mando la opcion: ", option)
	}
	return option, 0
}

func (move *Move) sendInfoMove(player *Player, options []int, playerError *PlayerError) (int, int) {
	fmt.Println("mando info")
	message := "Es tu turno, podes hacer las siguientes jugadas: " + "\n"
	for _, possibleOption := range options {
		if possibleOption == TIRAR_CARTA {
			message += common.BOLD + "(" + strconv.Itoa(TIRAR_CARTA) + ") " + common.NONE + common.CYAN + "Tirar" + common.NONE + " una carta" + "\n"
		} else if possibleOption == CANTAR_ENVIDO {
			message += common.BOLD + "(" + strconv.Itoa(CANTAR_ENVIDO) + ") " + common.NONE + "Cantar" + common.BYellow + " envido" + common.NONE + "\n"
		} else if possibleOption == CANTAR_TRUCO && move.trucoState < ACEPTAR_TRUCO {
			message += common.BOLD + "(" + strconv.Itoa(CANTAR_TRUCO) + ") " + common.NONE + "Cantar" + common.BRed + " truco " + common.NONE + "\n"
		} else if possibleOption == QUERER_ENVIDO {
			message += common.BOLD + "(" + strconv.Itoa(QUERER_ENVIDO) + ")" + common.NONE + common.GREEN + " Quiero envido " + common.NONE + "\n"
		} else if possibleOption == CANTAR_ENVIDO_ENVIDO {
			message += common.BOLD + "(" + strconv.Itoa(CANTAR_ENVIDO_ENVIDO) + ")" + common.NONE + common.GREEN + " Cantar envido envido" + common.NONE + "\n"
		} else if possibleOption == NO_QUERER_ENVIDO {
			message += common.BOLD + "(" + strconv.Itoa(NO_QUERER_ENVIDO) + ")" + common.NONE + common.RED + " No quiero envido " + common.NONE + "\n"
		} else if possibleOption == ACEPTAR_TRUCO {
			message += common.BOLD + "(" + strconv.Itoa(ACEPTAR_TRUCO) + ")" + common.NONE + common.GREEN + " Quiero truco " + common.NONE + "\n"
		} else if possibleOption == RECHAZAR_TRUCO {
			message += common.BOLD + "(" + strconv.Itoa(RECHAZAR_TRUCO) + ")" + common.NONE + common.RED + " Rechazar truco " + common.NONE + "\n"
		} else if possibleOption == QUERER_ENVIDO_ENVIDO {
			message += common.BOLD + "(" + strconv.Itoa(QUERER_ENVIDO_ENVIDO) + ")" + common.NONE + common.GREEN + " Quiero envido envido " + common.NONE + "\n"
		} else if possibleOption == IRSE_AL_MAZO {
			message += common.BOLD + "(" + strconv.Itoa(IRSE_AL_MAZO) + ")" + common.NONE + common.BWhite + " Irse al mazo " + common.NONE + "\n"
		} else if possibleOption == CANTAR_RETRUCO {
			message += common.BOLD + "(" + strconv.Itoa(CANTAR_RETRUCO) + ")" + common.NONE + common.GREEN + " Cantar retruco" + common.NONE + "\n"
		}else if possibleOption == ACEPTAR_RETRUCO {
			message += common.BOLD + "(" + strconv.Itoa(ACEPTAR_RETRUCO) + ")" + common.NONE + common.GREEN + " Quiero REtruco " + common.NONE + "\n"
		} else if possibleOption == RECHAZAR_RETRUCO {
			message += common.BOLD + "(" + strconv.Itoa(RECHAZAR_RETRUCO) + ")" + common.NONE + common.RED + " Rechazar REtruco " + common.NONE + "\n"
		} 
	}

	fmt.Println("sendInfoMove: todavia no entre al for")
	option, err := loopSendOptionsToPlayer(options, player, playerError, message)

	return option, err
}

func (move *Move) askPlayerToMove(player *Player, options []int, playerError *PlayerError) (int, int) {
	option := 0
	var err int
	fmt.Println("estado del envido: ", move.envidoState)

	if move.trucoState == CANTO_TRUCO {
		common.Send(player.socket, common.OpponentSingTruco)
		/*chequear errores*/ common.Receive(player.socket)
	} else if move.envidoState == CANTAR_ENVIDO {

		common.Send(player.socket, common.OpponetSingEnvido)
		/*chequear errores*/ common.Receive(player.socket)
	} else if move.envidoState == QUERER_ENVIDO {
		common.Send(player.socket, common.OpponetAcceptEnvido)
		common.Receive(player.socket)
		move.envidoState = 0
	} else if move.trucoState == ACEPTAR_TRUCO && !move.alreadyAceptedTruco {
		move.alreadyAceptedTruco = true

		common.Send(player.socket, common.OpponetAcceptTruco)
		/*chequear errores*/ common.Receive(player.socket)
	} else if move.trucoState == ACEPTAR_RETRUCO {

		common.Send(player.socket, common.OpponetAcceptRetruco)
		/*chequear errores*/ common.Receive(player.socket)
	} else if move.trucoState == RECHAZAR_TRUCO && !move.alreadyAceptedTruco {
		move.alreadyAceptedTruco = true

		common.Send(player.socket, common.OpponetRejectTruco)
		/*chequear errores*/ common.Receive(player.socket)
		return IRSE_AL_MAZO, err
	} else if move.trucoState == RECHAZAR_RETRUCO {
		common.Send(player.socket, common.OpponetRejectTruco)
		/*chequear errores*/ common.Receive(player.socket)
		return IRSE_AL_MAZO, err
	} else if len(move.cardsPlayed) > 0 {
		message := common.BBlue + "Tu oponente tiro una carta " + move.cardsPlayed[0].getFullName() + common.NONE + "\n"
		common.Send(player.socket, message)
		msg, _ := common.Receive(player.socket)
		fmt.Println("====respuesta de oponente tiro una carta", msg)
	}
	option, err = move.sendInfoMove(player, options, playerError)

	if playerError.err != nil {
		return -1, -1
	}

	switch option {
	case IRSE_AL_MAZO:
		move.handleFinishRound(player)
		return IRSE_AL_MAZO, err
	case TIRAR_CARTA:
		err = move.handleThrowACard(player, playerError)
	case CANTAR_ENVIDO:
		move.handleEnvido(player)
	case QUERER_ENVIDO:
		move.handleEnvido(player)
	case CANTAR_TRUCO:
		move.handleTruco(player, option)
	case CANTAR_RETRUCO:
		move.handleTruco(player, option)
	case ACEPTAR_TRUCO:
		move.handleTruco(player, option)
	case RECHAZAR_TRUCO:
		move.handleTruco(player, option)
	case RECHAZAR_RETRUCO:
		move.handleTruco(player, option)
	case ACEPTAR_RETRUCO:
		move.handleTruco(player, option)
	}
	if (option == ACEPTAR_TRUCO || option == ACEPTAR_RETRUCO) && len(move.cardsPlayed)%2 != 0 {
		return TIRAR_CARTA, err
	}
	return option, err
}
