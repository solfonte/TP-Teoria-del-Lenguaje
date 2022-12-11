package server

import (
	"truco/app/common"
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
	trucoState            int
	alreadyAceptedTruco   bool
	alreadyAceptedRetruco bool
	alreadySangTruco      bool
	envidoState           int
	hasSangFinishRound    bool
}

/*********************************** ENVIDO FUNCTIONS ***********************************************/

func (move *Move) canSingEnvido() bool {
	return move.typeMove == FIRST_MOVE && !move.alreadySangEnvido && !move.alreadyAceptedTruco && !move.alreadyAceptedRetruco
}

func (move *Move) handleEnvido(player *Player, option int, playerError *PlayerError) {
	if option == CANTAR_ENVIDO_ENVIDO {
		SendInfoPlayer(player, common.SingEnvidoEnvido, playerError)
		move.envidoState = CANTAR_ENVIDO_ENVIDO
	} else if option == NO_QUERER_ENVIDO_ENVIDO {
		SendInfoPlayer(player, common.RejectEnvidoEnvido, playerError)
		move.envidoState = NO_QUERER_ENVIDO_ENVIDO
	} else if option == NO_QUERER_ENVIDO {
		SendInfoPlayer(player, common.RejectEnvido, playerError)
		move.envidoState = NO_QUERER_ENVIDO
	} else if option == QUERER_ENVIDO {
		SendInfoPlayer(player, common.AcceptEnvido, playerError)
		move.envidoState = QUERER_ENVIDO
	} else if option == CANTAR_ENVIDO {
		SendInfoPlayer(player, common.SingEnvido, playerError)
		move.envidoState = CANTAR_ENVIDO
		move.alreadySangEnvido = true
	}

}

/*********************************** TRUCO FUNCTIONS ***********************************************/

func (move *Move) canSingRetruco(player *Player) bool {
	return !move.alreadyAceptedRetruco && move.alreadySangTruco && !player.hasSagnTruco
}

func (move *Move) setAlreadySangTruco(player1 *Player, player2 *Player) {
	move.alreadySangTruco = (player1.hasSagnTruco || player2.hasSagnTruco)
	move.alreadyAceptedRetruco = (player1.hasSangReTruco || player2.hasSangReTruco)
}

func (move *Move) handleTruco(player *Player, option int, playerError *PlayerError) {
	if option == CANTAR_RETRUCO {
		SendInfoPlayer(player, common.SingRetruco, playerError)
		move.trucoState = CANTAR_RETRUCO
		player.hasSangReTruco = true
	} else if option == RECHAZAR_TRUCO {
		move.trucoState = RECHAZAR_TRUCO
		SendInfoPlayer(player, common.RejectTruco, playerError)
	} else if option == RECHAZAR_RETRUCO {
		move.trucoState = RECHAZAR_RETRUCO
		SendInfoPlayer(player, common.RejectRetruco, playerError)
	} else if option == ACEPTAR_RETRUCO {
		move.trucoState = ACEPTAR_RETRUCO
		SendInfoPlayer(player, common.AcceptRetruco, playerError)
		move.alreadyAceptedRetruco = true
		player.setNotifyRetruco(true)
	} else if option == ACEPTAR_TRUCO {
		SendInfoPlayer(player, common.AcceptTruco, playerError)
		move.trucoState = ACEPTAR_TRUCO
		move.alreadyAceptedTruco = true
		player.setNotifyTruco(true)
	} else if option == CANTAR_TRUCO {
		SendInfoPlayer(player, common.SingTruco, playerError)
		move.trucoState = CANTO_TRUCO
		move.alreadySangTruco = true
		player.hasSagnTruco = true
	}
}

func (move *Move) handleTrucoResult(actual *Player, opponent *Player, finish *bool, playerError *PlayerError) bool {

	if actual.lastMove == RECHAZAR_TRUCO || actual.lastMove == RECHAZAR_RETRUCO {
		message := common.OpponetRejectTruco
		if actual.lastMove == RECHAZAR_RETRUCO {
			message = common.OpponetRejectRetruco
		}
		SendInfoPlayer(opponent, message, playerError)
		return move.finish_round(opponent, actual, finish, playerError)
	} else if opponent.lastMove == RECHAZAR_TRUCO || opponent.lastMove == RECHAZAR_RETRUCO {
		message := common.OpponetRejectTruco
		if opponent.lastMove == RECHAZAR_RETRUCO {
			message = common.OpponetRejectRetruco
		}
		SendInfoPlayer(actual, message, playerError)
		return move.finish_round(actual, opponent, finish, playerError)
	} else if actual.lastMove == CANTAR_RETRUCO && opponent.lastMove == CANTAR_TRUCO {
		actual.turn = false
	} else if opponent.lastMove == CANTAR_RETRUCO && actual.lastMove == CANTAR_TRUCO {
		opponent.turn = false
	} else {
		return false
	}
	return false
}

