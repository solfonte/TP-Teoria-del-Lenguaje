package server

import (
	"math/rand"
    "time"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
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

	cardNames := readCSV("server/cards.csv")

	rand.Seed(time.Now().UnixNano())
	var amountOfCards int = 0
	var assignedCards [3]Card

		for amountOfCards < 3 {
			card := rand.Int() % 40 // posicion en nuestro vector de cartas 
			if cardDealer.cards[card] != 0 { // si no fue asignada 
				card_value, _ := strconv.Atoi(cardNames[card][0])
				card_suit := cardNames[card][1]

				assignedCards[amountOfCards] = Card{id : card, value : card_value, suit : card_suit }

				//se le asigna cero para determinar que ya se repartio
				cardDealer.cards[card] = 0
				amountOfCards += 1
			}
		} 
		player.dealCards(assignedCards)
		
}

func readCSV(filePath string) [][]string{
	f, err := os.Open(filePath)
    if err != nil {
       // log.Fatal("Unable to read input file " + filePath, err)
       fmt.Println("Unable to read input file " + filePath, err)
    }
    defer f.Close()

	csvReader := csv.NewReader(f)
	csvReader.Comma = ','
    records, err := csvReader.ReadAll()
    if err != nil {
        fmt.Println("Unable to parse file as CSV for " + filePath, err)
    }

	return records
}
