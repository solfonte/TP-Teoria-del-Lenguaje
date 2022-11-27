package server

import (
	"fmt"
	"strconv"
	"strings"
	"truco/app/common"
)

func sendMenu(player Player) (string, error) {
	// msgCreate := common.GREEN + "CREATE" + common.NONE
	// //fmt.Println(msgCreate)
	// msgJoin := common.BLUE + "JOIN" + common.NONE
	// //fmt.Println(msgJoin)
	// message := "ingresa " + msgCreate + " para crear un juego O ingresa " + msgJoin + " para unirte a una partida ya creada"
	common.Send(player.socket, common.RequestMatchMessage)
	// receives its answer
	messagePlayer, error := common.Receive(player.socket)
	fmt.Println(messagePlayer)
	response := strings.ToUpper(messagePlayer)
	fmt.Println("RESPONSE QUE ME LLEGA ", response)

	for (response != "CREATE") && (response != "JOIN") {
		common.Send(player.socket, "Error: ingrese CREATE para crear un juego O ingresa JOIN para unirte a una partida ya creada")
		messagePlayer, error = common.Receive(player.socket)
		response = strings.ToUpper(messagePlayer)
	}
	return response, error
}

func getAmountOfPlayers(player Player) int {

	common.Send(player.socket, common.AmountOfMembersMessage)
	members, _ := common.Receive(player.socket)
	amount_members, _ := strconv.Atoi(members)

	for amount_members != 2 && amount_members != 4 {
		common.Send(player.socket, " Error! Elegir cantidad de integrantes 2 o 4")
		members, _ = common.Receive(player.socket)
		amount_members, _ = strconv.Atoi(members)
	}
	return amount_members
}

func getAmountOfPoints(player Player) int {
	common.Send(player.socket, common.DurationOfMatchMessage)
	duration, _ := common.Receive(player.socket)
	amout_duration_points, _ := strconv.Atoi(duration)

	for amout_duration_points != 15 && amout_duration_points != 30 {
		common.Send(player.socket, "Error! ingrese duracion de partida de valor 15 o 30 puntos")
		duration, _ := common.Receive(player.socket)
		amout_duration_points, _ = strconv.Atoi(duration)

	}
	return amout_duration_points
}

func processRequest(player Player, message string) map[string]int {
	match := make(map[string]int)

	if message == "CREATE" {
		match["create"] = 0
		getMatchParameters(match, player)
		common.Send(player.socket, common.CreateMatchMessage)
		return match
	} else {
		match["create"] = 1
		getMatchParameters(match, player)
		common.Send(player.socket, common.JoinMatchMessage)
		return match
	}

}

func getMatchParameters(match map[string]int, player Player) {
	members := getAmountOfPlayers(player)
	duration := getAmountOfPoints(player)
	match["members"] = members
	match["duration"] = duration
}

func startGame(player Player) {
	common.Send(player.socket, common.GameStartedMessage)
	message, _ := common.Receive(player.socket)
	fmt.Println(message)

}

func sendOtherPlayDisconnection(player Player, msg string) {
	common.Send(player.socket, msg)
	message, _ := common.Receive(player.socket)
	fmt.Println(message)
}

func getCardColors(card string) string {
	if strings.Contains(card, "oro") {
		return (common.YELLOW + card + common.NONE)
	} else if strings.Contains(card, "espada") {
		return (common.CYAN + card + common.NONE)
	} else if strings.Contains(card, "basto") {
		return (common.GREEN + card + common.NONE)
	} else {
		return (common.RED + card + common.NONE)
	}
}

func sendInfoCards(player Player) {

	message := ""
	for _, card := range player.getCards() {
		message += getCardColors(card.getFullName()) + common.BWhite + " | " + common.NONE
	}
	common.Send(player.socket, common.CardsMessage+message)
	common.Receive(player.socket) //receive de patch (ok)
}

func sendInfoPlayers(winner *Player, loser *Player, msgWinner string, msgLoser string) {
	common.Send(winner.socket, msgWinner)
	msg, _ := common.Receive(winner.socket)
	fmt.Println(msg)

	common.Send(loser.socket, msgLoser)
	msg, _ = common.Receive(loser.socket)
	fmt.Println(msg)
}

func SendWelcomeMessage(player *Player) {

	common.Send(player.socket, common.WelcomeMessage)
	common.Receive(player.socket)
	common.Send(player.socket, "Las reglas del juego son: ")
	common.Receive(player.socket)
}

func GetCardsToThrow(cards []Card) (string, []int) {

	message := common.BWhite + "Que carta queres tirar? " + "\n" + common.NONE

	var maxOptionsSelected []int
	for index, card := range cards {
		number := common.BOLD + strconv.Itoa(index+1) + ") " + common.NONE
		message += number
		message += getCardColors(card.getFullName()) + "\n"
		maxOptionsSelected = append(maxOptionsSelected, index+1)
	}
	return message, maxOptionsSelected

}

func sendInfoPointsPlayers(player1 *Player, player2 *Player) {

	common.Send(player1.socket, common.GetPointsMessage(player1.points, player2.points))
	common.Receive(player1.socket)
	common.Send(player2.socket, common.GetPointsMessage(player2.points, player1.points))
	common.Receive(player2.socket)
}
