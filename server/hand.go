package server

import (
	"fmt"
)

type Hand struct {
	cards            []Card
	cardsNotSelected []Card
}

const (
	BASTO  = 0
	ORO    = 1
	COPA   = 2
	ESPADA = 3
)

func (hand *Hand) removeCardSelected(posTodelete int) {
	hand.cardsNotSelected = append(hand.cardsNotSelected[:posTodelete], hand.cardsNotSelected[posTodelete+1:]...)
}

func (hand *Hand) distributeSuits() [4][]Card {
	var basto []Card
	var oro []Card
	var copa []Card
	var espada []Card
	suits := [4][]Card{basto, oro, copa, espada}

	for _, card := range hand.cards {
		if card.suit == "basto" {
			suits[BASTO] = append(suits[BASTO], card)
		} else if card.suit == "oro" {
			suits[ORO] = append(suits[ORO], card)
		} else if card.suit == "copa" {
			suits[COPA] = append(suits[COPA], card)
		} else {
			suits[ESPADA] = append(suits[ESPADA], card)
		}
	}
	return suits

}

func (hand *Hand) pointsForSuit() (int, []Card) {

	var repeatedSuit []Card
	points := 0
	suits := hand.distributeSuits()

	for _, suit := range suits {
		if len(suit) >= 2 {
			repeatedSuit = suit
			points += 20
		}
	}

	return points, repeatedSuit
}

func (hand *Hand) pointsForNumberEnvido(repeatedSuit []Card) int {
	points := 0
	if repeatedSuit != nil {
		greatestCardNumber := 0
		secondGreatestCardNumber := 0

		if repeatedSuit[0].value < 10 {
			greatestCardNumber = repeatedSuit[0].value
		}
		if repeatedSuit[1].value < 10 {
			secondGreatestCardNumber = repeatedSuit[1].value
		}

		if len(repeatedSuit) == 3 && repeatedSuit[2].value < 10 {
			if greatestCardNumber > repeatedSuit[2].value && secondGreatestCardNumber < repeatedSuit[2].value {
				secondGreatestCardNumber = repeatedSuit[2].value
			} else if greatestCardNumber < repeatedSuit[2].value && secondGreatestCardNumber > repeatedSuit[2].value {
				greatestCardNumber = repeatedSuit[2].value
			} else if greatestCardNumber < repeatedSuit[2].value && secondGreatestCardNumber < repeatedSuit[2].value {
				if greatestCardNumber > secondGreatestCardNumber {
					secondGreatestCardNumber = repeatedSuit[2].value
				} else {
					greatestCardNumber = repeatedSuit[2].value
				}
			}
		}
		points += greatestCardNumber + secondGreatestCardNumber
	} else {
		greatestCardNumber := hand.cards[0].value
		for _, card := range hand.cards {
			if card.value > greatestCardNumber {
				greatestCardNumber = card.value
			}
		}
		points = greatestCardNumber
	}
	return points
}

func (hand *Hand) calculatePointsEnvido() int {
	pointsForSuit, repeatedSuit := hand.pointsForSuit()
	pointsForNumber := hand.pointsForNumberEnvido(repeatedSuit)
	return pointsForNumber + pointsForSuit
}

func (hand *Hand) winsEnvidoOver(otherHand Hand) bool {
	fmt.Println("entre a wins over")
	sumForHand := hand.calculatePointsEnvido()
	sumForOtherHand := otherHand.calculatePointsEnvido()

	if sumForHand >= sumForOtherHand {
		//TODO: si es empate gana el que es mano en la ronda
		return true
	}
	return false
}
