package server

import (
	"strconv"
)

type Card struct {
	id    int
	value int
	suit  string
}

func (Card Card) getFullName() string {
	return (strconv.Itoa(Card.value) + " " + Card.suit)
}

// Devuelve 1 si la carta es mejor que la otra, 0 si son de igual jerarquia o -1 si es menor que la otra.
func (Card Card) compareCards(anotherCard Card) int {
	if Card.id >= 4 && Card.id <= 7 && anotherCard.id >= 4 && anotherCard.id <= 7 {
		return 0
	}
	if Card.id >= 8 && Card.id <= 11 && anotherCard.id >= 8 && anotherCard.id <= 11 {
		return 0
	}
	if Card.id >= 12 && Card.id <= 13 && anotherCard.id >= 12 && anotherCard.id <= 13 {
		return 0
	}
	if Card.id >= 14 && Card.id <= 17 && anotherCard.id >= 14 && anotherCard.id <= 17 {
		return 0
	}
	if Card.id >= 18 && Card.id <= 21 && anotherCard.id >= 18 && anotherCard.id <= 21 {
		return 0
	}
	if Card.id >= 22 && Card.id <= 25 && anotherCard.id >= 22 && anotherCard.id <= 25 {
		return 0
	}
	if Card.id >= 26 && Card.id <= 27 && anotherCard.id >= 26 && anotherCard.id <= 27 {
		return 0
	}
	if Card.id >= 28 && Card.id <= 31 && anotherCard.id >= 28 && anotherCard.id <= 31 {
		return 0
	}
	if Card.id >= 32 && Card.id <= 35 && anotherCard.id >= 32 && anotherCard.id <= 35 {
		return 0
	}
	if Card.id >= 36 && Card.id <= 39 && anotherCard.id >= 36 && anotherCard.id <= 39 {
		return 0
	}
	if Card.id < anotherCard.id {
		return 1
	} else {
		return -1
	}
}
