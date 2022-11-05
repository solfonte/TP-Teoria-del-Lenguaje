package server

import (
	"fmt"
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

func processRequest(player Player, message string) map[string]int {
	match := make(map[string]int)
	if message == "CREATE" {
		fmt.Println("Entre a create")
		common.Send(player.socket, "Creando una partia, ingresar de cuantos integrantes 2 o 4")
		members, _ := common.Receive(player.socket)
		common.Send(player.socket, "Ingrese duracion de partida partida: 15 o 30 puntos")
		duration, _ := common.Receive(player.socket)
		fmt.Println("duration ", duration)
		fmt.Println("members ", members)
		match["create"] = 0
		match["members"], _ = strconv.Atoi(members)
		match["duration"], _ = strconv.Atoi(duration)
		common.Send(player.socket, "Partida creada, esperando a que se una el resto de los jugadores")
		return match
	} else {
		common.Send(player.socket, "Buscando una partida, de cuantos integrantes; 2 o 4")
		members, _ := common.Receive(player.socket)
		common.Send(player.socket, "Ingrese duracion de partida partida: 15 o 30 puntos")
		duration, _ := common.Receive(player.socket)
		fmt.Println("duration ", duration)
		fmt.Println("members ", members)
		match["create"] = 1
		match["members"], _ = strconv.Atoi(members)
		match["duration"], _ = strconv.Atoi(duration)
		common.Send(player.socket, "Partida solicitada, se esta buscando una partida")
		return match
	}

}
