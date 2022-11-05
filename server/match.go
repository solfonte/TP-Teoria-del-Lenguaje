package server

import (
	//"fmt"
	"math/rand"
    "time"
)

type Match struct {
	duration   int
	maxPlayers int
	players    []Player
	started    bool
}

func deal_cards(players []Player){
	
	//inicializamos un numero por carta
	var cards [40]int
	for i := range cards {
		cards[i] = 1
	}

	//random generator para asignar al azar
	rand.Seed(time.Now().UnixNano())

	//se le asignan cartas a cada jugador. No se pueden repetir. 
	for _,p := range players {

		var amountOfCards int = 0
		var assignedCards [3]int

		for amountOfCards < 3 {
			card := rand.Int() % 40
			if cards[card] != 0 {
				assignedCards[amountOfCards] = card

				//se le asigna cero para determinar que ya se repartio
				cards[card] = 0
				amountOfCards += 1
			}
		} 
		p.dealCards(assignedCards)
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
