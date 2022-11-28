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

	sendInfoPlayers(winner, loser, common.GetWinningMoveMessage(move.typeMove), common.GetLossingMoveMessage(move.typeMove))
	fmt.Println("jugadas ganadas ", winner.winsPerPlay)

	fmt.Println("termino jugada")
	*finish = true
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
	if envidoRelatedOptions(actualoption, opponentOption) {
		fmt.Println("identifique envido")
		move.handleEnvidoResult(actualoption, opponentOption, actual, opponent, finish)
		return false
	} else if actualoption == CANTAR_TRUCO || opponentOption == CANTAR_TRUCO {
		fmt.Println("alguno canto truco")
		return false
	} else if len(move.cardsPlayed) == 2 {
		result := move.cardsPlayed[0].compareCards(move.cardsPlayed[1])
		return move.assingWinner(result, actual, opponent, finish)
	} else if actualoption == ACEPTAR_TRUCO || opponentOption == ACEPTAR_TRUCO {
		fmt.Println("alguno quiere truco")
		return false
	} else if opponentOption == RECHAZAR_TRUCO || opponentOption == IRSE_AL_MAZO {
		fmt.Println("No quiere truco")
		return move.finish_round(actual, opponent, finish)
	} else if actualoption == RECHAZAR_TRUCO {
		fmt.Println("No quiere truco")
		return move.finish_round(opponent, actual, finish)
	}
	return false
}

func (move *Move) handlePlayersMoves(orderChannel chan int, movesChannel chan int, player *Player) {
	var moveOrder int = -1
	var playerError PlayerError
	var opponentOption int = 0
	for moveOrder != STOP {
		moveOrder = <-orderChannel

		if moveOrder == WAIT {
			move.askPlayerToWait(orderChannel, player, &playerError)
			opponentOption = <-movesChannel

		} else if moveOrder == PLAY {
			options := move.definePlayerPossibleOptions(player.lastMove, opponentOption)
			actualPlayerOption, _ := move.askPlayerToMove(player, options, &playerError)
			player.lastMove = actualPlayerOption
			movesChannel <- actualPlayerOption
		}

	}
	fmt.Println("//////////////////////////salgo de handelear al jugador " + player.name + "//////////////////////////////")
}

func isTurnOfPlayer(player *Player) bool {
	return !(player.lastMove == CANTAR_ENVIDO_ENVIDO)
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
	go move.handlePlayersMoves(orderChannel1, movesChannel1, player1)
	go move.handlePlayersMoves(orderChannel2, movesChannel2, player2)
	fmt.Println("start_move lanza los hilos")
	fmt.Println("--------------------estoy entrando a la jugada " + strconv.Itoa(move.typeMove) + "-----------------------")
	for !moveFinished && err != -1 {

		move.setAlreadySangTruco(player1, player2) //TODO:chequear si va aca
		if isTurnOfPlayer(player1) && !moveFinished {
			orderChannel1 <- PLAY
			orderChannel2 <- WAIT
			option1 = <-movesChannel1
			moveFinished = move.handleResult(option1, option2, player1, player2, finish)
			fmt.Println("finish move: ", moveFinished)
			movesChannel2 <- option1 //al jugador 2 le mando la jugada del jugador 1
		}

		if isTurnOfPlayer(player2) && !moveFinished {
			orderChannel1 <- WAIT
			orderChannel2 <- PLAY

			option2 = <-movesChannel2
			movesChannel1 <- option2 //al jugador 2 le mando la jugada del jugador 1
			//ver si hay que resettear alguna opcion
			moveFinished = move.handleResult(option2, option1, player1, player2, finish)
			fmt.Println("finish move: ", moveFinished)
		} else {
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

	sendInfoPlayers(winner, loser, common.GetWinningMoveMessage(move.typeMove), common.GetLossingMoveMessage(move.typeMove))
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

func receiveWaitingRequests(socket net.Conn) int {
	common.Send(socket, common.WaitingOptionsPlayer)
	message, err := common.Receive(socket)
	fmt.Println("el hilo de receive waiting requests recibio " + message)
	fmt.Println(message)
	if err != nil {
		return -1
	}
	fmt.Println("pase waiting requests")
	option, _ := strconv.Atoi(message)
	return option
}

func (move *Move) handleWaitingOptions(status int, player *Player) {
	if status == VER_MIS_CARTAS {
		//tenemos las card played pero faltaria decir quien tiro que
		sendInfoCards(*player)
	}
	return
}

func (move *Move) handlePlayerActivity(orderChannel chan int, player *Player) {
	status := WAIT
	for status != RETURN_FROM_WAITING_OPTIONS && len(orderChannel) == 0 && status != -1 {
		move.handleWaitingOptions(status, player)
		status = receiveWaitingRequests(player.socket)
	}
	if len(orderChannel) == 1 {
		fmt.Println("me llego una orden para el jugador " + player.name)
	}
}

func (move *Move) askPlayerToWait(orderChannel chan int, player *Player, playerError *PlayerError) int {
	common.Send(player.socket, common.WaitPlayerToPlayMessage)
	_, err := common.Receive(player.socket)
	if err != nil {
		playerError.player = player
		playerError.err = err
		return -1
	}

	move.handlePlayerActivity(orderChannel, player)

	if err != nil {
		playerError.player = player
		playerError.err = err
		return -1
	}

	return 0
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

	//TODO: agregar info de que aceptaron envido
}

func (move *Move) handleTruco(player *Player) {
	if move.trucoState == CANTO_TRUCO {
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

func (move *Move) handleFinishRound() {
	// si se canto truco/envido/envidoenvido, se cuenta el punto
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
	for !containsOption(option, options) {
		fmt.Println("loop options player: entre a mandarle la info a los jugadores")
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
	} else if move.trucoState == RECHAZAR_TRUCO && !move.alreadyAceptedTruco {
		move.alreadyAceptedTruco = true

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

	if err == -1 {
		return -1, -1
	}
	switch option {
	case IRSE_AL_MAZO:
		move.handleFinishRound()
		return IRSE_AL_MAZO, err
	case TIRAR_CARTA:
		err = move.handleThrowACard(player, playerError)
	case CANTAR_ENVIDO:
		move.handleEnvido(player)
	case QUERER_ENVIDO:
		move.handleEnvido(player)
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
