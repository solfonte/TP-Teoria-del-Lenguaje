package server

import (
	"fmt"
	"strconv"
	"truco/app/common"
)

func definePlayerPossibleOptions(move *Move, actualOption int, opponentOption int) []int {
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

func isTurnOfPlayer(player *Player) bool {
	fmt.Println("999999999999999999999999999999 PLAYER LAST MOVE = ", player.lastMove)
	return !(player.lastMove == CANTAR_ENVIDO_ENVIDO) && !(player.lastMove == CANTAR_RETRUCO)
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
		}else if possibleOption == ACEPTAR_RETRUCO {
			message += common.BOLD + "(" + strconv.Itoa(ACEPTAR_RETRUCO) + ")" + common.NONE + common.GREEN + " Quiero REtruco " + common.NONE + "\n"
		} else if possibleOption == RECHAZAR_RETRUCO {
			message += common.BOLD + "(" + strconv.Itoa(RECHAZAR_RETRUCO) + ")" + common.NONE + common.RED + " Rechazar REtruco " + common.NONE + "\n"
		}
	}
	return message
}

func handleWaitingOptions(status int, player *Player, playerMove *int, playerError *PlayerError) {

	if status == VER_MIS_CARTAS {
		fmt.Println("STATUS ES VER MIS CARTAS")
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
	fmt.Println("el hilo de receive waiting requests recibio " + message)
	fmt.Println(message)
	fmt.Println("pase waiting requests")
	option, _ := strconv.Atoi(message)
	return option, nil
}

func handleEnvidoResult(actualOption int, opponentOption int, actual *Player, opponent *Player, finish *bool) {
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
