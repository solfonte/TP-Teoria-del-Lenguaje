package server

import (
	"fmt"
	"strconv"
	"truco/app/common"
)

func definePlayerPossibleOptions(move *Move, player *Player, opponentOption int) []int {
	var options []int
	if (player.lastMove == NO_QUERER_ENVIDO || player.lastMove == NO_QUERER_ENVIDO_ENVIDO || player.lastMove == QUERER_ENVIDO_ENVIDO) && opponentOption != CANTAR_TRUCO {
		options = append(options, TIRAR_CARTA)
		if !move.alreadySangTruco {
			options = append(options, CANTAR_TRUCO)
		}
		options = append(options, IRSE_AL_MAZO)
		player.lastMove = 0
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
		if move.canSingRetruco(player) {
			options = append(options, CANTAR_RETRUCO)
		}
	} else if opponentOption == CANTAR_ENVIDO {
		options = append(options, QUERER_ENVIDO)
		options = append(options, CANTAR_ENVIDO_ENVIDO)
		options = append(options, NO_QUERER_ENVIDO)
	} else if opponentOption == QUERER_ENVIDO_ENVIDO || opponentOption == QUERER_ENVIDO || opponentOption == NO_QUERER_ENVIDO || opponentOption == NO_QUERER_ENVIDO_ENVIDO {
		options = append(options, TIRAR_CARTA)
		if !move.alreadySangTruco {
			options = append(options, CANTAR_TRUCO)
		}
	} else if opponentOption == CANTAR_ENVIDO_ENVIDO {
		options = append(options, QUERER_ENVIDO_ENVIDO)
		options = append(options, NO_QUERER_ENVIDO_ENVIDO)
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
		if move.canSingRetruco(player) {
			options = append(options, CANTAR_RETRUCO)
		}
	}

	options = append(options, IRSE_AL_MAZO)
	return options
}

func isTurnOfPlayer(player *Player) bool {
	return !(player.lastMove == CANTAR_ENVIDO_ENVIDO) && !(player.lastMove == CANTAR_RETRUCO && !player.turn)
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

func loopSendOptionsToPlayer(options []int, player *Player, playerError *PlayerError, message string) (int, int) {
	option := 0
	msgError := ""
	for !containsOption(option, options) && playerError.err == nil {
		common.Send(player.socket, msgError+message)
		jugada, err := common.Receive(player.socket)
		if err != nil {
			fmt.Println("ERROR EN EL LOOP")
			playerError.player = player
			playerError.err = err
			return -1, -1
		}
		msgError = "Error: no elegiste una opcion valida. "
		option, _ = strconv.Atoi(jugada)
	}
	return option, 0
}

func getMessageInfoMoveToSend(move *Move, options []int) string {
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
		} else if possibleOption == ACEPTAR_RETRUCO {
			message += common.BOLD + "(" + strconv.Itoa(ACEPTAR_RETRUCO) + ")" + common.NONE + common.GREEN + " Quiero REtruco " + common.NONE + "\n"
		} else if possibleOption == RECHAZAR_RETRUCO {
			message += common.BOLD + "(" + strconv.Itoa(RECHAZAR_RETRUCO) + ")" + common.NONE + common.RED + " Rechazar REtruco " + common.NONE + "\n"
		} else if possibleOption == NO_QUERER_ENVIDO_ENVIDO {
			message += common.BOLD + "(" + strconv.Itoa(NO_QUERER_ENVIDO_ENVIDO) + ")" + common.NONE + common.RED + " No quiero envido envido " + common.NONE + "\n"
		}
	}
	return message
}

func handleWaitingOptions(status int, player *Player, playerMove *int, playerError *PlayerError) {

	if status == VER_MIS_CARTAS {
		sendInfoCards(*player, playerError)
	}
	if status == IRSE_AL_MAZO {
		*playerMove = IRSE_AL_MAZO
		common.Send(player.socket, common.SingFinishRound)
		/* chequear error */ common.Receive(player.socket)
	}
	return
}

func receiveWaitingRequests(player *Player) (int, error) {
	common.Send(player.socket, common.WaitingOptionsPlayer)
	message, err := common.Receive(player.socket)
	if err != nil {
		return -1, err
	}
	if message == common.ACK {
		return 0, nil
	}
	option, _ := strconv.Atoi(message)
	return option, nil
}

