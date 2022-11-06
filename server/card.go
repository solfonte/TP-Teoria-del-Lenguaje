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
