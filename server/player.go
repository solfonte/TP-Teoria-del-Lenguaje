package server

import (
	"fmt"
	"net"
	"truco/app/common"
)

type Player struct {
	id           int
	name         string
	socket       net.Conn
	points       int
	hand         Hand
	cardSelected Card
	winsPerPlay  int
}

func (player *Player) clearCards() {
	var cards []Card
	player.hand = Hand{cards: cards, cardsNotSelected: cards}
}

func (player *Player) getCards() []Card {
	return player.hand.cardsNotSelected
}

func (player *Player) verifyEnvidoWinnerAgainst(opponent *Player) *Player {
	if player.hand.winsOver(opponent.hand) {
		return player
	}
	return opponent
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

func (player *Player) dealCards(cards []Card) {
	player.hand = Hand{cards: cards, cardsNotSelected: cards}
}

func (player *Player) removeCardSelected(posToDelete int) {
	player.hand.removeCardSelected(posToDelete)
}

func (player *Player) sumPoints(points int) {
	fmt.Println("SUME PUNTOS POR ENVIDO ", points)
	player.points += points
}
func (player *Player) stop() {
	fmt.Println("player disconnect", player.name)
	player.socket.Close()
}