/*********************************** HANDLERS ***********************************************/

func (move *Move) handleResult(actual *Player, opponent *Player, finish *bool, playerError *PlayerError) bool {
	if actual.lastMove == TIRAR_CARTA && opponent.lastMove == CANTAR_RETRUCO {
		opponent.lastMove = 0
	} else if opponent.lastMove == TIRAR_CARTA && actual.lastMove == CANTAR_RETRUCO {
		actual.lastMove = 0
	}
	if actual.lastMove == IRSE_AL_MAZO {
		SendInfoPlayer(opponent, common.OpponetHasSangFinishRound, playerError)
		return move.finish_round(opponent, actual, finish, playerError)
	} else if opponent.lastMove == IRSE_AL_MAZO {
		SendInfoPlayer(actual, common.OpponetHasSangFinishRound, playerError)
		return move.finish_round(actual, opponent, finish, playerError)
	} else if len(move.cardsPlayed) == MAX_CARDS_FOR_MOVE {
		message := common.BBlue + "Tu oponente tiro una carta " + move.cardsPlayed[1].card.getFullName() + common.NONE + "\n"
		SendInfoPlayer(opponent, message, playerError)
		result := move.cardsPlayed[0].card.compareCards(move.cardsPlayed[1].card)
		return move.assingWinner(result, move.cardsPlayed[0].player, move.cardsPlayed[1].player, finish, playerError)
	} else if envidoRelatedOptions(actual.lastMove, opponent.lastMove) {
		handleEnvidoResult(move, actual, opponent, finish, playerError)
		return false
	} else if trucoRelatedOptions(actual.lastMove, opponent.lastMove) {
		return move.handleTrucoResult(actual, opponent, finish, playerError)
	}
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

func (move *Move) handlePlayerActivity(player *Player, playerMove *int, playerError *PlayerError) {
	status := WAIT
	var err error
	for status != RETURN_FROM_WAITING_OPTIONS && status != -1 && status != IRSE_AL_MAZO {
		status, err = receiveWaitingRequests(player)

		handleWaitingOptions(status, player, playerMove, playerError)
		if err != nil {
			playerError.player = player
			playerError.err = err
		}
	}
	return
}

func (move *Move) handleFinishRound(player *Player, playerError *PlayerError) {
	SendInfoPlayer(player, common.SingFinishRound, playerError)
	move.hasSangFinishRound = true
}

func (move *Move) handleThrowACard(player *Player, playerError *PlayerError) int {
	playerCards := player.getCards()
	message, options := GetCardsToThrow(playerCards)

	option, err := loopSendOptionsToPlayer(options, player, playerError, message)
	cardPlayer := CardPlayer{card: playerCards[option-1], player: player}
	move.cardsPlayed = append(move.cardsPlayed, cardPlayer)
	err = sendPlayerCardPlayed(player, playerCards[option-1], playerError)
	player.removeCardSelected(option - 1)
	return err
}

func (move *Move) assingWinner(result int, player1 *Player, player2 *Player, finish *bool, playerError *PlayerError) bool {
	if result == 1 || result == 0 {
		return move.process_winner(player1, player2, finish, playerError)
	} else {
		return move.process_winner(player2, player1, finish, playerError)
	}
}

/*********************************** MOVE FUCTIONS ***********************************************/

func (move *Move) finish_round(winner *Player, loser *Player, finish *bool, playerError *PlayerError) bool {
	move.winner.id = winner.id
	move.winner.points = 0
	move.loser.id = loser.id
	move.loser.points = 0

	if move.hasSangFinishRound && move.trucoState != ACEPTAR_TRUCO && move.trucoState != ACEPTAR_RETRUCO && move.envidoState != QUERER_ENVIDO && move.envidoState != QUERER_ENVIDO_ENVIDO && move.trucoState != RECHAZAR_RETRUCO && move.trucoState != CANTAR_RETRUCO {
		move.winner.points = 1
		winner.points += 1
	} else {
		if move.trucoState == ACEPTAR_TRUCO || move.envidoState == QUERER_ENVIDO {
			move.winner.points = 2
			winner.points += 2
		} else if move.envidoState == QUERER_ENVIDO_ENVIDO {
			move.winner.points = 4
			winner.points += 4
		} else if move.trucoState == RECHAZAR_RETRUCO || (move.hasSangFinishRound && move.trucoState == CANTAR_RETRUCO) {
			move.winner.points = 2
			winner.points += 2
		} else if move.trucoState == ACEPTAR_RETRUCO {
			move.winner.points = 3
			winner.points += 3
		} else {
			move.winner.points = 1
			winner.points += 1
		}
	}

	*finish = true

	sendInfoPlayers(winner, loser, common.GetWinningMoveMessage(move.typeMove), common.GetLossingMoveMessage(move.typeMove), playerError)

	return true
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

		move.setAlreadySangTruco(player1, player2)
		if isTurnOfPlayer(player1) && !moveFinished && playerError.err == nil {
			orderChannel1 <- PLAY
			orderChannel2 <- WAIT
			option1 = <-movesChannel1
			option_wait = <-mazoChannell2
			if option_wait == IRSE_AL_MAZO {
				option2 = option_wait
				player2.lastMove = option_wait
			}
			moveFinished = move.handleResult(player1, player2, finish, playerError)
			movesChannel2 <- option1 //al jugador 2 le mando la jugada del jugador 1
		} else {
			option1 = 0
			player1.lastMove = 0
		}
		if isTurnOfPlayer(player2) && !moveFinished && playerError.err == nil {
			orderChannel1 <- WAIT
			orderChannel2 <- PLAY

			option2 = <-movesChannel2
			option_wait = <-mazoChannell1
			if option_wait == IRSE_AL_MAZO {
				option1 = option_wait
				player1.lastMove = option_wait
			}

			moveFinished = move.handleResult(player2, player1, finish, playerError)
			movesChannel1 <- option2
		} else {
			option2 = 0
			player2.lastMove = 0
		}
	}
	orderChannel1 <- STOP
	orderChannel2 <- STOP
	player1.lastMove = 0
	player2.lastMove = 0
	return err

}