func handleEnvidoResult(move *Move, actual *Player, opponent *Player, finish *bool) {
	sendInfoEnvido(move, actual, opponent)
	if (actual.lastMove == CANTAR_ENVIDO || opponent.lastMove == CANTAR_ENVIDO) && (actual.lastMove == QUERER_ENVIDO || opponent.lastMove == QUERER_ENVIDO) {
		envidoWinner := actual.verifyEnvidoWinnerAgainst(opponent)
		envidoWinner.sumPoints(2)
	} else if (actual.lastMove == CANTAR_ENVIDO || opponent.lastMove == CANTAR_ENVIDO) && (actual.lastMove == NO_QUERER_ENVIDO || opponent.lastMove == NO_QUERER_ENVIDO) {
		playerToSumPoints := actual
		if opponent.lastMove == CANTAR_ENVIDO {
			playerToSumPoints = opponent
		}
		playerToSumPoints.sumPoints(1)
	} else if (actual.lastMove == CANTAR_ENVIDO_ENVIDO || opponent.lastMove == CANTAR_ENVIDO_ENVIDO) && (actual.lastMove == QUERER_ENVIDO_ENVIDO || opponent.lastMove == QUERER_ENVIDO_ENVIDO) {
		envidoWinner := actual.verifyEnvidoWinnerAgainst(opponent)
		envidoWinner.sumPoints(4)
	} else if (actual.lastMove == CANTAR_ENVIDO_ENVIDO || opponent.lastMove == CANTAR_ENVIDO_ENVIDO) && (actual.lastMove == NO_QUERER_ENVIDO_ENVIDO || opponent.lastMove == NO_QUERER_ENVIDO_ENVIDO) {
		playerToSumPoints := actual
		if opponent.lastMove == CANTAR_ENVIDO_ENVIDO {
			playerToSumPoints = opponent
		}
		playerToSumPoints.sumPoints(2)
	}
}

func sendInfoEnvido(move *Move, actual *Player, opponent *Player) {
	if actual.lastMove == QUERER_ENVIDO_ENVIDO {
		SendInfoPlayer(opponent, common.OpponentAcceptEnvidoEnvido)
		move.envidoState = 0
	} else if actual.lastMove == NO_QUERER_ENVIDO_ENVIDO {
		SendInfoPlayer(opponent, common.OpponentRejectEnvidoEnvido)
		move.envidoState = 0
	} else if actual.lastMove == NO_QUERER_ENVIDO {
		SendInfoPlayer(opponent, common.OpponetRejectEnvido)
		move.envidoState = 0
	} else if actual.lastMove == CANTAR_ENVIDO {
		SendInfoPlayer(opponent, common.OpponetSingEnvido)
	} else if actual.lastMove == QUERER_ENVIDO {
		SendInfoPlayer(opponent, common.OpponetAcceptEnvido)
		move.envidoState = 0
	} else if actual.lastMove == CANTAR_ENVIDO_ENVIDO {
		SendInfoPlayer(opponent, common.OpponentSingEnvidoEnvido)
	}

}

func sendInfoOpponent(move *Move, player *Player) {
	if move.trucoState == CANTO_TRUCO {
		SendInfoPlayer(player, common.OpponentSingTruco)
	} else if move.trucoState == CANTAR_RETRUCO {
		SendInfoPlayer(player, common.OpponentSingRetruco)
	} else if move.trucoState == ACEPTAR_TRUCO && !player.notifyTruco {
		move.alreadyAceptedTruco = true
		player.setNotifyTruco(true)
		SendInfoPlayer(player, common.OpponetAcceptTruco)
	} else if move.trucoState == ACEPTAR_RETRUCO && !player.notifyRetruco {
		player.setNotifyRetruco(true)
		SendInfoPlayer(player, common.OpponetAcceptRetruco)
	} else if len(move.cardsPlayed) > 0 && move.envidoState == 0 {
		message := common.BBlue + "Tu oponente tiro una carta " + move.cardsPlayed[0].card.getFullName() + common.NONE + "\n"
		SendInfoPlayer(player, message)
	}

}

func envidoRelatedOptions(playerOption int, anotherPlayerOption int) bool {
	options := []int{CANTAR_ENVIDO, QUERER_ENVIDO, QUERER_ENVIDO_ENVIDO, NO_QUERER_ENVIDO_ENVIDO, NO_QUERER_ENVIDO, QUERER_ENVIDO_ENVIDO}

	for _, option := range options {
		if playerOption == option || anotherPlayerOption == option {
			return true
		}
	}
	return false
}

func trucoRelatedOptions(playerOption int, anotherPlayerOption int) bool {
	options := []int{CANTAR_TRUCO, CANTO_TRUCO, RECHAZAR_TRUCO, CANTAR_RETRUCO, RECHAZAR_RETRUCO, ACEPTAR_TRUCO, ACEPTAR_RETRUCO}
	for _, option := range options {
		if playerOption == option || anotherPlayerOption == option {
			return true
		}
	}
	return false
}
