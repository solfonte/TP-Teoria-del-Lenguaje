package common

import (
	"strconv"
)

const (
	WelcomeMessage = YELLOW + `
	_     _                           _     _               _   _                        
   | |   (_)                         (_)   | |             | | | |                       
   | |__  _  ___ _ ____   _____ _ __  _  __| | ___     __ _| | | |_ _ __ _   _  ___ ___  
   |  _ \| |/ _ \  _ \ \ / / _ \  _ \| |/ _  |/ _ \   / _  | | | __|  __| | | |/ __/ _ \ 
   | |_) | |  __/ | | \ V /  __/ | | | | (_| | (_) | | (_| | | | |_| |  | |_| | (_| (_) |
   |_.__/|_|\___|_| |_|\_/ \___|_| |_|_|\__,_|\___/   \__,_|_|  \__|_|   \____|\___\___/ 
																						 																				 
  ` + NONE
	FinishGame = YELLOW + `
	███████ ██ ███    ██ ██ ███████ ██   ██      ██████   █████  ███    ███ ███████ 
	██      ██ ████   ██ ██ ██      ██   ██     ██       ██   ██ ████  ████ ██      
	█████   ██ ██ ██  ██ ██ ███████ ███████     ██   ███ ███████ ██ ████ ██ █████   
	██      ██ ██  ██ ██ ██      ██ ██   ██     ██    ██ ██   ██ ██  ██  ██ ██      
	██      ██ ██   ████ ██ ███████ ██   ██      ██████  ██   ██ ██      ██ ███████ 
																					
																					
	` + NONE
	WinMatchMessage = BGreen + `
	╔═╗┌─┐┌┐┌┌─┐┌─┐┌┬┐┌─┐  ┬  ┌─┐  ┌─┐┌─┐┬─┐┌┬┐┬┌┬┐┌─┐  
	║ ╦├─┤│││├─┤└─┐ │ ├┤   │  ├─┤  ├─┘├─┤├┬┘ │ │ ││├─┤  
	╚═╝┴ ┴┘└┘┴ ┴└─┘ ┴ └─┘  ┴─┘┴ ┴  ┴  ┴ ┴┴└─ ┴ ┴─┴┘┴ ┴  
	` + NONE
	LoseMatchMessage = BRed + `
	╔═╗┌─┐┬─┐┌┬┐┬┌─┐┌┬┐┌─┐  ┬  ┌─┐  ┌─┐┌─┐┬─┐┌┬┐┬┌┬┐┌─┐
	╠═╝├┤ ├┬┘ │││└─┐ │ ├┤   │  ├─┤  ├─┘├─┤├┬┘ │ │ ││├─┤
	╩  └─┘┴└──┴┘┴└─┘ ┴ └─┘  ┴─┘┴ ┴  ┴  ┴ ┴┴└─ ┴ ┴─┴┘┴ ┴
	` + NONE
	AskPlayerName       = BWhite + "Porfavor Ingrese su nombre: " + NONE
	RequestMatchMessage = BWhite + "Ingresa " + BGreen + "CREATE" + BWhite + " para creare un juego o ingresa " + BBlue + "JOIN" + BWhite + " para unirte a una partida ya creada" + NONE

	DurationOfMatchMessage  = BWhite + "Ingrese duracion de partida partida:" + BCyan + "15 " + BWhite + "o " + BCyan + "30 " + BWhite + "puntos" + NONE
	CreateMatchMessage      = BWhite + "OK, Partida creada, esperando a que se una el resto de los jugadores" + NONE
	JoinMatchMessage        = BWhite + "OK, Partida solicitada, se esta buscando una partida" + NONE
	GameStartedMessage      = BWhite + "El juego comenzó" + NONE
	WaitingOptionsPlayer    = "Mientras esperas a que sea tu turno, podes realizar las siguientes acciones" + "\n" + GRAY + "(11) " + BYellow + "Irse" + NONE + " al mazo" + "\n" + GRAY + "(12) " + BGreen + "Consultar Cartas." + NONE + "\n"
	CardsMessage            = BWhite + "Estas son tus cartas: " + NONE
	WaitPlayerToPlayMessage = BBlue + "Espera a que juegue tu oponente..." + NONE + "\n"
	SingEnvido              = BPurple + "Cantaste ENVIDO" + NONE + "\n"
	SingEnvidoEnvido        = BPurple + "Cantaste ENVIDO ENVIDO" + NONE + "\n"
	SingTruco               = BPurple + "Cantaste TRUCO" + NONE + "\n"
	SingRetruco             = BPurple + "Cantaste RETRUCO" + NONE + "\n"
	AcceptEnvido            = BPurple + "Aceptaste ENVIDO" + NONE + "\n"
	AcceptTruco             = BPurple + "Aceptaste TRUCO" + NONE + "\n"
	AcceptRetruco           = BPurple + "Aceptaste RETRUCO" + NONE + "\n"
	SingFinishRound         = BPurple + "Te fuiste al MAZO" + NONE + "\n"
	RejectTruco             = BPurple + "Rechazaste TRUCO" + NONE + "\n"
	RejectRetruco           = BPurple + "Rechazaste RETRUCO" + NONE + "\n"
	RejectEnvido            = BPurple + "Rechazaste ENVIDO" + NONE + "\n"
	RejectEnvidoEnvido      = BPurple + "Rechazaste ENVIDO ENVIDO" + NONE + "\n"

	OpponentSingTruco          = BBlue + "Tu oponente canto TRUCO" + NONE + "\n"
	OpponentSingRetruco        = BBlue + "Tu oponente canto RETRUCO" + NONE + "\n"
	OpponetSingEnvido          = BBlue + "Tu oponente canto ENVIDO" + NONE + "\n"
	OpponetAcceptTruco         = BBlue + "Tu oponente Acepto el TRUCO" + NONE + "\n"
	OpponetAcceptEnvido        = BBlue + "Tu oponente Acepto el ENVIDO" + NONE + "\n"
	OpponetRejectTruco         = BBlue + "Tu oponente Rechazo el TRUCO" + NONE + "\n"
	OpponetHasSangFinishRound  = BBlue + "Tu oponente se fue AL MAZO" + NONE + "\n"
	OpponetRejectRetruco       = BBlue + "Tu oponente Rechazo el RETRUCO" + NONE + "\n"
	OpponetAcceptRetruco       = BBlue + "Tu oponente Acepto el RETRUCO" + NONE + "\n"
	OpponetRejectEnvido        = BBlue + "Tu oponente Rechazo el ENVIDO" + NONE + "\n"
	OpponentSingEnvidoEnvido   = BBlue + "Tu oponente canto ENVIDO ENVIDO" + NONE + "\n"
	OpponentAcceptEnvidoEnvido = BBlue + "Tu oponente Acepto ENVIDO" + NONE + "\n"
	OpponentRejectEnvidoEnvido = BBlue + "Tu oponente Rechazo ENVIDO ENVIDO" + NONE + "\n"
)

