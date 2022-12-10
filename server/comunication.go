package server

import (
	"fmt"
	"strconv"
	"strings"
	"truco/app/common"
)

func sendMenu(player Player) (string, error) {
	common.Send(player.socket, common.RequestMatchMessage)
	messagePlayer, error := common.Receive(player.socket)
	fmt.Println(messagePlayer)
	response := strings.ToUpper(messagePlayer)
	fmt.Println("RESPONSE QUE ME LLEGA ", response)

	for (response != "CREATE") && (response != "JOIN") {
		common.Send(player.socket, common.ErrorCreateOrJoin)
		messagePlayer, error = common.Receive(player.socket)
		response = strings.ToUpper(messagePlayer)
	}
	return response, error
}

func getAmountOfPoints(player Player) int {
	common.Send(player.socket, common.DurationOfMatchMessage)
	duration, _ := common.Receive(player.socket)
	amout_duration_points, _ := strconv.Atoi(duration)

	for amout_duration_points != 15 && amout_duration_points != 30 {
		common.Send(player.socket, common.ErrorMaxPoints)
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
	duration := getAmountOfPoints(player)
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

func sendInfoCards(player Player, playerError *PlayerError) {

	message := ""
	for _, card := range player.getCards() {
		message += getCardColors(card.getFullName()) + common.BWhite + " | " + common.NONE
	}
	common.Send(player.socket, common.CardsMessage+message)
	_, err := common.Receive(player.socket) //receive de patch (ok)

	if err != nil {
		playerError.player = &player
		playerError.err = err
	}
}

func sendInfoPlayers(winner *Player, loser *Player, msgWinner string, msgLoser string) {
	common.Send(winner.socket, msgWinner)
	msg, _ := common.Receive(winner.socket)
	fmt.Println(msg)

	common.Send(loser.socket, msgLoser)
	msg1, _ := common.Receive(loser.socket)
	fmt.Println(msg1)
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

func sendPlayerCardPlayed(player *Player, card Card) {
	common.Send(player.socket, common.GetCardPlayed(getCardColors(card.getFullName())))
	common.Receive(player.socket)
}

func SendInfoPlayer(player *Player, message string) {
	common.Send(player.socket, message)
	msg, _ := common.Receive(player.socket)
	fmt.Println("Le nvio dato especial y espero un OK recibo: ", msg)
}
