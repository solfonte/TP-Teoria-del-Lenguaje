package server

import (
	"fmt"
	"net"
	"truco/app/common"
)

type Player struct {
	id             int
	name           string
	socket         net.Conn
	points         int
	hand           Hand
	cardSelected   Card
	winsPerPlay    int
	hasSagnTruco   bool
	lastMove       int
	hasSangReTruco bool
	notifyTruco    bool
	notifyRetruco  bool
}

func (player *Player) clearCards() {
	var cards []Card
	player.hand = Hand{cards: cards, cardsNotSelected: cards}
}

func (player *Player) getCards() []Card {
	return player.hand.cardsNotSelected
}

func (player *Player) verifyEnvidoWinnerAgainst(opponent *Player) *Player {
	if player.hand.winsEnvidoOver(opponent.hand) {
		msgWinning := common.GetWinningEnvidoMessage(player.hand.calculatePointsEnvido(), opponent.hand.calculatePointsEnvido())
		msgLossing := common.GetLossingEnvidoMessage(opponent.hand.calculatePointsEnvido(), player.hand.calculatePointsEnvido())
		sendInfoPlayers(player, opponent, msgWinning, msgLossing)
		return player
	}
	msgWinning := common.GetWinningEnvidoMessage(opponent.hand.calculatePointsEnvido(), player.hand.calculatePointsEnvido())
	msgLossing := common.GetLossingEnvidoMessage(player.hand.calculatePointsEnvido(), opponent.hand.calculatePointsEnvido())
	sendInfoPlayers(opponent, player, msgWinning, msgLossing)
	return opponent
}

func (player *Player) welcomePlayer() {
	SendWelcomeMessage(player)
	player.askPlayerName()
}

func (player *Player) askPlayerName() {
	common.Send(player.socket, common.AskPlayerName)
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

func (player *Player) setHasSangTruco(truco bool) {
	player.hasSagnTruco = truco
}

func (player *Player) setHasSangRetruco(retruco bool) {
	player.hasSangReTruco = retruco
}

func (player *Player) setNotifyTruco(notify bool) {
	player.notifyTruco = notify
}

func (player *Player) setNotifyRetruco(notify bool) {
	player.notifyRetruco = notify
}
