package server

import (
	"fmt"
	"math"
	"strconv"
	"truco/app/common"
)

const (
	opponentMessageForEnvido = "Tu oponente canto envido. Tus opciones son: (1) Quiero (2) Quiero envido envido (3) No quiero"
	TIRAR_CARTA				 = 4
	CANTAR_ENVIDO			 = 5
	CANTAR_TRUCO			 = 6
	QUERER_ENVIDO            = 7
	QUERER_ENVIDO_ENVIDO     = 8
	NO_QUERER_ENVIDO_ENVIDO  = 9
	NO_QUERER_ENVIDO         = 10
)

type InfoPlayer struct {
	id     int
	points int
}

type Move struct {
	winner            InfoPlayer
	loser             InfoPlayer
	points            int
	typeMove          int
	cardsPlayed       []Card
	alreadySangEnvido bool
} 

func (move *Move) canSingEnvido() bool {
	return move.typeMove == 1 && !move.alreadySangEnvido
}

func (move *Move) definePlayerPossibleOptions(opponentOption int) []int {
	var options []int
	if (opponentOption == TIRAR_CARTA){
		options = append(options, TIRAR_CARTA)
		if (move.canSingEnvido()){
			options = append(options, CANTAR_ENVIDO)
		}
		options = append(options, CANTAR_TRUCO)
	} else if (opponentOption == CANTAR_ENVIDO){
		options = append(options, QUERER_ENVIDO)
		options = append(options, QUERER_ENVIDO_ENVIDO)
		options = append(options, NO_QUERER_ENVIDO)
	}else if (opponentOption == QUERER_ENVIDO || opponentOption == NO_QUERER_ENVIDO || opponentOption == NO_QUERER_ENVIDO_ENVIDO){
		options = append(options, TIRAR_CARTA)
	}else if (opponentOption == QUERER_ENVIDO_ENVIDO){
		options = append(options, TIRAR_CARTA)
		options = append(options, NO_QUERER_ENVIDO)
	}else {
		options = append(options, TIRAR_CARTA)
		if (move.canSingEnvido()){
			options = append(options, CANTAR_ENVIDO)
		}
		options = append(options, CANTAR_TRUCO)
	}
	return options
	//TODO: DESPUES SE AGREGA TRUCO
}

func (move *Move) handleResult(option1 int, option2 int, actual *Player, opponent *Player, finish *bool) bool {
	if option1 == QUERER_ENVIDO || option2 == QUERER_ENVIDO {
		fmt.Println("alguno quiere envido")
		//TODO: el oponent es el q no es mano???? importante
		envidoWinner := actual.verifyEnvidoWinnerAgainst(opponent)
		envidoWinner.sumPoints(2)
		return false
	} else if option1 == NO_QUERER_ENVIDO {
		fmt.Println("Uno no quiere envido")
		opponent.sumPoints(1)
		return false
	} else if option2 == NO_QUERER_ENVIDO {
		fmt.Println("Dos no quiere envido")
		actual.sumPoints(1)
		return false
	} else if option1 == TIRAR_CARTA && option2 == TIRAR_CARTA {
		fmt.Println("ambos tiraron carta")
		result := move.cardsPlayed[0].compareCards(move.cardsPlayed[1])
		return move.assingWinner(result, actual, opponent, finish)
	} else if option1 == CANTAR_ENVIDO || option2 == CANTAR_ENVIDO {
		return false
	} 
	return true
}

