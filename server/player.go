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
	cards [3]int
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

func (player *Player) dealCards(cards [3]int){
	//despues hay que mapear los numeros a las cartas. Por ahora queda asi
	player.cards = cards
}
