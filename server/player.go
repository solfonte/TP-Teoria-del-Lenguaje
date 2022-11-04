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

func askPlayerName(player Player) {
	common.Send(player.socket, "Podes ingresar tu nombre")
	playerName, error := common.Receive(player.socket)
	if error != nil {
		fmt.Println("Error reciving from client: ", error.Error())
	}
	player.name = playerName
}
