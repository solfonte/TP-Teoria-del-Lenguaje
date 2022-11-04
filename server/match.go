package server

import "fmt"

type Match struct {
	duration   int
	maxPlayers int
	players    []Player
	started    bool
}

func (match *Match) addPlayerToMatch(player Player) {
	match.players = append(match.players, player)
	if len(match.players) == match.maxPlayers {
		match.started = true
		fmt.Println("Comenzo partida")

	}
}
