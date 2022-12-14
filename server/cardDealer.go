package server

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

type CardDealer struct {
	cards [40]int
}

func (cardDealer *CardDealer) initialize() {
	for i := range cardDealer.cards {
		cardDealer.cards[i] = 1
	}
}

func (cardDealer *CardDealer) assignCards(player *Player) {
	cardNames := readCSV("../server/cards.csv")

	rand.Seed(time.Now().UnixNano())
	var amountOfCards int = 0
	var assignedCards []Card

	for amountOfCards < 3 {
		card := rand.Int() % 40
		if cardDealer.cards[card] != 0 {
			card_value, _ := strconv.Atoi(cardNames[card][0])
			card_suit := cardNames[card][1]

			assignedCards = append(assignedCards, Card{id: card, value: card_value, suit: card_suit})
			cardDealer.cards[card] = 0
			amountOfCards += 1
		}
	}
	player.dealCards(assignedCards)

}

func readCSV(filePath string) [][]string {

	f, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	csvReader.Comma = ','
	records, err := csvReader.ReadAll()
	if err != nil {
		fmt.Println("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}
