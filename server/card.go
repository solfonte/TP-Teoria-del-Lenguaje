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
	if Card.id >= RANGO_I_TRES && Card.id <= RANGO_F_TRES && anotherCard.id >= RANGO_I_TRES && anotherCard.id <= RANGO_F_TRES {
		return 0
	}
	if Card.id >= RANGO_I_DOS && Card.id <= RANGO_F_DOS && anotherCard.id >= RANGO_I_DOS && anotherCard.id <= RANGO_F_DOS {
		return 0
	}
	if Card.id >= RANGO_I_UNO && Card.id <= RANGO_F_UNO && anotherCard.id >= RANGO_I_UNO && anotherCard.id <= RANGO_F_UNO {
		return 0
	}
	if Card.id >= RANGO_I_DOCE && Card.id <= RANGO_F_DOCE && anotherCard.id >= RANGO_I_DOCE && anotherCard.id <= RANGO_F_DOCE {
		return 0
	}
	if Card.id >= RANGO_I_ONCE && Card.id <= RANGO_F_ONCE && anotherCard.id >= RANGO_I_ONCE && anotherCard.id <= RANGO_F_ONCE {
		return 0
	}
	if Card.id >= RANGO_I_DIEZ && Card.id <= RANGO_F_DIEZ && anotherCard.id >= RANGO_I_DIEZ && anotherCard.id <= RANGO_F_DIEZ {
		return 0
	}
	if Card.id >= RANGO_I_SIETE && Card.id <= RANGO_F_SIETE && anotherCard.id >= RANGO_I_SIETE && anotherCard.id <= RANGO_F_SIETE {
		return 0
	}
	if Card.id >= RANGO_I_SEIS && Card.id <= RANGO_F_SEIS && anotherCard.id >= RANGO_I_SEIS && anotherCard.id <= RANGO_F_SEIS {
		return 0
	}
	if Card.id >= RANGO_I_CINCO && Card.id <= RANGO_F_CINCO && anotherCard.id >= RANGO_I_CINCO && anotherCard.id <= RANGO_F_CINCO {
		return 0
	}
	if Card.id >= RANGO_I_CUATRO && Card.id <= RANGO_F_CUATRO && anotherCard.id >= RANGO_I_CUATRO && anotherCard.id <= RANGO_F_CUATRO {
		return 0
	}
	if Card.id < anotherCard.id {
		return 1
	} else {
		return -1
	}
}