func (move *Move) process_winner(winner *Player, loser *Player, finish *bool, playerError *PlayerError) bool {
	move.winner.id = winner.id
	move.winner.points = 0
	move.loser.id = loser.id
	move.loser.points = 0
	if !move.hasSangFinishRound {
		winner.winsPerPlay += 1
		if move.typeMove == LAST_MOVE || winner.winsPerPlay >= 2 {
			if move.alreadyAceptedRetruco {
				move.winner.points = 3
				winner.points += 3
			} else if winner.hasSagnTruco || loser.hasSagnTruco {
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
	sendInfoPlayers(winner, loser, common.GetWinningMoveMessage(move.typeMove), common.GetLossingMoveMessage(move.typeMove), playerError)
	if winner.winsPerPlay == 2 || move.hasSangFinishRound {
		*finish = true
	} else {
		*finish = false
	}
	return true
}

func (move *Move) askPlayerToWait(player *Player, playerOption *int, playerError *PlayerError) int {
	common.Send(player.socket, common.WaitPlayerToPlayMessage)
	_, err := common.Receive(player.socket)
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

func (move *Move) sendInfoMove(player *Player, options []int, playerError *PlayerError) (int, int) {
	message := getMessageInfoMoveToSend(move, options)
	option, err := loopSendOptionsToPlayer(options, player, playerError, message)
	return option, err
}

func (move *Move) askPlayerToMove(player *Player, options []int, playerError *PlayerError) (int, int) {
	option := 0
	var err int

	sendInfoOpponent(move, player, playerError)
	if playerError.err != nil {
		return -1, -1
	}
	option, err = move.sendInfoMove(player, options, playerError)

	if playerError.err != nil {
		return -1, -1
	}

	if option == IRSE_AL_MAZO {
		move.handleFinishRound(player, playerError)
		return IRSE_AL_MAZO, err
	} else if option == TIRAR_CARTA {
		err = move.handleThrowACard(player, playerError)
	} else {
		move.handleEnvido(player, option, playerError)
		move.handleTruco(player, option, playerError)
	}

	return option, err
}
