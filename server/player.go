package server

import (
	"fmt"
	"net"
	"truco/app/common"
)

type Player struct {
	id     int
	name   string
	socket net.Conn
	points int
}

func (player *Player) askPlayerName() {
	common.Send(player.socket, "Podes ingresar tu nombre")
	playerName, error := common.Receive(player.socket)
	if error != nil {
		fmt.Println("Error reciving from client: ", error.Error())
	}
	player.name = playerName
	fmt.Println("nombre del jugador: ", player.name)
}