func GetWinningMoveMessage(number int) string {
	return BGreen + "Ganaste la jugada " + strconv.Itoa(number) + NONE + "\n"
}

func GetWinningRoundMessage(number int) string {
	return BGreen + "Ganaste la ronda " + strconv.Itoa(number) + NONE + "\n"
}

func GetLossingRoundMessage(number int) string {
	return BRed + "Perdiste la ronda " + strconv.Itoa(number) + NONE + "\n"
}

func GetLossingMoveMessage(number int) string {
	return BRed + "Perdiste la jugada " + strconv.Itoa(number) + NONE + "\n"
}

func GetPointsMessage(player1Points int, player2Points int) string {
	return BPurple + "Tus puntos son: " + strconv.Itoa(player1Points) + " Y los de tu oponente son: " + strconv.Itoa(player2Points) + NONE + "\n"
}

func GetWinningEnvidoMessage(player1Points int, player2Points int) string {
	message := BPurple + "Ganaste el Envido con " + strconv.Itoa(player1Points) + " puntos"
	return message + " y tu oponente perdio con " + strconv.Itoa(player2Points) + " puntos" + "\n"
}

func GetLossingEnvidoMessage(player1Points int, player2Points int) string {
	message := BPurple + "Perdiste el Envido con " + strconv.Itoa(player1Points) + " puntos"
	return message + " y tu oponente gano con " + strconv.Itoa(player2Points) + " puntos" + "\n"
}

func GetCardPlayed(card string) string {
	return BPurple + "Tiraste la carta " + NONE + card + "\n"
}
