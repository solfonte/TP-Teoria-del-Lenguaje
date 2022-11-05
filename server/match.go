package server
/*
import (
	"fmt"
	"math/rand"
    "time"
)
*/
type Match struct {
	duration   int
	maxPlayers int
	players    []Player
	started    bool
}

func deal_cards(players []Player){
	
	var cardDealer  = CardDealer{}
	cardDealer.initialize()

	for _,p := range players {

		cardDealer.assignCards(&p)
	}
}

func (match *Match) addPlayerToMatch(player Player) {
	if match != nil {
		match.players = append(match.players, player)
		if len(match.players) == match.maxPlayers {
			match.started = true
			deal_cards(match.players)
		}
	}
}
