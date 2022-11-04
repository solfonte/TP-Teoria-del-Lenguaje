package server

import "fmt"

type MatchManager struct {
	matches []Match
}

func process_player(matchManager *MatchManager, player Player) {
	askPlayerName(player)
	messageClient, _ := sendMenu(player)
	requestedmatch := processRequest(player, messageClient)
	if requestedmatch["create"] == 0 {
		newMatch := Match{duration: requestedmatch["duration"], maxPlayers: requestedmatch["members"], started: false, players: []Player{}}
		addPlayerToMatch(&newMatch, player)
		matchManager.matches = append(matchManager.matches, newMatch)
		fmt.Println("Matches ", matchManager.matches)
	} else {
		fmt.Println("Matches antes de pasarlos: ", matchManager.matches)
		match := look_matches_with_criteria(matchManager.matches, requestedmatch["duration"], requestedmatch["members"])
		addPlayerToMatch(match, player)
	}
}

func look_matches_with_criteria(matches []Match, duration int, maxPlayers int) *Match {
	fmt.Println("Entre a buscar matches")
	fmt.Println("Matches ", matches)
	for _, match := range matches {
		fmt.Println("ewntre al for")
		if match.duration == duration && match.maxPlayers == maxPlayers && !match.started {
			fmt.Println(match.duration, match.maxPlayers, match.players)
			return &match
		}
	}
	return nil
}
