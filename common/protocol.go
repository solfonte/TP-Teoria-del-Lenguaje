package common

import (
	"strconv"
)

const (
	WelcomeMessage = YELLOW + `
	_     _                           _     _               _   _                        
   | |   (_)                         (_)   | |             | | | |                       
   | |__  _  ___ _ ____   _____ _ __  _  __| | ___     __ _| | | |_ _ __ _   _  ___ ___  
   | '_ \| |/ _ \ '_ \ \ / / _ \ '_ \| |/ _  |/ _ \   / _  | | | __| '__| | | |/ __/ _ \ 
   | |_) | |  __/ | | \ V /  __/ | | | | (_| | (_) | | (_| | | | |_| |  | |_| | (_| (_) |
   |_.__/|_|\___|_| |_|\_/ \___|_| |_|_|\__,_|\___/   \__,_|_|  \__|_|   \__,_|\___\___/ 
																						 																				 
  ` + NONE
	AskPlayerName       = BWhite + "Porfavor Ingrese su nombre: " + NONE
	RequestMatchMessage = BWhite + "Ingresa " + BGreen + "CREATE" + BWhite + " para creare un juego o ingresa " + BBlue + "JOIN" + BWhite + " para unirte a una partida ya creada" + NONE

	AmountOfMembersMessage = BWhite + "ingrese cantidad integrantes: " + BCyan + "2 " + BWhite + "o " + BCyan + "4" + NONE
	DurationOfMatchMessage = BWhite + "Ingrese duracion de partida partida:" + BCyan + "15 " + BWhite + "o " + BCyan + "30 " + BWhite + "puntos" + NONE
	CreateMatchMessage     = BWhite + "OK, Partida creada, esperando a que se una el resto de los jugadores" + NONE
	JoinMatchMessage       = BWhite + "OK, Partida solicitada, se esta buscando una partida" + NONE
	GameStartedMessage     = BWhite + "El juego comenzó" + NONE
)

func GetWinningMessage(number int) string {
	return BGreen + "Ganaste la jugada " + strconv.Itoa(number) + NONE
}

func GetLossingMessage(number int) string {
	return BRed + "Perdiste la jugada " + strconv.Itoa(number) + NONE
}
