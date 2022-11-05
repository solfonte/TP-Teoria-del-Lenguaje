package server

import (
	"strconv"
	"truco/app/common"
)

func sendMenu(player Player) (string, error) {
	common.Send(player.socket, "Bienvenido al truco " + player.name)
	common.Receive(player.socket)
	common.Send(player.socket, "Las reglas del juego son sencillas: .....")
	common.Receive(player.socket)
	common.Send(player.socket, "ingresa CREATE para crear un juego O ingresa JOIN para unirte a una partida ya creada")
	
	// receives its answer
	messagePlayer, error := common.Receive(player.socket)
	for (messagePlayer != "CREATE") && (messagePlayer != "JOIN") {
		common.Send(player.socket, "ingresa CREATE para crear un juego O ingresa JOIN para unirte a una partida ya creada")
		messagePlayer, error = common.Receive(player.socket)
	}
	return messagePlayer, error
}

func getAmountOfPlayers(player Player) int{
	
	common.Send(player.socket, "Buscando una partida, de cuantos integrantes; 2 o 4")
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
		common.Send(player.socket, "Partida creada, esperando a que se una el resto de los jugadores")
		return match
	} else {
		match["create"] = 1
		getMatchParameters(match, player)
		common.Send(player.socket, "Partida solicitada, se esta buscando una partida")
		return match
	}
		
}

func getMatchParameters(match map[string]int, player Player){
	members := getAmountOfPlayers(player)
	duration:= getAmountOfPoints(player)
	match["members"] = members
	match["duration"] = duration
}

func startGame(player Player){
	common.Send(player.socket, "El juego ya arrancado")
	common.Send(player.socket, "Estas son tus cartas")
	card1 := strconv.Itoa(player.cards[0].value) + " " strconv.Itoa(player.cards[0].type)
	card2 := strconv.Itoa(player.cards[1].value) + " " strconv.Itoa(player.cards[1].type)
	card3 := strconv.Itoa(player.cards[2].value) + " " strconv.Itoa(player.cards[2].type)
	common.Send(player.socket, card1 + " " + card2 + " " + card3)
}
