package server

import (
	"fmt"
	"strconv"
	"strings"
	"truco/app/common"
)

func sendMenu(player Player) (string, error) {
	common.Send(player.socket, "Bienvenido al truco "+player.name)
	messagePlayer, error := common.Receive(player.socket)
	common.Send(player.socket, "Las reglas del juego son sencillas: .....")
	messagePlayer, error = common.Receive(player.socket)
	common.Send(player.socket, "ingresa CREATE para crear un juego O ingresa JOIN para unirte a una partida ya creada")
	fmt.Println("paso")
	// receives its answer
	messagePlayer, error = common.Receive(player.socket)
	fmt.Println(messagePlayer)
	response := strings.ToUpper(messagePlayer)
	fmt.Println(response)

	for (response != "CREATE") && (response != "JOIN") {
		common.Send(player.socket, "Error: ingrese CREATE para crear un juego O ingresa JOIN para unirte a una partida ya creada")
		messagePlayer, error = common.Receive(player.socket)
		response = strings.ToUpper(messagePlayer)
	}
	return response, error
}

func getAmountOfPlayers(player Player) int {

	common.Send(player.socket, "ingrese cantidad integrantes; 2 o 4")
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
	common.Send(player.socket, "Ingrese duracion de partida partida: 15 o 30 puntos")
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
		common.Send(player.socket, "OK, Partida creada, esperando a que se una el resto de los jugadores")
		return match
	} else {
		match["create"] = 1
		getMatchParameters(match, player)
		common.Send(player.socket, "OK, Partida solicitada, se esta buscando una partida")
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
	common.Send(player.socket, "El juego comenzó")
	message, _ := common.Receive(player.socket)
	fmt.Println(message)

}

func sendOtherPlayDisconnection(player Player, msg string) {
	common.Send(player.socket, msg)
	message, _ := common.Receive(player.socket)
	fmt.Println(message)
}

func sendInfoCards(player Player) {

	//TODO: ver si no conviene que sea dinamico para cuando le queden dos o una?
	cards := player.getCards()//hand.cards[0].getFullName()
	card1 := cards[0].getFullName()
	card2 := cards[1].getFullName()
	card3 := cards[2].getFullName()
	common.Send(player.socket, "Estas son tus cartas: "+card1+" "+card2+" "+card3)
	message, _ := common.Receive(player.socket)
	fmt.Println(message)
	fmt.Println("cartas: "+card1, card2, card3)

}

func sendInfoPlayers(winner *Player, loser *Player, msgWinner string, msgLoser string) {
	common.Send(winner.socket, msgWinner)
	common.Receive(winner.socket)

	common.Send(loser.socket, msgLoser)
	common.Receive(loser.socket)
}