func (move *Move) start_move(player1 *Player, player2 *Player, playerError *PlayerError, finish *bool) int {
	var moveFinished bool
	err := 0
	var option1 int = 0 
	var option2 int = 0
	for !moveFinished && err != -1 {
		fmt.Println("entro a una nueva vuelta de 3 tirosss y el type move es ", strconv.Itoa(move.typeMove))
		err = move.askPlayerForWait(player2, playerError)
		fmt.Println("el jugador 2 espera")

		if err != -1{
			fmt.Println("juega el jugador 1")
			options := move.definePlayerPossibleOptions(option2)
			option1, err = move.askPlayerForMove(player1, options, playerError)
		}
		if err != -1 {
			fmt.Println("manejamos el resultado de opciones ingresadas que son: ", option1, " and ", option2)
			moveFinished = move.handleResult(option1, option2, player1, player2, finish)
			fmt.Println("Finish: ", moveFinished)
		}
		if err != -1 {
		fmt.Println("el jugador 1 espera")

			err = move.askPlayerForWait(player1, playerError)
		}
		if err != -1{
			fmt.Println("juega el jugador 1")
			options := move.definePlayerPossibleOptions(option1)
			option2, err = move.askPlayerForMove(player2, options, playerError)
		}
		if err != -1 {
			fmt.Println("manejamos el resultado de opciones ingresadas que son: ", option1, " and ", option2)
			moveFinished = move.handleResult(option1, option2, player1, player2, finish)
			fmt.Println("Finish: ", moveFinished)
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
		move.winner.points = 1
		winner.points += 1
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
		*finish = true
	} else {
		fmt.Println("No termino jugada")
		*finish = false
	}
	// termino jugada
	return true
}

func (move *Move) askPlayerForWait(player *Player, playerError *PlayerError) int {
	common.Send(player.socket, "Espera a que juegue tu oponente...")
	_, err := common.Receive(player.socket)
	if err != nil {
		playerError.player = player
		playerError.err = err
		return -1
	}
	//fmt.Println(message)
	return 0
}



func (move *Move) handleEnvido(player *Player) {
	common.Send(player.socket, "cantaste ENVIDO")
	move.alreadySangEnvido = true
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
	player.removeCardSelected(option - 1)
	return 0
}

func (move *Move) sendInfoMove(player *Player, options []int,  playerError *PlayerError) (int, int) {

	message := "Es tu turno, podes hacer las siguientes jugadas: "
	for _, possibleOption := range options {
		if possibleOption == TIRAR_CARTA{
			message += "(" + strconv.Itoa(TIRAR_CARTA) + ") Tirar una carta "
		}else if possibleOption == CANTAR_ENVIDO {
			message += "(" + strconv.Itoa(CANTAR_ENVIDO) + ") Cantar envido "
		}else if possibleOption == CANTAR_TRUCO {
			message += "(" + strconv.Itoa(CANTAR_TRUCO) + ") Cantar truco "
		}else if possibleOption == QUERER_ENVIDO {
			message += "(" + strconv.Itoa(QUERER_ENVIDO) + ") Quiero envido "
		}else if possibleOption == QUERER_ENVIDO_ENVIDO {
			message += "(" + strconv.Itoa(QUERER_ENVIDO_ENVIDO) + ") Quiero envido envido "
		}else if possibleOption == NO_QUERER_ENVIDO{
			message += "(" + strconv.Itoa(NO_QUERER_ENVIDO) + ") No quiero envido "
		}
	}
	//esto va en otra funcion
	option := 0
	msgError := ""
	//TODO: tanto en sendInfo move como en handlethwro se hace el mismo loop, ver de meterlos
	// en una misma funcion.(recibea mensaje y cant de optiones y devuelvan la opcion elegida)
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
		fmt.Println("opcion ingresada: ", option)
	}
	return option, 0
}


func (move *Move) askPlayerForMove(player *Player, options []int, playerError *PlayerError) (int, int) {
	option := 0
	var err int

	if len(move.cardsPlayed) > 0 {
		message := "Tu oponente tiro una carta " + move.cardsPlayed[0].getFullName() //TODO:porque 0?
		common.Send(player.socket, message)
		//messageClient, _ := common.Receive(player.socket)
	}

	option, err = move.sendInfoMove(player, options, playerError)

	if err == -1 {
		return -1, -1
	}

	switch option {
		case TIRAR_CARTA:
			err = move.handleThrowACard(player, playerError)
		case CANTAR_ENVIDO:
			move.handleEnvido(player)
		case CANTAR_TRUCO:
			fmt.Println("canto truco")
		case QUERER_ENVIDO:	// si quiere envido es porque alguien lo canto
			fmt.Println("QUIERE ENVIDO")
		case QUERER_ENVIDO_ENVIDO: // si quiere envido envido es porque alguien canto envido
			fmt.Println("QUIERE ENVIDO ENVIDO")
		case NO_QUERER_ENVIDO:	// si no quiere envido es porque alguien lo canto
			fmt.Println("NO QUIERE ENVIDO")
	}
	return option, err
}
