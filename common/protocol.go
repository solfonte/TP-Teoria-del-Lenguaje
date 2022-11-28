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
	FinishGame = YELLOW + `
	███████ ██ ███    ██ ██ ███████ ██   ██      ██████   █████  ███    ███ ███████ 
	██      ██ ████   ██ ██ ██      ██   ██     ██       ██   ██ ████  ████ ██      
	█████   ██ ██ ██  ██ ██ ███████ ███████     ██   ███ ███████ ██ ████ ██ █████   
	██      ██ ██  ██ ██ ██      ██ ██   ██     ██    ██ ██   ██ ██  ██  ██ ██      
	██      ██ ██   ████ ██ ███████ ██   ██      ██████  ██   ██ ██      ██ ███████ 
																					
																					
	` + NONE
	AskPlayerName       = BWhite + "Porfavor Ingrese su nombre: " + NONE
	RequestMatchMessage = BWhite + "Ingresa " + BGreen + "CREATE" + BWhite + " para creare un juego o ingresa " + BBlue + "JOIN" + BWhite + " para unirte a una partida ya creada" + NONE

	AmountOfMembersMessage  = BWhite + "ingrese cantidad integrantes: " + BCyan + "2 " + BWhite + "o " + BCyan + "4" + NONE
	DurationOfMatchMessage  = BWhite + "Ingrese duracion de partida partida:" + BCyan + "15 " + BWhite + "o " + BCyan + "30 " + BWhite + "puntos" + NONE
	CreateMatchMessage      = BWhite + "OK, Partida creada, esperando a que se una el resto de los jugadores" + NONE
	JoinMatchMessage        = BWhite + "OK, Partida solicitada, se esta buscando una partida" + NONE
	GameStartedMessage      = BWhite + "El juego comenzó" + NONE
	WaitingOptionsPlayer    = BWhite + "Mientras esperas a que sea tu turno, podes realizar las siguientes acciones" + "\n" + "(11) Irse al mazo" + "\n" + "(12) Consultar Cartas." + "\n" + "Ingresa (0) si no queres realizar ninguna de estas acciones" + NONE
	CardsMessage            = BWhite + "Estas son tus cartas: " + NONE
	WaitPlayerToPlayMessage = BBlue + "Espera a que juegue tu oponente..." + NONE + "\n"
	SingEnvido              = BPurple + "Cantaste ENVIDO" + NONE
	SingTruco               = BPurple + "Cantaste Truco" + NONE
	AcceptEnvido            = BPurple + "Aceptaste ENVIDO" + NONE
	AcceptTruco             = BPurple + "Aceptaste TRUCO" + NONE
	OpponentSingTruco       = BBlue + "Tu oponente canto TRUCO" + NONE + "\n"
	OpponetSingEnvido       = BBlue + "Tu oponente canto ENVIDO" + NONE + "\n"
	OpponetAcceptTruco      = BBlue + "Tu oponente Acepto el TRUCO" + NONE + "\n"
	OpponetAcceptEnvido     = BBlue + "Tu oponente Acepto el ENVIDO" + NONE + "\n"
	OpponetRejectTruco      = BBlue + "Tu oponente Rechazo el TRUCO" + NONE + "\n"
	WinMatchMessage         = BGreen + "Ganaste la partida :)" + NONE
	LoseMatchMessage        = BRed + "Perdiste la partida :(" + NONE
)

func GetWinningMoveMessage(number int) string {
	return BGreen + "Ganaste la jugada " + strconv.Itoa(number) + NONE
}

func GetWinningRoundMessage(number int) string {
	return BGreen + "Ganaste la ronda " + strconv.Itoa(number) + NONE
}

func GetLossingRoundMessage(number int) string {
	return BRed + "Perdiste la ronda " + strconv.Itoa(number) + NONE
}

func GetLossingMoveMessage(number int) string {
	return BRed + "Perdiste la jugada " + strconv.Itoa(number) + NONE
}

func GetPointsMessage(player1Points int, player2Points int) string {
	return BPurple + "Tus puntos son: " + strconv.Itoa(player1Points) + " Y los de tu oponente son: " + strconv.Itoa(player2Points) + NONE
}

func GetWinningEnvidoMessage(player1Points int, player2Points int) string {
	message := BPurple + "Ganaste el Envido con " + strconv.Itoa(player1Points) + " puntos"
	return message + " y tu oponente perdio con " + strconv.Itoa(player2Points) + " puntos"
}

func GetLossingEnvidoMessage(player1Points int, player2Points int) string {
	message := BPurple + "Perdiste el Envido con " + strconv.Itoa(player1Points) + " puntos"
	return message + " y tu oponente gano con " + strconv.Itoa(player2Points) + " puntos"
}
