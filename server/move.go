package server

import (
	"fmt"
	"math"
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
	CANTO_TRUCO      = 20
	ACEPTAR_TRUCO    = 21
	RECHAZAR_TRUCO   = 22
	CANTAR_RETRUCO   = 23
	ACEPTAR_RETRUCO  = 24
	RECHAZAR_RETRUCO = 25
)

type InfoPlayer struct {
	id     int
	points int
}

type CardPlayer struct {
	card   Card
	player *Player
}

type Move struct {
	winner                InfoPlayer
	loser                 InfoPlayer
	points                int
	typeMove              int
	cardsPlayed           []CardPlayer
	alreadySangEnvido     bool
	trucoState            int //20 canto truco, 21 se acepto truco, 22 se rechaza truco
	alreadyAceptedTruco   bool
	alreadyAceptedRetruco bool
	alreadySangTruco      bool
	envidoState           int
	hasSangFinishRound    bool
}

func (move *Move) canSingEnvido() bool {
	return move.typeMove == 1 && !move.alreadySangEnvido && !move.alreadyAceptedTruco && !move.alreadyAceptedRetruco
}

func (move *Move) canSingRetruco(player *Player) bool {
	fmt.Println("ya se acepto retruco: ", move.alreadyAceptedRetruco)
	return !move.alreadyAceptedRetruco && move.alreadySangTruco && !player.hasSagnTruco
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
	move.alreadyAceptedRetruco = (player1.hasSangReTruco || player2.hasSangReTruco)
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

		} else if move.trucoState == RECHAZAR_RETRUCO || (move.hasSangFinishRound && move.trucoState == CANTAR_RETRUCO) {
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

func (move *Move) handleResult(actual *Player, opponent *Player, finish *bool) bool {
	if actual.lastMove == TIRAR_CARTA && opponent.lastMove == CANTAR_RETRUCO {
		fmt.Println("entre actaul tiro una carta le seteo al otro last move en 0")
		opponent.lastMove = 0
	} else if opponent.lastMove == TIRAR_CARTA && actual.lastMove == CANTAR_RETRUCO {
		fmt.Println("entre oponente tiro una carta le seteo al otro last move en 0")
		actual.lastMove = 0
	}
	if actual.lastMove == IRSE_AL_MAZO {
		SendInfoPlayer(opponent, common.OpponetHasSangFinishRound)
		fmt.Println("ACTUAL se fue al MAZO")
		return move.finish_round(opponent, actual, finish)
	} else if opponent.lastMove == IRSE_AL_MAZO {
		fmt.Println("OPONENT se fue al  MAZO")
		SendInfoPlayer(actual, common.OpponetHasSangFinishRound)
		return move.finish_round(actual, opponent, finish)
	} else if envidoRelatedOptions(actual.lastMove, opponent.lastMove) {
		fmt.Println("identifique envido")
		handleEnvidoResult(actual.lastMove, opponent.lastMove, actual, opponent, finish)
		return false
	} else if actual.lastMove == CANTAR_TRUCO || opponent.lastMove == CANTAR_TRUCO {
		fmt.Println("alguno canto truco")
		return false
	} else if actual.lastMove == CANTAR_RETRUCO && opponent.lastMove != RECHAZAR_RETRUCO || opponent.lastMove == CANTAR_RETRUCO  && actual.lastMove != RECHAZAR_RETRUCO {
		fmt.Println("alguno canto REtruco")
		return false
	} else if len(move.cardsPlayed) == 2 {
		fmt.Println("hay dos cartas chequeo quien gana jugada")
		result := move.cardsPlayed[0].card.compareCards(move.cardsPlayed[1].card)
		return move.assingWinner(result, move.cardsPlayed[0].player, move.cardsPlayed[1].player, finish)
	} else if actual.lastMove == ACEPTAR_TRUCO || opponent.lastMove == ACEPTAR_TRUCO {
		fmt.Println("alguno quiere truco")
		return false
	} else if opponent.lastMove == RECHAZAR_TRUCO || opponent.lastMove == RECHAZAR_RETRUCO {
		fmt.Println("No quiere truco")
		return move.finish_round(actual, opponent, finish)
	} else if actual.lastMove == RECHAZAR_TRUCO || actual.lastMove == RECHAZAR_RETRUCO {
		fmt.Println("No quiere truco")
		return move.finish_round(opponent, actual, finish)
	}
	fmt.Println("no entro en ninguna de las opciones")

	return false
}

func (move *Move) handlePlayersMoves(orderChannel chan int, movesChannel chan int, mazoChannel chan int, player *Player, playerError *PlayerError) {
	var moveOrder int = -1
	var opponentOption int = 0
	for moveOrder != STOP {
		moveOrder = <-orderChannel

		if moveOrder == WAIT {
			move.askPlayerToWait(player, &opponentOption, playerError)
			mazoChannel <- opponentOption
			if playerError.err != nil {
				return
			} else {
				opponentOption = <-movesChannel
			}
		} else if moveOrder == PLAY {
			options := definePlayerPossibleOptions(move, player, opponentOption)
			actualPlayerOption, _ := move.askPlayerToMove(player, options, playerError)
			if playerError.err != nil {
				return
			} else {
				player.lastMove = actualPlayerOption
				movesChannel <- actualPlayerOption
			}
		}

	}
}

func (move *Move) start_move(player1 *Player, player2 *Player, playerError *PlayerError, finish *bool) int {
	move.envidoState = 0
	err := 0
	var moveFinished bool
	var option1 int = 0
	var option2 int = 0
	var option_wait int = 0
	orderChannel1 := make(chan int)
	orderChannel2 := make(chan int)
	movesChannel1 := make(chan int)
	movesChannel2 := make(chan int)
	mazoChannell1 := make(chan int)
	mazoChannell2 := make(chan int)
	//TODO: el player error va con mutex
	go move.handlePlayersMoves(orderChannel1, movesChannel1, mazoChannell1, player1, playerError)
	go move.handlePlayersMoves(orderChannel2, movesChannel2, mazoChannell2, player2, playerError)
	for !moveFinished && playerError.err == nil {

		move.setAlreadySangTruco(player1, player2) //TODO:chequear si va aca
		if isTurnOfPlayer(player1) && !moveFinished && playerError.err == nil {
			orderChannel1 <- PLAY
			orderChannel2 <- WAIT
			option1 = <-movesChannel1
			option_wait = <-mazoChannell2
			if option_wait == IRSE_AL_MAZO {
				option2 = option_wait
				player2.lastMove = option_wait
			}
			moveFinished = move.handleResult(player1, player2, finish)
			movesChannel2 <- option1 //al jugador 2 le mando la jugada del jugador 1
		}
		fmt.Println("finish move: ", moveFinished)
		if isTurnOfPlayer(player2) && !moveFinished && playerError.err == nil {
			orderChannel1 <- WAIT
			orderChannel2 <- PLAY

			option2 = <-movesChannel2
			option_wait = <-mazoChannell1
			if option_wait == IRSE_AL_MAZO {
				option1 = option_wait
				player1.lastMove = option_wait
			}

			moveFinished = move.handleResult(player2, player1, finish)
			movesChannel1 <- option2
		}/*creo qe este else no va*/
		 else {
			fmt.Println("esta bien que entre aca si alguien tiro irse al mazo")
			player1.lastMove = 0
		}
	}
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
			if move.alreadyAceptedRetruco {
				fmt.Println("PUNTORS RETRUCO")
				move.winner.points = 3
				winner.points += 3
			} else if winner.hasSagnTruco || loser.hasSagnTruco {
				fmt.Println("PUNTOS TRUCO")
				move.winner.points = 2
				winner.points += 2
			} else {
				fmt.Println("PUNTORS NORMALES")
				move.winner.points = 1
				winner.points += 1
			}
		} else {
			fmt.Println("NINGUN PUNTO")
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
	return true
}

func (move *Move) handlePlayerActivity(player *Player, playerMove *int, playerError *PlayerError) {
	status := WAIT
	var err error
	for status != RETURN_FROM_WAITING_OPTIONS && status != -1 && status != IRSE_AL_MAZO {
		status, err = receiveWaitingRequests(player)
		handleWaitingOptions(status, player, playerMove, playerError)
		if err != nil {
			fmt.Println("detecte error del q espera")
			playerError.player = player
			playerError.err = err
		}
	}
	fmt.Println("Salgo del for del waiting options ")
	return
}

func (move *Move) askPlayerToWait(player *Player, playerOption *int, playerError *PlayerError) int {
	common.Send(player.socket, common.WaitPlayerToPlayMessage)
	message, err := common.Receive(player.socket)
	fmt.Println("mESNAJE QUE ME LLEGA EN AK PLAYER TO WAIT: ", message)
	if err != nil {
		playerError.player = player
		playerError.err = err
		return -1
	}

	move.handlePlayerActivity(player, playerOption, playerError)

	if err != nil {
		playerError.player = player
		playerError.err = err
		return -1
	}
	return 0
}

func (move *Move) handleEnvido(player *Player) {
	if move.envidoState == CANTAR_ENVIDO {
		SendInfoPlayer(player, common.AcceptEnvido)
		move.envidoState = QUERER_ENVIDO
	} else {
		SendInfoPlayer(player, common.SingEnvido)
		move.envidoState = CANTAR_ENVIDO
		move.alreadySangEnvido = true
	}

}

func (move *Move) handleTruco(player *Player, option int) {
	if option == CANTAR_RETRUCO {
		SendInfoPlayer(player, common.SingRetruco)
		move.trucoState = CANTAR_RETRUCO
		player.hasSangReTruco = true
	} else if option == RECHAZAR_TRUCO {
		move.trucoState = RECHAZAR_TRUCO
		SendInfoPlayer(player, common.RejectTruco)
	} else if option == RECHAZAR_RETRUCO {
		move.trucoState = RECHAZAR_RETRUCO
		SendInfoPlayer(player, common.RejectRetruco)
	} else if option == ACEPTAR_RETRUCO {
		move.trucoState = ACEPTAR_RETRUCO
		SendInfoPlayer(player, common.AcceptRetruco)
		move.alreadyAceptedRetruco = true
	} else if option == ACEPTAR_TRUCO {
		SendInfoPlayer(player, common.AcceptTruco)
		move.trucoState = ACEPTAR_TRUCO
		move.alreadyAceptedTruco = true
	} else {
		SendInfoPlayer(player, common.SingTruco)
		move.trucoState = CANTO_TRUCO
		move.alreadySangTruco = true
		player.hasSagnTruco = true
	}
}

func (move *Move) handleFinishRound(player *Player) {
	SendInfoPlayer(player, common.SingFinishRound)
	move.hasSangFinishRound = true
}

func (move *Move) handleThrowACard(player *Player, playerError *PlayerError) int {
	playerCards := player.getCards()
	message, options := GetCardsToThrow(playerCards)

	option, err := loopSendOptionsToPlayer(options, player, playerError, message)
	cardPlayer := CardPlayer{card: playerCards[option-1], player: player}
	move.cardsPlayed = append(move.cardsPlayed, cardPlayer)
	sendPlayerCardPlayed(player, playerCards[option-1])
	player.removeCardSelected(option - 1)
	return err
}

func (move *Move) sendInfoMove(player *Player, options []int, playerError *PlayerError) (int, int) {
	fmt.Println("mando info")
	message := getMessageInfoMoveToSend(move, options)
	fmt.Println("sendInfoMove: todavia no entre al for")
	option, err := loopSendOptionsToPlayer(options, player, playerError, message)
	return option, err
}

func (move *Move) askPlayerToMove(player *Player, options []int, playerError *PlayerError) (int, int) {
	option := 0
	var err int
	fmt.Println("estado del envido: ", move.envidoState)

	if move.trucoState == CANTO_TRUCO {
		SendInfoPlayer(player, common.OpponentSingTruco)
	} else if move.trucoState == CANTAR_RETRUCO {
		SendInfoPlayer(player, common.OpponentSingRetruco)
	} else if move.envidoState == CANTAR_ENVIDO {
		SendInfoPlayer(player, common.OpponetSingEnvido)
	} else if move.envidoState == QUERER_ENVIDO {
		SendInfoPlayer(player, common.OpponetAcceptEnvido)
		move.envidoState = 0
	} else if move.envidoState == NO_QUERER_ENVIDO {
		SendInfoPlayer(player, common.OpponetRejectEnvido)
		move.envidoState = 0
	} else if move.trucoState == ACEPTAR_TRUCO && !move.alreadyAceptedTruco {
		move.alreadyAceptedTruco = true
		SendInfoPlayer(player, common.OpponetAcceptTruco)
	} else if move.trucoState == ACEPTAR_RETRUCO && !move.alreadyAceptedRetruco {
		SendInfoPlayer(player, common.OpponetAcceptRetruco)
	} else if move.trucoState == RECHAZAR_TRUCO && !move.alreadyAceptedTruco {
		move.alreadyAceptedTruco = true
		SendInfoPlayer(player, common.OpponetRejectTruco)
		return IRSE_AL_MAZO, err
	} else if move.trucoState == RECHAZAR_RETRUCO && !move.alreadyAceptedRetruco {
		SendInfoPlayer(player, common.OpponetRejectRetruco)
		return IRSE_AL_MAZO, err
	} else if len(move.cardsPlayed) > 0 {
		message := common.BBlue + "Tu oponente tiro una carta " + move.cardsPlayed[0].card.getFullName() + common.NONE + "\n"
		SendInfoPlayer(player, message)
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
	case NO_QUERER_ENVIDO:
		move.envidoState = NO_QUERER_ENVIDO
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
	return option, err
}
