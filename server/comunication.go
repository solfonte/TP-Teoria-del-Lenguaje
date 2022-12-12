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

	response := strings.ToUpper(messagePlayer)

	for (response != CREATE) && (response != JOIN) {
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

	for amout_duration_points != DURATION_15 && amout_duration_points != DURATION_30 {
		common.Send(player.socket, common.ErrorMaxPoints)
		duration, _ := common.Receive(player.socket)
		amout_duration_points, _ = strconv.Atoi(duration)

	}
	return amout_duration_points
}

func processRequest(player Player, message string) map[string]int {
	match := make(map[string]int)

	if message == CREATE {
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

func sendOtherPlayDisconnection(player Player, msg string) error {
	common.Send(player.socket, msg)
	_, err := common.Receive(player.socket)
	return err
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

func sendInfoPlayers(winner *Player, loser *Player, msgWinner string, msgLoser string, playerError *PlayerError) {
	common.Send(winner.socket, msgWinner)
	_, err := common.Receive(winner.socket)
	if err != nil {
		playerError.player = winner
		playerError.err = err
	}
	common.Send(loser.socket, msgLoser)
	_, err = common.Receive(loser.socket)
	if err != nil {
		playerError.player = loser
		playerError.err = err
	}
}

func SendWelcomeMessage(player *Player) {

	common.Send(player.socket, common.WelcomeMessage)
	common.Receive(player.socket)
	common.Send(player.socket, common.Rules)
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

func sendInfoPointsPlayers(player1 *Player, player2 *Player, playerError *PlayerError) {

	common.Send(player1.socket, common.GetPointsMessage(player1.points, player2.points))
	_, err := common.Receive(player1.socket)
	if err != nil {
		playerError.player = player1
		playerError.err = err
	}
	common.Send(player2.socket, common.GetPointsMessage(player2.points, player1.points))
	_, err = common.Receive(player2.socket)
	if err != nil {
		playerError.player = player2
		playerError.err = err
	}
}

func sendPlayerCardPlayed(player *Player, card Card, playerError *PlayerError) int {
	common.Send(player.socket, common.GetCardPlayed(getCardColors(card.getFullName())))
	_, err := common.Receive(player.socket)
	if err != nil {
		playerError.player = player
		playerError.err = err
		return -1
	}
	return 0
}

func SendInfoPlayer(player *Player, message string, playerError *PlayerError) {
	common.Send(player.socket, message)
	_, err := common.Receive(player.socket)
	if err != nil {
		playerError.player = player
		playerError.err = err
	}
}
