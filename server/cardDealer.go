package server

import (
	"math/rand"
    "time"
)

type CardDealer struct {
	cards [40]int
}

func (cardDealer *CardDealer) initialize (){
	for i := range cardDealer.cards {
		cardDealer.cards[i] = 1
	}
}

func (cardDealer *CardDealer) assignCards (player *Player) {
	//random generator para asignar al azar
	rand.Seed(time.Now().UnixNano())
	var amountOfCards int = 0
		var assignedCards [3]int

		for amountOfCards < 3 {
			card := rand.Int() % 40
			if cardDealer.cards[card] != 0 {
				assignedCards[amountOfCards] = card

				//se le asigna cero para determinar que ya se repartio
				cardDealer.cards[card] = 0
				amountOfCards += 1
			}
		} 
		player.dealCards(assignedCards)
}