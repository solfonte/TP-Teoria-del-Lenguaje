package server

import "fmt"

type MatchManager struct {
	matches []Match
}

func process_player(matchManager MatchManager, player Player) {
	askPlayerName(player)
	messageClient, _ := sendMenu(player)
	requestedmatch := processRequest(player, messageClient)
	if requestedmatch["create"] == 0 {
		newMatch := Match{duration: requestedmatch["duration"], maxPlayers: requestedmatch["members"]}
		newMatch.players = append(newMatch.players, player)
		matchManager.matches = append(matchManager.matches, newMatch)
		fmt.Println("Matches ", matchManager.matches)
	} else {
		go look_matches_with_criteria(requestedmatch["duration"], requestedmatch["members"])
	}
}

func look_matches_with_criteria(duration int, cantPlayers int) {
	
}
